package events

import (
	"context"
	"fmt"
)

const (
	aggregateStartVersion                = -1
	aggregateAppliedEventsInitialCap     = 10
	aggregateUncommittedEventsInitialCap = 10
)

type AggregateType string

type HandleCommand interface {
	HandleCommand(ctx context.Context, command Command) error
}

type When interface {
	When(event Event) error
}

type when func(event Event) error

type Apply interface {
	Apply(event Event) error
}

type Load interface {
	Load(events []Event) error
}

type Aggregate interface {
	AggregateRoot
}

type AggregateRoot interface {
	GetUncommittedEvents() []Event
	GetId() string
	SetId(id string) *AggregateBase
	GetVersion() int64
	ClearUncommittedEvents()
	ToSnapshot()
	SetType(aggregateType AggregateType)
	GetType() AggregateType
	SetAppliedEvents(events []Event)
	GetAppliedEvents() []Event
	RaiseEvent(event Event) error
	String() string
	Load
	Apply
}

type AggregateBase struct {
	Id                string
	Version           int64
	AppliedEvents     []Event
	UncommittedEvents []Event
	Type              AggregateType
	withAppliedEvents bool
	when              when
}

func NewAggregateBase(when when) *AggregateBase {
	if when == nil {
		return nil
	}

	return &AggregateBase{
		Version:           aggregateStartVersion,
		AppliedEvents:     make([]Event, 0, aggregateAppliedEventsInitialCap),
		UncommittedEvents: make([]Event, 0, aggregateUncommittedEventsInitialCap),
		when:              when,
		withAppliedEvents: false,
	}
}

func (a *AggregateBase) SetId(id string) *AggregateBase {
	a.Id = fmt.Sprintf("%s-%s", a.GetType(), id)
	return a
}

func (a *AggregateBase) GetId() string {
	return a.Id
}

func (a *AggregateBase) SetType(aggregateType AggregateType) {
	a.Type = aggregateType
}

func (a *AggregateBase) GetType() AggregateType {
	return a.Type
}

func (a *AggregateBase) GetVersion() int64 {
	return a.Version
}

func (a *AggregateBase) ClearUncommittedEvents() {
	a.UncommittedEvents = make([]Event, 0, aggregateUncommittedEventsInitialCap)
}

func (a *AggregateBase) GetAppliedEvents() []Event {
	return a.AppliedEvents
}

func (a *AggregateBase) SetAppliedEvents(events []Event) {
	a.AppliedEvents = events
}

func (a *AggregateBase) GetUncommittedEvents() []Event {
	return a.UncommittedEvents
}

func (a *AggregateBase) Load(events []Event) error {
	for _, evt := range events {
		if evt.GetAggregateId() != a.GetId() {
			return ErrInvalidAggregate
		}

		if err := a.when(evt); err != nil {
			return err
		}

		if a.withAppliedEvents {
			a.AppliedEvents = append(a.AppliedEvents, evt)
		}
		a.Version++
	}

	return nil
}

func (a *AggregateBase) Apply(event Event) error {
	if event.GetAggregateId() != a.GetId() {
		return ErrInvalidAggregateId
	}

	event.SetAggregateType(a.GetType())

	if err := a.when(event); err != nil {
		return err
	}

	a.Version++
	event.SetVersion(a.GetVersion())
	a.UncommittedEvents = append(a.UncommittedEvents, event)
	return nil
}

func (a *AggregateBase) RaiseEvent(event Event) error {
	if event.GetAggregateId() != a.GetId() {
		return ErrInvalidAggregateId
	}
	if a.GetVersion() >= event.GetVersion() {
		return ErrInvalidEventVersion
	}

	event.SetAggregateType(a.GetType())

	if err := a.when(event); err != nil {
		return err
	}

	if a.withAppliedEvents {
		a.AppliedEvents = append(a.AppliedEvents, event)
	}
	a.Version = event.GetVersion()
	return nil
}

func (a *AggregateBase) ToSnapshot() {
	if a.withAppliedEvents {
		a.AppliedEvents = append(a.AppliedEvents, a.UncommittedEvents...)
	}
	a.ClearUncommittedEvents()
}

func (a *AggregateBase) String() string {
	return fmt.Sprintf("Id: {%s}, Version: {%v}, Type: {%v}, AppliedEvents: {%v}, UncommittedEvents: {%v}",
		a.GetId(),
		a.GetVersion(),
		a.GetType(),
		len(a.GetAppliedEvents()),
		len(a.GetUncommittedEvents()),
	)
}
