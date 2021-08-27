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

# Service

To run the backend service:

```
cd backend
go run ./cmd/service
```

or show usage

```
cd backend
go run ./cmd/service -help
```
