package config

type KafkaProducerReturnConfig struct {
	// If enabled, successfully delivered messages will be returned on the
	// Successes channel (default disabled).
	Successes bool

	// If enabled, messages that failed to deliver will be returned on the
	// Errors channel, including error (default enabled).
	Errors bool
}

type KafkaProducerFlushConfig struct {
	// The best-effort number of bytes needed to trigger a flush. Use the
	// global sarama.MaxRequestSize to set a hard upper limit.
	Bytes int
	// The best-effort number of messages needed to trigger a flush. Use
	// `MaxMessages` to set a hard upper limit.
	Messages int
	// The best-effort frequency of flushes. Equivalent to
	// `queue.buffering.max.ms` setting of JVM producer.
	Frequency Duration
	// The maximum number of messages the producer will send in a single
	// broker request. Defaults to 0 for unlimited. Similar to
	// `queue.buffering.max.messages` in the JVM producer.
	MaxMessages int
}

type KafkaProducerRetryConfig struct {
	// The total number of times to retry sending a message (default 3).
	// Similar to the `message.send.max.retries` setting of the JVM producer.
	Max int
	// How long to wait for the cluster to settle between retries
	// (default 100ms). Similar to the `retry.backoff.ms` setting of the
	// JVM producer.
	Backoff Duration
}

type KafkaProducerConfig struct {
	// The maximum permitted size of a message (defaults to 1000000). Should be
	// set equal to or smaller than the broker's `message.max.bytes`.
	MaxMessageBytes *int

	// The level of acknowledgement reliability needed from the broker (defaults
	// to WaitForLocal). Equivalent to the `request.required.acks` setting of the
	// JVM producer.
	// NoResponse doesn't send any response, the TCP ACK is all you get.
	// NoResponse RequiredAcks = 0
	// WaitForLocal waits for only the local commit to succeed before responding.
	// WaitForLocal RequiredAcks = 1
	// WaitForAll waits for all in-sync replicas to commit before responding.
	// The minimum number of in-sync replicas is configured on the broker via
	// the `min.insync.replicas` configuration key.
	// WaitForAll RequiredAcks = -1
	RequiredAcks *int

	// The maximum duration the broker will wait the receipt of the number of
	// RequiredAcks (defaults to 10 seconds). This is only relevant when
	// RequiredAcks is set to WaitForAll or a number > 1. Only supports
	// millisecond resolution, nanoseconds will be truncated. Equivalent to
	// the JVM producer's `request.timeout.ms` setting.
	Timeout *Duration

	// The type of compression to use on messages (defaults to no compression).
	// Similar to `compression.codec` setting of the JVM producer.
	// CompressionNone   CompressionCodec = 0
	// CompressionGZIP   CompressionCodec = 1
	// CompressionSnappy CompressionCodec = 2
	// CompressionLZ4    CompressionCodec = 3
	Compression *int

	// Return specifies what channels will be populated. If they are set to true,
	// you must read from the respective channels to prevent deadlock. If,
	// however, this config is used to create a `SyncProducer`, both must be set
	// to true and you shall not read from the channels since the producer does
	// this internally.
	// this configuration is unavailable on DCL
	Return *KafkaProducerReturnConfig

	// The following config options control how often messages are batched up and
	// sent to the broker. By default, messages are sent as fast as possible, and
	// all messages received while the current batch is in-flight are placed
	// into the subsequent batch.
	Flush *KafkaProducerFlushConfig
	Retry *KafkaProducerRetryConfig
}

type KafkaConsumerRetryConfig struct {
	// How long to wait after a failing to read from a partition before
	// trying again (default 2s).
	Backoff Duration
}

type KafkaConsumerFetchConfig struct {
	// The minimum number of message bytes to fetch in a request - the broker
	// will wait until at least this many are available. The default is 1,
	// as 0 causes the consumer to spin when no messages are available.
	// Equivalent to the JVM's `fetch.min.bytes`.
	Min int32
	// The default number of message bytes to fetch from the broker in each
	// request (default 32768). This should be larger than the majority of
	// your messages, or else the consumer will spend a lot of time
	// negotiating sizes and not actually consuming. Similar to the JVM's
	// `fetch.message.max.bytes`.
	Default int32
	// The maximum number of message bytes to fetch from the broker in a
	// single request. Messages larger than this will return
	// ErrMessageTooLarge and will not be consumable, so you must be sure
	// this is at least as large as your largest message. Defaults to 0
	// (no limit). Similar to the JVM's `fetch.message.max.bytes`. The
	// global `sarama.MaxResponseSize` still applies.
	Max int32
}

type KafkaConsumerReturnConfig struct {
	// If enabled, any errors that occurred while consuming are returned on
	// the Errors channel (default disabled).
	Errors bool
}

type KafkaConsumerOffsetsConfig struct {
	// How frequently to commit updated offsets. Defaults to 1s.
	CommitInterval Duration

	// OffsetNewest stands for the log head offset, i.e. the offset that will be
	// assigned to the next message that will be produced to the partition. You
	// can send this to a client's GetOffset method to get this offset, or when
	// calling ConsumePartition to start consuming new messages.
	//       OffsetNewest int64 = -1
	// OffsetOldest stands for the oldest offset available on the broker for a
	// partition. You can send this to a client's GetOffset method to get this
	// offset, or when calling ConsumePartition to start consuming from the
	// oldest offset that is still available on the broker.
	//       OffsetOldest int64 = -2

	// The initial offset to use if no offset was previously committed.
	// Should be OffsetNewest or OffsetOldest. Defaults to OffsetNewest.
	Initial int64

	// The retention duration for committed offsets. If zero, disabled
	// (in which case the `offsets.retention.minutes` option on the
	// broker will be used).  Kafka only supports precision up to
	// milliseconds; nanoseconds will be truncated. Requires Kafka
	// broker version 0.9.0 or later.
	// (default is 0: disabled).
	Retention Duration
}

