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

type when func(event Event) error

type Aggregate interface {
	AggregateRoot
}

type AggregateRoot interface {
	GetUncommittedEvents() []Event
	GetID() string
	GetVersion() int64
	GetUncommittedVersion() int64
	CommitEvents()
	ClearUncommittedEvents()
	ToSnapshot()
	GetType() AggregateType
	SetAppliedEvents(events []Event)
	GetAppliedEvents() []Event
	Add(event Event) error
	String() string
	Apply(event Event) error
	Load(events []Event) error
}

type aggregateRoot struct {
	id                 string
	version            int64
	appliedEvents      []Event
	uncommittedEvents  []Event
	typ                AggregateType
	withAppliedEvents  bool
	when               when
	appliedVersion     int64
	uncommittedVersion int64
}

func NewAggregateRoot(id string, typ AggregateType, when when) *aggregateRoot {
	// delay panic
	if when == nil {
		return nil
	}

	return &aggregateRoot{
		id:                id,
		typ:               typ,
		version:           aggregateStartVersion,
		appliedEvents:     make([]Event, 0, aggregateAppliedEventsInitialCap),
		uncommittedEvents: make([]Event, 0, aggregateUncommittedEventsInitialCap),
		when:              when,
		withAppliedEvents: false,
	}
}

func (a *aggregateRoot) GetID() string {
	return a.id
}

func (a *aggregateRoot) GetType() AggregateType {
	return a.typ
}

func (a *aggregateRoot) GetVersion() int64 {
	return a.version
}

func (a *aggregateRoot) GetUncommittedVersion() int64 {
	return a.uncommittedVersion
}

func (a *aggregateRoot) ClearUncommittedEvents() {
	a.uncommittedEvents = make([]Event, 0, aggregateUncommittedEventsInitialCap)
}

func (a *aggregateRoot) GetAppliedEvents() []Event {
	return a.appliedEvents
}

func (a *aggregateRoot) CommitEvents() {
	a.appliedEvents = append(a.appliedEvents, a.uncommittedEvents...)
	a.withAppliedEvents = true
	a.ClearUncommittedEvents()
}

func (a *aggregateRoot) SetAppliedEvents(events []Event) {
	a.appliedEvents = events
	a.withAppliedEvents = true
}

func (a *aggregateRoot) GetUncommittedEvents() []Event {
	return a.uncommittedEvents
}

func (a *aggregateRoot) Load(events []Event) error {
	for _, evt := range events {
		if evt.GetAggregateID() != a.GetID() {
			return ErrInvalidAggregate
		}

		if err := a.Apply(evt); err != nil {
			return err
		}
	}
	return nil
}

func (a *aggregateRoot) Add(event Event) error {
	if event.GetAggregateID() != a.GetID() {
		return ErrInvalidAggregateId
	}

	event.SetAggregateType(a.GetType())

	if err := a.when(event); err != nil {
		return err
	}

	a.version++
	event.SetVersion(a.GetVersion())
	a.uncommittedEvents = append(a.uncommittedEvents, event)
	return nil
}

func (a *aggregateRoot) Apply(event Event) error {
	if event.GetAggregateID() != a.GetID() {
		return ErrInvalidAggregateId
	}
	if a.GetVersion() >= event.GetVersion() {
		fmt.Println(a.GetVersion(), event.GetVersion(), event.GetEventType())
		return ErrInvalidEventVersion
	}

	event.SetAggregateType(a.GetType())

	if err := a.applyEvent(event); err != nil {
		return err
	}

	a.version = event.GetVersion()
	return nil
}

func (a *aggregateRoot) applyEvent(event Event) error {
	if err := a.when(event); err != nil {
		return err
	}
	if !a.withAppliedEvents {
		a.withAppliedEvents = true
	}
	a.appliedEvents = append(a.appliedEvents, event)
	return nil
}

func (a *aggregateRoot) ToSnapshot() {
	if a.withAppliedEvents {
		a.appliedEvents = append(a.appliedEvents, a.uncommittedEvents...)
	}
	a.ClearUncommittedEvents()
}

func (a *aggregateRoot) String() string {
	return fmt.Sprintf("Id: {%s}, Version: {%v}, Type: {%v}, AppliedEvents: {%v}, UncommittedEvents: {%v}",
		a.GetID(),
		a.GetVersion(),
		a.GetType(),
		len(a.GetAppliedEvents()),
		len(a.GetUncommittedEvents()),
	)
}
