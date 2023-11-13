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

# Kubernetes Cluster
* `kind create cluster --config=kindClusterWithSeveralNodes.yaml`
  * Problems: 
    * Problem1: "failed to create cluster: node(s) already exist for a cluster with the name "kind"
      * Solution: `docker ps` and check the containerName of kind process. `docker kill ContainerNameOfKindContainer`. Or
  * `kubectl get nodes`
    * cluster is created with 3 nodes
      * 1 control plane == master node 
      * 2 worker nodes == worker node
    * Problems:
      * Problem1: Why the nodes declared as worker don't appear as worker
        * Attempt1: `kind create cluster --config=kindClusterWithoutSpecifyingImage.yaml`
        * Attempt2: `kind create cluster --config=kindClusterWithSeveralNodes.yaml`
* Check 'tutorials/' to realize that any containerized application can be run there


# TODO

# Node Components

## kubelet

## kube-proxy
* `kubectl describe pod/KubeProxyPodId -n kube-system`
  * Check detailed information about the kube-proxy pod related
    * which lives in the namespace kube-system
* `kubectl logs -n kube-system -l k8s-app=kube-proxy`
  * Check logs
* `kubectl describe configmap kube-proxy -n kube-system`
  * Check the kube-proxy configuration served in a configmap
  * `kubectl get configmap kube-proxy -n kube-system -o json | jq '.data' | jq -r '.["config.conf"]'`
    * If you want to extract the KubeProxyConfiguration