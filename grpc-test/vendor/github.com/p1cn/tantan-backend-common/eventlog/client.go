package eventlog

import (
	"errors"

	"github.com/pquerna/ffjson/ffjson"

	"github.com/p1cn/tantan-backend-common/eventlog/domain"
)

var EventLog *EventLogClient

var (
	ErrTimeout     = errors.New("EventLog RPC ERROR : Timtout")
	ErrInvalidDATA = errors.New("EventLog RPC ERROR : Invalid DATA")
	ErrDisabled    = errors.New("EventLog RPC ERROR : it is disabled")
)

type IClient interface {
	Start() error
	Close() error
	Send(message *RpcMessage) error
}

type EventLogClient struct {
	client  IClient
	enabled bool
}

func (e *EventLogClient) start() error {
	err := e.client.Start()
	if err == nil {
		e.enabled = true
	}
	return err
}

func (e *EventLogClient) close() error {
	if !e.enabled {
		return ErrDisabled
	}
	e.enabled = false
	return e.client.Close()
}

func (e *EventLogClient) sendData(message *RpcMessage) error {
	if !e.enabled {
		return ErrDisabled
	}
	return e.client.Send(message)
}

type RpcMessage struct {
	Key   []byte
	Topic string
	Data  []byte
	ID    int64
}

type EventLogRpcMessage struct {
	Topic string
	Event *domain.Event
	ID    int64
}

//
//  functions
//

func IsEnable() bool {
	return EventLog != nil
}

func Close() error {
	if !IsEnable() {
		return nil
	}
	return EventLog.close()
}

func SendEvent(msg *EventLogRpcMessage) error {
	if !IsEnable() {
		return nil
	}

	if msg == nil || len(msg.Topic) == 0 || msg.Event == nil {
		return ErrInvalidDATA
	}

	data, err := ffjson.Marshal(msg.Event)
	if err != nil {
		return err
	}

	return EventLog.sendData(&RpcMessage{
		Topic: msg.Topic,
		ID:    msg.ID,
		Data:  data,
	})
}

func SendData(msg *RpcMessage) error {
	if !IsEnable() {
		return nil
	}

	if msg == nil || len(msg.Topic) == 0 || len(msg.Data) == 0 {
		return ErrInvalidDATA
	}

	return EventLog.sendData(msg)
}
