# Kubernetes and multi-cloud portability

A quick demo of how to run things across clouds

## Setup

```
kubectl get cs --context=harbor0
kubectl get cs --context=aws
kubectl get cs --context=gke
kubectl config use-context harbor0
```

## Deploy wordpress or some other app

Let's talk about portability.

We're going to deploy an app to our private cloud (or public cloud, setup is up to you)

Deploy storage
```
kubectl create -f mysql-pvc.yaml
kubectl create -f wordpress-pvc.yaml
```
Verify

```
kubectl get pv
kubectl get pvc
```

