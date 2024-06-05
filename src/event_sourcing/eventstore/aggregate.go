package eventstore

import (
	"reflect"
	"test/src/event_sourcing/event"
)

type Aggregator interface {
	GetId() string
	SetId(id string)
	GetType() string
	SetType(typ string)
	GetVersion() event.Version
	SetVersion(version event.Version)
	ListCommittedEvents() []event.Eventer
	ListUncommittedEvents() []event.Eventer
	Apply(event event.Eventer) error
	ApplyCommitted(event event.Eventer) error
	Commit(event event.Eventer) error
}

type AggregateCluster struct {
	currentId      string
	currentType    string
	currentVersion event.Version
	// internal.
	committedEvents   []event.Eventer
	uncommittedEvents *linkedList
	transitionfn      event.Transition
}

var _ (event.Aggregator) = &AggregateCluster{}

func New(agg event.Aggregator, transition event.Transition, idgenfn IDGenerator) *AggregateCluster {
	return &AggregateCluster{
		currentId:         idgenfn(idDefaultAlphabet, idDefaultSize),
		currentType:       reflect.TypeOf(agg).Elem().Name(),
		committedEvents:   make([]event.Eventer, 0, 8),
		uncommittedEvents: new(linkedList),
		transitionfn:      transition,
	}
}

func (r *AggregateCluster) GetId() string {
	return r.currentId
}

func (r *AggregateCluster) SetId(id string) {
	r.currentId = id
}

func (r *AggregateCluster) GetType() string {
	return r.currentType
}

func (r *AggregateCluster) SetType(typ string) {
	r.currentType = typ
}

func (r *AggregateCluster) GetVersion() event.Version {
	return r.currentVersion
}

func (r *AggregateCluster) SetVersion(version event.Version) {
	r.currentVersion = version
}