type KafkaConsumerConfig struct {
	Retry *KafkaConsumerRetryConfig

	// Fetch is the namespace for controlling how many bytes are retrieved by any
	// given request.
	Fetch *KafkaConsumerFetchConfig
	// The maximum amount of time the broker will wait for Consumer.Fetch.Min
	// bytes to become available before it returns fewer than that anyways. The
	// default is 250ms, since 0 causes the consumer to spin when no events are
	// available. 100-500ms is a reasonable range for most cases. Kafka only
	// supports precision up to milliseconds; nanoseconds will be truncated.
	// Equivalent to the JVM's `fetch.wait.max.ms`.
	MaxWaitTime *Duration

	// The maximum amount of time the consumer expects a message takes to
	// process for the user. If writing to the Messages channel takes longer
	// than this, that partition will stop fetching more messages until it
	// can proceed again.
	// Note that, since the Messages channel is buffered, the actual grace time is
	// (MaxProcessingTime * ChanneBufferSize). Defaults to 100ms.
	// If a message is not written to the Messages channel between two ticks
	// of the expiryTicker then a timeout is detected.
	// Using a ticker instead of a timer to detect timeouts should typically
	// result in many fewer calls to Timer functions which may result in a
	// significant performance improvement if many messages are being sent
	// and timeouts are infrequent.
	// The disadvantage of using a ticker instead of a timer is that
	// timeouts will be less accurate. That is, the effective timeout could
	// be between `MaxProcessingTime` and `2 * MaxProcessingTime`. For
	// example, if `MaxProcessingTime` is 100ms then a delay of 180ms
	// between two messages being sent may not be recognized as a timeout.
	MaxProcessingTime *Duration

	// Return specifies what channels will be populated. If they are set to true,
	// you must read from them to prevent deadlock.
	Return *KafkaConsumerReturnConfig

	// Offsets specifies configuration for how and when to commit consumed
	// offsets. This currently requires the manual use of an OffsetManager
	// but will eventually be automated.
	Offsets *KafkaConsumerOffsetsConfig
}

type KafkaGroupOffsetsRetryConfig struct {
	// The numer retries when committing offsets (defaults to 3).
	Max int
}

type KafkaGroupOffsetsSynchronizationConfig struct {
	// The duration allowed for other clients to commit their offsets before resumption in this client, e.g. during a rebalance
	// NewConfig sets this to the Consumer.MaxProcessingTime duration of the Sarama configuration
	DwellTime Duration
}

type KafkaGroupOffsetsConfig struct {
	Retry           *KafkaGroupOffsetsRetryConfig
	Synchronization *KafkaGroupOffsetsSynchronizationConfig
}

type KafkaGroupSessionConfig struct {
	// The allowed session timeout for registered consumers (defaults to 30s).
	// Must be within the allowed server range.
	Timeout Duration
}

type KafkaGroupHeartbeatConfig struct {
	// Interval between each heartbeat (defaults to 3s). It should be no more
	// than 1/3rd of the Group.Session.Timout setting
	Interval Duration
}

type KafkaGroupReturnConfig struct {
	// If enabled, rebalance notification will be returned on the
	// Notifications channel (default disabled).
	Notifications bool
}

type KafkaGroupMemberConfig struct {
	// Custom metadata to include when joining the group. The user data for all joined members
	// can be retrieved by sending a DescribeGroupRequest to the broker that is the
	// coordinator for the group.
	UserData []byte
}

type KafkaGroupConfig struct {

	// The strategy to use for the allocation of partitions to consumers (defaults to StrategyRange)
	// StrategyRange is the default and assigns partition ranges to consumers.
	// Example with six partitions and two consumers:
	//   C1: [0, 1, 2]
	//   C2: [3, 4, 5]
	//StrategyRange Strategy = "range"

	// StrategyRoundRobin assigns partitions by alternating over consumers.
	// Example with six partitions and two consumers:
	//   C1: [0, 2, 4]
	//   C2: [1, 3, 5]
	//StrategyRoundRobin Strategy = "roundrobin"
	PartitionStrategy *string

	// By default, messages and errors from the subscribed topics and partitions are all multiplexed and
	// made available through the consumer's Messages() and Errors() channels.
	//
	// Users who require low-level access can enable ConsumerModePartitions where individual partitions
	// are exposed on the Partitions() channel. Messages and errors must then be consumed on the partitions
	// themselves.
	Mode *uint8

	Offsets   *KafkaGroupOffsetsConfig
	Session   *KafkaGroupSessionConfig
	Heartbeat *KafkaGroupHeartbeatConfig

	// Return specifies which group channels will be populated. If they are set to true,
	// you must read from the respective channels to prevent deadlock.
	Return *KafkaGroupReturnConfig
	Member *KafkaGroupMemberConfig
}

type KafkaConfig struct {
	Brokers    []string
	Zookeepers []string
	Producer   *KafkaProducerConfig
	Consumer   *KafkaConsumerConfig
	Group      *KafkaGroupConfig
}
