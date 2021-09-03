Frontend
========

This directory holds code for the frontend targeting web and mobile platforms.

# Prerequisites

The following is required to run, build, and maintain the application:
- Flutter: https://flutter.dev
- Firebase CLI: https://firebase.google.com/docs/cli
- Firebase project: https://console.firebase.google.com
- Deployed and configured backend (See [backend/README](../backend/README.md))

# Setup

The following setup is needed to run and deploy the application.

## assets/config.json

The `assets/config.json` file is required to inform the application about the backend endpoints.
See [assets/README](assets/README.md) for details.

## web/config/firebase.js

The `web/config/firebase.js` file is required to inform the web application about firebase 
configuration details.
See [web/config/README](web/config/README.md) for details.
Additional information may be found at https://firebase.flutter.dev/docs/installation/web.

## Android setup

Before using this app on Android, you must first connect to your Firebase project with your Android application.
See https://firebase.flutter.dev/docs/installation/android for details.

## iOS setup
Before using this app on iOS, you must first connect to your Firebase project with your Android application.
See https://firebase.flutter.dev/docs/installation/ios/ for details.

# Run application

To run the application, follow standard `flutter run` documentation.  For convenience, the [Makefile](Makefile)
contains a target for running a local web server.

# Build application

To build the application, follow standard `flutter build` documentation.

# Deploy application

Follow firebase hosting instructions to deploy the application (https://firebase.google.com/docs/hosting).
