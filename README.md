# About

This repository shows an example using the following technologies:
- Google Cloud Run
- gRPC and gRPC-web
- Firebase and Firestore

# Requirements

- Prerequisites listed for https://grpc.io/docs/languages/go/quickstart/#prerequisites
- Prerequisites listed for https://grpc.io/docs/languages/dart/quickstart/#prerequisites

# Building

## Update BUILD files

To update bazel BUILD files

```
bazel run //:gazelle
```

## Listing build targets

To list available build targets:

```
bazel query ...
```

# proto

To update proto code generation:

```
bazel build //counter/v1:counter
```

# Service

To run the backend service:

```
bazel run //backend/cmd/service
```