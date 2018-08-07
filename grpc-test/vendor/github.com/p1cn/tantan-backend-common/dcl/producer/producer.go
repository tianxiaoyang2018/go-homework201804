package producer

// TODO: metrics

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/p1cn/tantan-backend-common/config"
	"github.com/p1cn/tantan-backend-common/dcl/eventmeta"
	slog "github.com/p1cn/tantan-backend-common/log"
	"github.com/p1cn/tantan-backend-common/metrics"
	"github.com/p1cn/tantan-backend-common/util/tracing"
	"github.com/p1cn/tantan-domain-schema/golang/event"
)

var (
	ErrInvalidEvent = errors.New("invalid event")
)

func twoUsersInteraction(id1 string, id2 string) string {
	stringInteraction := func(id1S, id2S string) string {
		if id1S < id2S {
			return id1S + ":" + id2S
		}
		return id2S + ":" + id1S
	}

	intInteraction := func(id1I, id2I int) string {
		if id1I < id2I {
			return fmt.Sprintf("%d:%d", id1I, id2I)
		}
		return fmt.Sprintf("%d:%d", id2I, id1I)
	}

	idInt1, err := strconv.Atoi(id1)
	if err != nil {
		return stringInteraction(id1, id2)
	}

	idInt2, err := strconv.Atoi(id2)
	if err != nil {
		return stringInteraction(id1, id2)
	}

	return intInteraction(idInt1, idInt2)
}

var topicKeys = map[string]func(*event.Event) string{
	eventmeta.UserTopic:      func(e *event.Event) string { return e.User.New.Id },
	eventmeta.DeviceTopic:    func(e *event.Event) string { return e.Device.New.OwnerId },
	eventmeta.BuildInfoTopic: func(e *event.Event) string { return e.BuildInfo.New.OwnerId },
	eventmeta.MatchTopic: func(e *event.Event) string {
		return twoUsersInteraction(e.Match.New.GetUserId(), e.Match.New.GetOtherUserId())
	},
	eventmeta.MessageTopic: func(e *event.Event) string {
		return twoUsersInteraction(e.Message.GetNew().GetUserId(), e.Message.GetNew().GetOtherUserId())
	},
	eventmeta.ConversationTopic:        func(e *event.Event) string { return e.GetConversation().GetNew().GetUserId() },
	eventmeta.MomentTopic:              func(e *event.Event) string { return e.GetMoment().GetNew().GetUserId() },
	eventmeta.MomentLikeTopic:          func(e *event.Event) string { return e.GetMomentLike().GetNew().GetUserId() },
	eventmeta.MomentCommentTopic:       func(e *event.Event) string { return e.GetMomentComment().GetNew().GetUserId() },
	eventmeta.RelationshipTopic:        func(e *event.Event) string { return e.GetRelationship().GetNew().GetUserId() },
	eventmeta.ScenarioUserCounterTopic: func(e *event.Event) string { return e.GetScenarioUserCounter().GetNew().GetUserId() },
	eventmeta.BlockTopic:               func(e *event.Event) string { return e.GetBlock().GetNew().GetOtherUserId() },
	eventmeta.FriendshipTopic:          func(e *event.Event) string { return e.GetFriendship().GetNew().GetOtherUserID() },
	eventmeta.VoteTopic:                func(e *event.Event) string { return e.GetVote().GetNew().GetVoteID() },
	eventmeta.VerificationTopic:        func(e *event.Event) string { return e.GetVerification().GetNew().GetUserID() },
	eventmeta.UserVerificationTopic:    func(e *event.Event) string { return e.GetUserVerification().GetNew().GetUserID() },
	eventmeta.UserStatusChangeTopic:    func(e *event.Event) string { return e.GetUser().GetNew().GetId() },
	eventmeta.UserActivityTopic:        func(e *event.Event) string { return e.GetUserActivity().GetNew().GetUserID() },
	eventmeta.FeedbackReportTopic:      func(e *event.Event) string { return e.GetReport().GetNew().GetId() },
	eventmeta.ContactTopic:             func(e *event.Event) string { return contactEventId(e) },
	eventmeta.PopularityTopic:          func(e *event.Event) string { return e.GetPopularity().GetUserId() },
	eventmeta.UserLinkTopic:            func(e *event.Event) string { return e.GetUserLink().GetNew().GetUserID() },
}

func contactEventId(e *event.Event) string {
	contactsNew := e.GetContact().GetNew()
	if len(contactsNew) > 0 {
		return contactsNew[0].GetId()
	}
	return ""
}

// ProducerConfig ...
type ProducerConfig struct {
	Name        string
	KafkaConfig config.KafkaConfig
	Partitioner sarama.PartitionerConstructor
}

