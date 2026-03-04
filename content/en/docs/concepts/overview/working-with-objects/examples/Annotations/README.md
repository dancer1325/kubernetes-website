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

# := metadata which
* ``{
  "key1" : "value1",
  "key2" : "value2",
  …
  }``
  * 👁`key` & `value` just can be string 👁️
  * key’s syntax — `optionalPrefix/name`
    * `optionalPrefix`
      * `kubernetes.io` & `k8s.io` are allowed, although they are reserved to Kubernetes core components
    * `name`
      * mandatory
      * - _ .     are invalid characters to start or ending the name
* uses
  * attach to Kubernetes' objects -- 'pods.yaml'
* != labels
  * uses
  * annotations have less restrictions
* + labels / same object -- can be used at the same time

# Note
* `kubectl apply -f pod.yaml`   to confirm all