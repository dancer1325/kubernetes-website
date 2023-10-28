# Pre requisites
* Locally or cluster you fully control (=== configure control plane), where you can
  * Run some docker daemon
    * [Docker desktop](https://www.docker.com/products/docker-desktop/)
  * Install some local cluster tool
    * [minikube](https://minikube.sigs.k8s.io/docs/start/)
    * [kind](https://kind.sigs.k8s.io/)
  * kubectl > v1.25

# Via kind
* `kind create cluster --name psa-ns-level`
  * Create a cluster with no Pod Security Standards
* `kubectl config current-context`
  * Check the current Kubernetes context
  * If it's not the current kind one -> `kubectl cluster-info --context kind-psa-ns-level`
* `kubectl create ns example`
  * Create a namespace
* `kubectl label --overwrite ns example \
  pod-security.kubernetes.io/enforce=baseline \
  pod-security.kubernetes.io/enforce-version=latest \
  pod-security.kubernetes.io/warn=restricted \
  pod-security.kubernetes.io/warn-version=latest \
  pod-security.kubernetes.io/audit=restricted \
  pod-security.kubernetes.io/audit-version=latest`
  * `kubectl label`
    * add / update / remove label
  * `--overwrite`
    * if the label already exists -> it must be overwritten
  * `pod-security.kubernetes.io/enforce=baseline` or `pod-security.kubernetes.io/enforce-version=latest`
    * key-valur pair to set as label
    * baseline PSS enforced 
  * `pod-security.kubernetes.io/warn=restricted` or `pod-security.kubernetes.io/warn-version=latest`
    * key-valur pair to set as label
    * restricted PSS warned
  * `pod-security.kubernetes.io/audit=restricted` or `pod-security.kubernetes.io/audit-version=latest`
    * key-valur pair to set as label
    * restricted PSS restricted
* `kubectl get namespace NamespaceName --show-labels`
* `kubectl apply -n example -f https://k8s.io/examples/security/example-baseline-pod.yaml`
  * Pod.yaml got it from this documentation
  * Output displays that the pod is started but with a warning
    * === baseline pod created in the example namespace
* `kubectl apply -n default -f https://k8s.io/examples/security/example-baseline-pod.yaml`
  * Same pod, but created in the default namespace
  * Since nor same PSS nor modes in the default namespace ->  no warnings got

---

* `kind delete cluster --name psa-ns-level`
  * Delete the previous cluster

# Via minikube
* TODO: 
