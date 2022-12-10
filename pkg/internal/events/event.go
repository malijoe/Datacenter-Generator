package events

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	uuid "github.com/satori/go.uuid"
)

type EventType string

type Event struct {
	EventId       string
	EventType     EventType
	Data          []byte
	Timestamp     time.Time
	AggregateType AggregateType
	AggregateId   string
	Version       int64
	Metadata      []byte
}

func NewBaseEvent(aggregate Aggregate, eventType EventType) Event {
	return Event{
		EventId:       uuid.NewV4().String(),
		AggregateType: aggregate.GetType(),
		AggregateId:   aggregate.GetId(),
		Version:       aggregate.GetVersion(),
		EventType:     eventType,
		Timestamp:     time.Now().UTC(),
	}
}

func NewEventFromRecorded(event *esdb.RecordedEvent) Event {
	return Event{
		EventId:     event.EventID.String(),
		EventType:   EventType(event.EventType),
		Data:        event.Data,
		Timestamp:   event.CreatedDate,
		AggregateId: event.StreamID,
		Version:     int64(event.EventNumber),
		Metadata:    event.UserMetadata,
	}
}

func NewEventFromEventData(event esdb.EventData) Event {
	return Event{
		EventId:   event.EventID.String(),
		EventType: EventType(event.EventType),
		Data:      event.Data,
		Metadata:  event.Metadata,
	}
}

func (e *Event) ToEventData() esdb.EventData {
	return esdb.EventData{
		EventType:   string(e.EventType),
		ContentType: esdb.JsonContentType,
		Data:        e.Data,
		Metadata:    e.Metadata,
	}
}

func (e *Event) GetEventId() string {
	return e.EventId
}

func (e *Event) GetTimestamp() time.Time {
	return e.Timestamp
}

func (e *Event) GetData() []byte {
	return e.Data
}

func (e *Event) SetData(data []byte) *Event {
	e.Data = data
	return e
}

func (e *Event) GetJsonData(data any) error {
	return json.Unmarshal(e.GetData(), data)
}

func (e *Event) SetJsonData(data any) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	e.Data = dataBytes
	return nil
}

func (e *Event) GetEventType() EventType {
	return e.EventType
}

func (e *Event) GetAggregateType() AggregateType {
	return e.AggregateType
}

func (e *Event) SetAggregateType(aggregateType AggregateType) {
	e.AggregateType = aggregateType
}

func (e *Event) GetAggregateId() string {
	return e.AggregateId
}

func (e *Event) GetVersion() int64 {
	return e.Version
}

func (e *Event) SetVersion(version int64) {
	e.Version = version
}

func (e *Event) GetMetadata() []byte {
	return e.Metadata
}

func (e *Event) SetMetadata(metadata any) error {
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	e.Metadata = metadataBytes
	return nil
}

func (e *Event) GetJsonMetadata(metadata any) error {
	return json.Unmarshal(e.GetMetadata(), metadata)
}

func (e *Event) GetString() string {
	return fmt.Sprintf("event: %+v", e)
}

func (e *Event) String() string {
	return fmt.Sprintf("(Event): AggregateId: {%s}, Version: {%d}, EventType: {%s}, AggregateType: {%s}, Metadata: {%s}. Timestamp: {%s}",
		e.AggregateId,
		e.Version,
		e.EventType,
		e.AggregateType,
		string(e.Metadata),
		e.Timestamp.UTC().String(),
	)
}
