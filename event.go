package spacedk

import (
	"time"
	"container/heap"
)

// Represents an event that can be transmitted from sdk
type Event struct {
	Type EventType
	Context map[string]string
}

func (e *Event) IsEqualTo(other *Event) bool {
	if e.Type != other.Type {
		return false
	}

	if len(e.Context) != len(other.Context) {
		return false
	}

	for key, value := range e.Context {
		if value != other.Context[key] {
			return false
		}
	}

	return true
}

type EventType int
const (
	// Will be sent when cooldown of the ship expires
	// Additionally, provides ['ship']ShipSymbol as context
	CooldownExpiredEvent EventType = iota

	// Sent when ship arrives at its destination (navInfo)
	// Has 'shipSymbol' and 'destinationSymbol' contexts
	//
	// Although its here, im still not sure how to implement
	// the change in arrival date, when ship's fly mode is modified.
	// For now it is never sent
	ShipArrivedEvent

	// Sent when contract accept deadline has been reached
	// Has 'contractId' in context
	ContractExpiredEvent

	// Sent when contract has reached its delivery deadline
	// Has 'contractId' in context
	ContractDeadlineEvent
)

type EventScheduler struct {
	schedule *EventPriorityQueue

	Events chan *Event
}

func NewEventScheduler(channel chan *Event) *EventScheduler {
	return &EventScheduler{
		schedule: &EventPriorityQueue{},
		Events: channel,
	}
}

func (es *EventScheduler) Add(event *Event, transmitAt time.Time) {
	element := &EventElement{
		value: event,
		TransmitAt: transmitAt,
	}

	heap.Push(es.schedule, element)
}

func (es *EventScheduler) Remove(event *Event) {
	element := &EventElement{
		value: event,
	}

	es.schedule.Delete(element)
}

func (es *EventScheduler) Transmit() bool {
	if time.Now().After(es.schedule.Peak().TransmitAt) {
		es.Events<-heap.Pop(es.schedule).(*EventElement).value
		return true
	}

	return false
}

func (es *EventScheduler) StartDispatch(stop <-chan struct{}) {
	go func() {
		ticker := time.NewTicker(50 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-stop:
				return
			case <-ticker.C:
				for {
					transmited := es.Transmit()
					if !transmited {
						break
					}
				}
			}
		}
	}()
}
