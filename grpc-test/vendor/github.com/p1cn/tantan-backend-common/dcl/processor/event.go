package processor

import (
	"context"

	"github.com/p1cn/tantan-backend-common/dcl/eventmeta"
	"github.com/p1cn/tantan-domain-schema/golang/event"
)

// EventSyncProcessor ...
type EventSyncProcessor interface {
	Process(ctx context.Context, event *event.Event, meta *eventmeta.EventMetaData) error
}

// EventProcessor ...
type EventProcessor interface {
	Process(ctx context.Context, event *event.Event, meta *eventmeta.EventMetaData, done chan error)
}
