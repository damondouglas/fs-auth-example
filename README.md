# About

This repository shows an example using the following technologies:
- Google Cloud Run
- gRPC and gRPC-web
- Firebase and Firestore

# Requirements

- Prerequisites listed for https://grpc.io/docs/languages/go/quickstart/#prerequisites
- Prerequisites listed for https://grpc.io/docs/languages/dart/quickstart/#prerequisites

# proto

For convenience, this project uses [buf](https://docs.buf.build/introduction) to build protobuf.
If you prefer not to use this utility, follow specific language generation instructions as documented (See https://grpc.io).

```
buf generate
```

# Backend

See [backend/README](backend/README.md) for details on setup and running.

# Frontend

See [frontend/README](frontend/README.md) for details on setup and running.
