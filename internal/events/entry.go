package events

import (
	"encoding/json"

	"github.com/malijoe/DatacenterGenerator/internal/encoding"
)

type Entry struct {
	AggregateID   string  `json:"aggregateID"`
	AggregateType string  `json:"aggregateType"`
	Version       int64   `json:"version"`
	EventStream   []Event `json:"event_stream"`
}

func (e Entry) Marshal() any {
	return struct {
		AggregateID   string  `json:"aggregateID"`
		AggregateType string  `json:"aggregateType"`
		Version       int64   `json:"version"`
		EventStream   []Event `json:"eventStream"`
	}{
		AggregateID:   e.AggregateID,
		AggregateType: e.AggregateType,
		Version:       e.Version,
		EventStream:   e.EventStream,
	}
}

func (e Entry) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Marshal())
}

func (e *Entry) Unmarshal(unmarshal func(any) error) (err error) {
	var obj struct {
		AggregateID   string  `json:"aggregateID"`
		AggregateType string  `json:"aggregateType"`
		Version       int64   `json:"version"`
		EventStream   []Event `json:"eventStream"`
	}
	if err = unmarshal(&obj); err != nil {
		return
	}
	*e = Entry{
		AggregateID:   obj.AggregateID,
		AggregateType: obj.AggregateType,
		Version:       obj.Version,
		EventStream:   obj.EventStream,
	}
	return
}

func (e *Entry) UnmarshalJSON(data []byte) error {
	return e.Unmarshal(encoding.JSONUnmarshalFunc(data))
}
