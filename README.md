# AeroStream Go Test

Testing [AeroStream](https://github.com/gradientgeeks/aerostream) with `segmentio/kafka-go` — a Go producer and consumer against a Kafka-compatible streaming broker.

## Setup

```bash
# Start AeroStream (1 controller + 1 broker)
docker compose up -d

# Install dependency
go mod tidy
```

## Run

```bash
# Terminal 1 — consumer (listens for messages)
go run consumer/main.go

# Terminal 2 — producer (sends 10 messages)
go run producer/main.go
```

## Findings

See [FINDINGS.md](FINDINGS.md) for compatibility notes.
