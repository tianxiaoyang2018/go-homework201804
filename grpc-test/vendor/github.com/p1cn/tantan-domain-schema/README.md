# tantan-domain-schema

```TODO: change repo name to putong-domain-tools```

Golang gRPC install guide:
```https://grpc.io/docs/quickstart/go.html```

## Generate go code from schemas

```sh
# make sure GOPATH is set

go get -u github.com/golang/protobuf/protoc-gen-go

export PATH=$PATH:$GOPATH/bin

git clone git@github.com:p1cn/tantan-domain-schema.git

cd tantan-domain-schema

_proto/build.sh
```

## What is the Domain Commit Log?
The Domain Commit Log is a record of transactions (insert, update, delete) of domain objects. Each domain object type has it's own record (topic), where all transactions of that type will be recorded. Services can consume on transactions per topic. Commits within a topic are ordered, and therefore they can be consumed in order.

## What is the difference between Commit Log, Message Queue and Pub/Sub?

### Commit Log
A commit log is a record of transactions. It's used to keep track of what has happened. Commits are written to the log before being applied, so transactions that are ongoing when a service is downed will get handled when the service gets back up again.

### Message Queue
MQ, see: https://aws.amazon.com/message-queue/

Message queues allow different parts of a system to communicate and process operations asynchronously. A message queue provides a lightweight buffer which temporarily stores messages, and endpoints that allow software components to connect to the queue in order to send and receive messages. Many producers and consumers can use the queue, but each message is processed only once, by a single consumer. For this reason, this messaging pattern is often called one-to-one, or point-to-point, communications. When a message needs to be processed by more than one consumer, message queues can be combined with Pub/Sub messaging in a fanout design pattern. 

### Pub/Sub
Pub/Sub, see https://aws.amazon.com/pub-sub-messaging/

A sibling to a message queue, a message topic provides a lightweight mechanism to broadcast asynchronous event notifications, and endpoints that allow software components to connect to the topic in order to send and receive those messages. To broadcast a message, a component called a publisher simply pushes a message to the topic. Unlike message queues, which batch messages until they are retrieved, message topics transfer messages with no or very little queuing, and push them out immediately to all subscribers. All components that subscribe to the topic will receive every message that is broadcast, unless a message filtering policy is set by the subscriber.
