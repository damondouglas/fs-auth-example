# About

This directory holds static application assets.

## config.json

Before running and deploying the app, a `config.json` file is required with the following details.
The host details are obtained from the deployed [backend](../../backend).

```
{
  "tcp": {
    "host": "<backend tcp host>",
    "port": 443
  },
  "http": {
    "host": "<backend http host>",
    "port": 443
  }
}
```

Example:

```
{
  "tcp": {
    "host": "counter-tcp-ne33nljv4q-uc.a.run.app",
    "port": 443
  },
  "http": {
    "host": "counter-http-ne33nljv4q-uc.a.run.app",
    "port": 443
  }
}
```
