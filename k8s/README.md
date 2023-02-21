# Kubernetes infra

## Prerequisites

* curl
* Kind, https://kind.sigs.k8s.io/ , for local Kubernetes testing
* kubectl, for interacting with Kubernetes cluster

## Setting up cluster

### Create local cluster with kind, specify some extra port mappings

```
kind create cluster --kubeconfig $HOME/.kube/config
export KUBECONFIG=$HOME/.kube/config
```


### Install our microservice into local kind cluster 

```
kind load docker-image craigcryptoapp:main
kubectl apply -f ./
```

### Verify that service is up
```
kubectl get svc cryptoserver
kubectl get pods cryptoserver-0
```

### Port forward, to access service
```
kubectl port-forward service/cryptoserver 8081:8081
```

**TODO: In future, replace kubectl port-forward with an Ingress controller or LoadBalancer**
`

### Test the crypto server with curl

```
curl -v  http://localhost:8081/currency/all
curl -v  http://localhost:8081/currency/btc
curl -v  http://localhost:8081/currency/nonexistent
```

### Test the crypto server with `go test`

```
cd ..
make test
```
