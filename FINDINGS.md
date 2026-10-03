# Findings

Tested AeroStream with `segmentio/kafka-go` v0.4.51 on a single-node Docker setup (1 controller + 1 broker).

## What works

- Connect, handshake, and API version negotiation — all good, 34 API keys supported
- Topic creation via CreateTopics API
- Producing messages (single and batch)
- Consuming messages with consumer groups — offsets tracked correctly

Basically the core Kafka produce/consume flow works out of the box with a standard Go Kafka client.

## What doesn't

**Auto topic creation is not supported.** If you produce to a topic that doesn't exist, you get `Unknown Topic Or Partition` instead of the topic being auto-created. In Kafka this is on by default (`auto.create.topics.enable=true`). Here you have to create topics explicitly before producing. Could trip up users migrating from Kafka.

**Broker advertises `0.0.0.0:9092` in metadata.** Works fine locally but would break in multi-machine setups. Probably needs an `advertised.listeners` equivalent.

## Setup

- Image: `quay.io/gradientgeeks/aerostream:latest`
- Machine: Linux, ~4 GB free RAM, 12 cores
- Ran fine, no crashes, no OOM
