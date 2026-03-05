# Pre requisites
* Locally or cluster you fully control (=== configure control plane), where you can
  * Run some docker daemon
    * [Docker desktop](https://www.docker.com/products/docker-desktop/)
  * Install some local cluster
    * tool
      * [minikube](https://minikube.sigs.k8s.io/docs/start/)
      * [kind](https://kind.sigs.k8s.io/)
  * Run a local cluster
    * [minikube]  `minikube start`
    * [kind] `kind create cluster`


# Simple Pods
* `kubectl apply -f pod.yaml`
* `kubectl apply -f podSpecifyingAPIGroup.yaml`
  * WRONG declared pod, since 'core' API Group mustn't be specified!!
## Possible fields to update
* `metadata` fields can not be updated most of them
* if `metadata.deletionTimestamp` is set → no new entry to `metadata.finalizers` can be added
  * Problems:
    * Problem1: How to configure it? TODO: ?
* `spec.containers[*].image`
* `spec.initContainers[*].image` updated, but not added new
* `spec.activeDeadlineSeconds` can be updated by a smaller positive number
* `spec.tolerations` just adding them, not modifying exists ones

# Managed by a workload resource
## by a Job
* `kubectl apply -f podByWorkloadResourceJob.yaml`
* `kubectl get all`
* If you modify it and run `kubectl apply -f podByWorkloadResourceJob.yaml` ->
  * Problems:
    * Problem1: Which fields can be modified to be able to update it?
      * Solution: TODO: ?
## by a Deployment
* `kubectl apply -f podByWorkloadResourceDeployment.yaml`
* `kubectl get all`
* If you modify it and run `kubectl apply -f podByWorkloadResourceDeployment.yaml` ->
  * existing pods are terminated
  * new pods are created
## by a DaemonSet
* `kubectl apply -f podByWorkloadResourceDaemonSet.yaml`
* `kubectl get all`
* If you modify it and run `kubectl apply -f podByWorkloadResourceDaemonSet.yaml` ->
  * existing pods are terminated
  * new pods are created
 
# Static pods
## Create a static pod
* `docker ps`
  * Check the containerId of the cluster running locally
* `docker cp staticPod.yaml ContainerNameRunningCluster:/etc/kubernetes/manifests/`
  * They are created under kube-system
  * '/etc/kubernetes/manifests/' is the default directory where kubelet looks for static pod manifests
    * You can check it
      * `docker exec -it ContainerNameRunningCluster sh`
      * `cat /var/lib/kubelet/config.yaml` and check 'staticPodPath' field
  * `docker exec -it ContainerNameRunningCluster sh`
    * `cd /etc/kubernetes/manifests/` &  `ls`
      * Check that static pod manifest has been pasted
  * `kubectl get pods -n kube-system`
    * Check that static pod is running 
## Check kubelet is managing the static pod
* `docker exec -it ContainerNameRunningCluster sh`
* `systemctl status kubelet`
  * Check the logs, showing static pod references
## Check the existence of a mirror pod
* `kubectl get pods -n kube-system`
  * Check that there is a pod with naming 'nginx-static-pod...'
    * -> it's handled this one by Kubernetes API server
## Can not refer to other Kubernetes API objects
* `kubectl apply -f serviceAccount.yaml`
* `kubectl apply -f MirrorStaticPod -n kube-system`
* `docker exec -it ContainerNameRunningCluster sh`
  * 
  * `systemctl restart kubelet`
