# Cryptocurrency Microservice

This repository contains a simple microservice in Go which accesses the api.hitbtc.com API server to obtain prices from various cryptocurrencies.

### Prerequisites

To build this microservice, the following tools must be installed
in your development environment.

- Go 1.16
- GNU make
- Docker
- Kind, https://kind.sigs.k8s.io/ , for local Kubernetes testing
- kubectl, for interacting with Kubernetes cluster

### Preparing environment

To obtain golint and staticcheck:

```
make get-gotools
```

### Building Microservice locally

To build the microservice locally:

```
make build
```

### Building Docker image

To build the Docker image, named `craigcryptoapp:main` containing the microservice:

```
make image
```

### Run in Kubernetes

After building the Docker image

```
cd ./k8s
```

and read [these instructions](./k8s/README.md) to run on Kubernetes

### Running tests

After deploying in Kubernetes, run the following to run tests:

```
make test
```
