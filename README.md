# About

This repository shows an example using the following technologies:
- Google Cloud Run
- gRPC and gRPC-web
- Firebase and Firestore

# Requirements

- Prerequisites listed for https://grpc.io/docs/languages/go/quickstart/#prerequisites
- Prerequisites listed for https://grpc.io/docs/languages/dart/quickstart/#prerequisites

# Build

To build the project, run:

```
bazel run //:gazelle
```

# Service

To run the backend service:

```
bazel run //backend/cmd/service
```