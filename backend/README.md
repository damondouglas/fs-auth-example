Backend
=======

This directory holds code for the backend service.

# Prerequisites

The following is required to maintain and run the example.

- https://golang.org/
- https://github.com/google/ko
- https://cloud.google.com/sdk/docs/install

Additionally, the project requires setting up the following Google Cloud artifacts.

- https://firebase.google.com/products/firestore
- https://cloud.google.com/identity-platform

# Run 

To run the service locally:

```
make run
```

# Build and Deploy

## TCP service

To deploy the TCP service:

```
make build-deploy-tcp
```

## HTTP service

To deploy the HTTP service:

```
make build-deploy-http
```
