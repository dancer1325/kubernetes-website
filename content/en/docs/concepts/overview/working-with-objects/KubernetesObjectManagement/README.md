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

# Imperative commands
* Operates on live Kubernetes Objects
  * Example: `kubectl create deployment nginx --image nginx`


# Imperative object configuration
* Operates on individual files
  * Example:
    * `kubectl create -f pod.yaml` / `kubectl create -f pod.yaml -f pod2.yaml` 
    * `kubectl delete -f pod.yaml`


# Declarative object configuration 
* Operates on directories of files, creating or patching
  * Example:
    * `kubectl diff -f declarativeObjectConfig`
    * `kubectl apply -f declarativeObjectConfig`