type Producer interface {
	Commit(context.Context, *event.Event) error
	Close() error
}

// Producer ...
type producer struct {
	producer sarama.SyncProducer
	counter  *metrics.Timer
	cfg      ProducerConfig
}

func (p *producer) Close() error {
	return p.producer.Close()
}

var (
	counterVec     *metrics.Timer
	prometheusOnce sync.Once
)

// NewProducer ...
func NewProducer(cfg *ProducerConfig) (Producer, error) {

	kConfig := sarama.NewConfig()
	kConfig.Version = sarama.V0_10_0_0
	kConfig.Producer.Return.Errors = true
	kConfig.Producer.Return.Successes = true

	if cfg.KafkaConfig.Producer != nil {
		if cfg.KafkaConfig.Producer.MaxMessageBytes != nil {
			kConfig.Producer.MaxMessageBytes = *cfg.KafkaConfig.Producer.MaxMessageBytes
		}
		if cfg.KafkaConfig.Producer.RequiredAcks != nil {
			kConfig.Producer.RequiredAcks = sarama.RequiredAcks(*cfg.KafkaConfig.Producer.RequiredAcks)
		}

		if cfg.KafkaConfig.Producer.Timeout != nil {
			kConfig.Producer.Timeout = time.Duration(*cfg.KafkaConfig.Producer.Timeout)
		}

		if cfg.KafkaConfig.Producer.Compression != nil {
			kConfig.Producer.Compression = sarama.CompressionCodec(*cfg.KafkaConfig.Producer.Compression)
		}

		if cfg.Partitioner != nil {
			kConfig.Producer.Partitioner = cfg.Partitioner
		}
		if cfg.KafkaConfig.Producer.Flush != nil {
			kConfig.Producer.Flush.Bytes = cfg.KafkaConfig.Producer.Flush.Bytes
			kConfig.Producer.Flush.Messages = cfg.KafkaConfig.Producer.Flush.Messages
			kConfig.Producer.Flush.Frequency = time.Duration(cfg.KafkaConfig.Producer.Flush.Frequency)
			kConfig.Producer.Flush.MaxMessages = cfg.KafkaConfig.Producer.Flush.MaxMessages
		}
		if cfg.KafkaConfig.Producer.Retry != nil {
			kConfig.Producer.Retry.Max = cfg.KafkaConfig.Producer.Retry.Max
			kConfig.Producer.Retry.Backoff = time.Duration(cfg.KafkaConfig.Producer.Retry.Backoff)
		}
	}

	p, err := sarama.NewSyncProducer(cfg.KafkaConfig.Brokers, kConfig)
	if err != nil {
		return nil, err
	}

	prometheusOnce.Do(func() {
		counterVec = metrics.NewTimer(metrics.NameSpaceTantan, "dcl_producer", "DCL producer", []string{"name", "topic", "ret"})
	})

	return &producer{
		producer: p,
		counter:  counterVec,
		cfg:      *cfg,
	}, nil
}

// Commit ...
// @todo 生产者是否更新ServiceContext的ServiceTraces
// 1. 服务自己产生的dcl事件 ： 没有更新
// 2. rpc或者http产生的dcl事件：middleware已经更新
// @tod ：dcl应该打印tracing日志
func (p *producer) Commit(ctx context.Context, event *event.Event) error {
	record := p.counter.Timer()

	if event == nil {
		record(p.cfg.Name, "", "failure_event_nil")
		return ErrInvalidEvent
	}
	if event.GetContext() == nil {
		serviceContext := tracing.GetServiceContext(ctx)
		if serviceContext == nil {
			slog.Warning("context.Content doesn't contain service context.(ref util.NewServiceContextToContext)")
			serviceContext = tracing.NewServiceContext()
		}
		event.Context = serviceContext
	}

	data, err := proto.Marshal(event)
	if err != nil {
		record(p.cfg.Name, "", "failure_event_marshal")
		slog.Err("err : %+v , : event : %+v", err, event)
		return err
	}

	key := GetEventKey(event)
	msg := &sarama.ProducerMessage{
		Topic: event.Topic,
		Value: sarama.ByteEncoder(data),
	}

	if len(key) != 0 {
		msg.Key = sarama.StringEncoder(key)
	}

	_, _, err = p.producer.SendMessage(msg)
	if err != nil {
		record(p.cfg.Name, event.GetTopic(), "failure_send")
		slog.Err("err : %+v , event : %+v", err, event)
		return err
	} else {
		record(p.cfg.Name, event.GetTopic(), "OK")
	}
	return nil
}

// TODO: Add topic and key to event
func GetEventKey(event *event.Event) string {
	e, ok := topicKeys[event.Topic]
	if !ok {
		return ""
	}
	return e(event)
}
