# AVX 512 Demo

## Setup
Switch to the right context
```
kubectl config use-context harbor0
kubectl get cs
kubectl get clusters
kubectl get nodes
```

## Demo
Try running a job on our cluster

```
kubectl create -f dotproduct-job.yaml
```

How's our job doing?

```
kubectl get pods
kubectl logs dotproduct-xxxx
```

Job is dying?

Ah, this needs a certain type of CPU, we can target it accordingly

```
kubectl delete job dotproduct
kubectl create -f dotproduct-job-cpu.yaml
```
Stuck pending? Looks like our cluster doesn't have the right CPUs.

Let's see which of our clusters has the right ones.

```
kubectl get clusters --context=federation
```

We have 3 clusters, let's launch the task across all clusters
```
kubectl delete job dotproduct
kubectl create -f dotproduct-rs.yaml --context=federation
```

We can see that the job is pending on this cluster.

```
kubectl get pods --context=harbor0
kubectl describe pod dotproduct-xxxx
```

Let's take a look at how the CPU targeting works

```
kubectl describe node node4
```

```
less dotproduct-rs.yaml
```

Let's see how things are doing now

```
kubectl get pods --context=harbor0
kubectl get pods --context=aws
kubectl get pods --context=gke
kubectl logs dotproduct-xxxx --context=gke
```

Awesome! Kubernetes has noticed that pods have been pending for too long, so it naturally gravitated the pods to where they would run!


Knowing this, let's modify it so that our workload only runs on our Google Cloud cluster.

```
less dotproduct-rs-gke.yaml
kubectl apply -f dotproduct-rs-gke.yaml --context=federation
```

Now you're playing with power, Federation Power!

## Cleanup

```
kubectl delete -f . --context=federation
kubectl delete -f .
```
