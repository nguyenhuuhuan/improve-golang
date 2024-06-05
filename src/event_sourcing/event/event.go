package event

import (
	"database/sql/driver"
	"time"
)

type Eventer interface {
	GetAggregateId() string
	SetAggregateId(id string)
	GetAggregateType() string
	SetAggregateType(typ string)
	GetReason() string
	SetReason(reason string)
	GetVersion() Version
	SetVersion(version Version)
	GetTimestamp() Timestamp
	SetTimestamp(tstamp Timestamp)
	GetPayload() Payload
	SetPayload(payload Payload)
	GetSerializer() SerializerType
	SetSerializer(s SerializerType)
}

type Event struct {
	aggregateId    string
	aggregateType  string
	reason         string
	version        Version
	tstamp         Timestamp
	payload        Payload
	serializerType SerializerType
}

type Version int

const (
	DirtyVersion Version = -1
	EmptyVersion Version = 0
	NextVersion  Version = 1
)

type Timestamp time.Time

func (t *Timestamp) Scan(value interface{}) error {
	*t = Timestamp(value.(time.Time))
	return nil
}

func (t Timestamp) Value() (driver.Value, error) {
	return time.Time(t), nil
}

type Payload []byte

func (evt *Event) GetAggregateId() string {
	return evt.aggregateId
}

func (evt *Event) SetAggregateId(aggId string) {
	evt.aggregateId = aggId
}

func (evt *Event) GetAggregateType() string {
	return evt.aggregateType
}

func (evt *Event) SetAggregateType(aggType string) {
	evt.aggregateType = aggType
}

func (evt *Event) GetReason() string {
	return evt.reason
}

func (evt *Event) SetReason(reason string) {
	evt.reason = reason
}

func (evt *Event) GetVersion() Version {
	return evt.version
}

func (evt *Event) SetVersion(version Version) {
	evt.version = version
}

func (evt *Event) GetTimestamp() Timestamp {
	return evt.tstamp
}

func (evt *Event) SetTimestamp(tstamp Timestamp) {
	evt.tstamp = tstamp
}

func (evt *Event) GetPayload() Payload {
	return evt.payload
}

func (evt *Event) SetPayload(payload Payload) {
	evt.payload = payload
}

func (evt *Event) GetSerializer() SerializerType {
	return evt.serializerType
}

func (evt *Event) SetSerializer(typ SerializerType) {
	evt.serializerType = typ
}

var _ (Eventer) = &Event{}

type Transition func(event Eventer) error
