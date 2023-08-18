# Trendle

---
>A self-hosted, real-time polling system.

## About

Trendle is a poll management system capable of persisting 
and responding to poll results in real-time through WebSocket connections.

## How to get started

Trendle uses PostgresSQL to persist the data, Redis to handle
the voting pubsub and Docker to containerize the application,
so make sure you have them in your environment.

First, set the following environment variables:
- **PORT** (where Trendle will listen)
- **REDIS_ADDR** (Redis address)
- **REDIS_PASS** (Redis password)
- **DATABASE_URL** (PostgresSQL connection string)

Then, build the Docker image running the following command:
```
docker build --target prod -t trendle-app .
```

Finally, run the image running the following command:
```
docker run trendle-app:latest
```