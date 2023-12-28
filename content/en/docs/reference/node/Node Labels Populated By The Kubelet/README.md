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

---
# Pre-populated labels on nodes
* Check that they are preset
  * `kubectl get nodes --show-labels`
    * 'node.kubernetes.io/instance-type', 'topology.kubernetes.io/region' & 'topology.kubernetes.io/zone' not preset since it's local ?
* Check that they are cloud-provider specific
  * Deploy a kubernetes cluster and check TODO:

# Custom own labels on nodes
* Via
  * kubelet configuration
    * `kind create cluster --config=clusterCustomizingKubelet.yaml` & `kubectl get nodes --show-labels` to check that a new label is added 
  * Kubernetes API
    * let's use kubectl (from all the possible ways)
      * `kubectl label node cluster-control-plane key2=value2` & `kubectl get nodes --show-labels` to check that a new label is added