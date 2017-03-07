# Federation Hello World

## Setup

```
kubectl use-context harbor0
kubectl get cs 
kubectl get nodes
```

## Deply Services before Demo!!!

We did not set up the federated service feature, create them ahead of time and say you can use any sort of load balancing to tie this together.

```
kubectl create -f hello-federation-svc.yaml
kubectl create -f hello-federation-svc.yaml --context=aws
kubectl create -f hello-federation-svc.yaml --context=gke
```

## Deply hello-federation

We're just going to deploy a replicaset for our ugly demo app
```
kubectl create -f hello-federation-dp.yaml --context=federation
```

Let's check it out

```
kubectl get rs --context=federation
kubect get pods --context=harbor0
kubect get pods --context=aws
kubect get pods --context=gke
```

Now we can connect to the app
```
kubectl get svc hello-federation --context=harbor0
kubectl get svc hello-federation --context=aws
kubectl get svc hello-federation --context=gke
```

Grab the external-ip or service ip and put it in a browser, there's our app!

