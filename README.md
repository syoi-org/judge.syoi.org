# judy

Forever work-in-progress repository for building a online judge.

## Prerequisites

None! Really. You can run a minimal setup without external dependencies. It's just a go binary. :D However, consider the
following when scaling up.

- Cache: Redis
- Database: postgreSQL
- Storage: S3
- Message Queue: RabbitMQ, Redis

## Getting Started

The minimal setup can be run without any configuration! Open two terminals and run the following:

```bash
judy migrate
judy serve
```

```bash
judy worker
```
