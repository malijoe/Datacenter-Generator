package repositories

import (
	"bytes"
	"context"
	"encoding/json"

	boltdb "github.com/boltdb/bolt"
	"github.com/malijoe/DatacenterGenerator/internal/errorz"
	"github.com/malijoe/DatacenterGenerator/internal/events"
)

const (
	collection = "events"
	// terminator is the null character, used as a terminator.
	terminator = "\x00"
)

type (
	conText   = context.Context
	aggregate = events.Aggregate
	ntry      = events.Entry
)

type localStore struct {
	db *boltdb.DB
}

func NewLocalStore(db *boltdb.DB) (repo *localStore, err error) {
	var tx *boltdb.Tx
	tx, err = db.Begin(true)
	if err != nil {
		return
	}
	defer tx.Rollback()

	_, err = tx.CreateBucketIfNotExists([]byte(collection))
	if err != nil {
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}

	repo = &localStore{
		db: db,
	}
	return
}

func aggregateToEntry(agg aggregate) (entry *ntry) {
	agg.CommitEvents()
	entry = &ntry{
		AggregateID:   agg.GetID(),
		AggregateType: string(agg.GetType()),
		Version:       agg.GetVersion(),
		EventStream:   agg.GetAppliedEvents(),
	}
	return
}

func getAggregateKey(agg aggregate) (key []byte) {
	key = []byte(string(agg.GetType()) + " | " + agg.GetID() + terminator)
	return
}

func (r *localStore) withTxn(f func(bckt *boltdb.Bucket) error) (err error) {
	err = r.db.Update(func(tx *boltdb.Tx) error {
		bckt := tx.Bucket([]byte(collection))
		return f(bckt)
	})
	return
}

func (r *localStore) withSnapshot(f func(bckt *boltdb.Bucket) error) (err error) {
	err = r.db.View(func(tx *boltdb.Tx) error {
		bckt := tx.Bucket([]byte(collection))
		return f(bckt)
	})
	return
}

func (r *localStore) getAggregateEntry(aggType, id string) (entry *ntry, err error) {
	err = r.withSnapshot(func(bckt *boltdb.Bucket) error {
		var (
			cursor = bckt.Cursor()
			data   []byte
		)
		for k, v := cursor.Seek([]byte(aggType)); k != nil && bytes.HasPrefix(k, []byte(aggType)); k, v = cursor.Next() {
			idBytes := bytes.TrimPrefix(k, []byte(aggType+" | "))
			idBytes = bytes.TrimSuffix(idBytes, []byte(terminator))

			if id == string(idBytes) {
				data = v
				break
			}
		}
		if data == nil {
			return errorz.ErrNotFound
		}
		var e ntry
		if uErr := json.Unmarshal(data, &e); uErr != nil {
			return uErr
		}
		entry = &e
		return nil
	})
	return
}

func (r *localStore) SaveAggregate(_ conText, agg aggregate) (err error) {
	entry := aggregateToEntry(agg)

	var data []byte
	data, err = json.Marshal(entry)
	if err != nil {
		return
	}

	err = r.withTxn(func(bckt *boltdb.Bucket) error {
		return bckt.Put(getAggregateKey(agg), data)
	})
	return
}

func (r *localStore) LoadAggregate(_ conText, agg aggregate) (err error) {
	var entry *ntry
	entry, err = r.getAggregateEntry(string(agg.GetType()), agg.GetID())
	if err != nil {
		return
	}

	for _, e := range entry.EventStream {
		if err = agg.Apply(e); err != nil {
			return
		}
	}
	return
}

func (r *localStore) DeleteAggregate(_ conText, aggregateType string, id string) (err error) {
	err = r.withTxn(func(bckt *boltdb.Bucket) error {
		return bckt.Delete([]byte(aggregateType + " | " + id + terminator))
	})
	return
}

func (r *localStore) GetAggregateTypeEntries(aggregateType string) (entries []*ntry, err error) {
	err = r.withSnapshot(func(bckt *boltdb.Bucket) error {
		cursor := bckt.Cursor()
		entries = make([]*ntry, 0, bckt.Stats().KeyN)
		for k, v := cursor.Seek([]byte(aggregateType)); k != nil && bytes.HasPrefix(k, []byte(aggregateType)); k, v = cursor.Next() {
			var entry ntry
			if uErr := json.Unmarshal(v, &entry); uErr != nil {
				return uErr
			}
			entries = append(entries, &entry)
		}
		return nil
	})
	return
}
