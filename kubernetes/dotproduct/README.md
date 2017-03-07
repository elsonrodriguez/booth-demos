# AVX 512 Demo

Switch to the right context
```
kubectl use-context harbor0
```

Try running a job on each cluster

```
kubectl create -f dotproduct-job.yaml
```

Job died?

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
kubectl create -f dotproduct-rs.yaml --context=federation
```

Let's see how things are doing

```
kubectl get pods --context=harbor0
kubectl get pods --context=aws
kubectl get pods --context=gke
```

Awesome!

Now let's modify it so that our workload only runs on our GKE cluster.

```
kubectl apply -f dotproduct-rs-gke.yaml --context=federation
```
