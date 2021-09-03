# About

The purpose of this example is to demonstrate a web and mobile UI communicating
via gRPC to a backend service.  Combining the use of firestore and gRPC within a
web and mobile context may be redundant since there are secure ways to connect
a web and mobile application directly to firestore/firebase.  Nonetheless, the
value of this example may additionally show how a gRPC service may connect and
manage firestore data.

The following technologies are used in this example:
- Google Cloud Run
- gRPC and gRPC-web
- Firebase and Firestore
- Protocol Buffers

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
