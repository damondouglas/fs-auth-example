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

## Firestore specific setup

The example assumes a "counts" collection exists in the database.

## Identity Platform specific setup

The example assumes email/password authentication is turned on as the frontend
example uses this form of identity platform authentication.

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
