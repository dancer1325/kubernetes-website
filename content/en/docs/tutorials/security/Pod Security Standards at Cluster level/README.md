# Pre requisites
* Locally or cluster you fully control (=== configure control plane), where you can
  * Run some docker daemon
    * [Docker desktop](https://www.docker.com/products/docker-desktop/)
  * Install some local cluster tool
    * [minikube](https://minikube.sigs.k8s.io/docs/start/)
    * [kind](https://kind.sigs.k8s.io/)
  * kubectl > v1.25

# Via kind
* `kind create cluster --name psa-wo-cluster-pss`
  * Create a cluster with no Pod Security Standards
* `kubectl config current-context`
  * Check the current Kubernetes context
  * If it's not the current kind one -> `kubectl cluster-info --context kind-psa-wo-cluster-pss`
* `kubectl get namespace`
  * Check the current active namespaces, created by default by kind
* `kubectl label --dry-run=server --overwrite ns --all \
  pod-security.kubernetes.io/enforce=privileged`
  * `kubectl label`
    * add / update / remove label
  * `--dry-run`
    * testing or validation purposes, without applying changes to the cluster!!
    * `kubectl get namespace NamespaceName --show-labels`
      * You can check that the labels have not been applied
    * `=server`
      * validated by the server
      * output displayed running the command
  * `--overwrite`
    * if the label already exists -> it must be overwritten
  * `pod-security.kubernetes.io/enforce=privileged` or `pod-security.kubernetes.io/enforce=baseline` or `pod-security.kubernetes.io/enforce=restricted`
    * key-valur pair to set as label
    * Pod Security Standards
      * 'privileged'
      * 'baseline' -- some warnings are get for certain namespaces --
      * 'restricted' -- some warnings are get for certain namespaces --
    * Pod Security Admission modes
      * 'enforce'
      * 'audit'
      * 'warn'

---

* `kind create cluster --name psa-with-cluster-pss --config tmp/pss/cluster-config.yaml`
  * Create a cluster with custom PSS
  * Problems
    * Problem1: Time out, starting the control plane
      * Solution: If you are using Docker Daemon -> Add the folder containing these .yaml to 'Preferences.Resources.FileSharing'
* `kubectl config current-context`
  * Check the current Kubernetes context
* `kubectl cluster-info --context kind-psa-with-cluster-pss`
  * Point kubectl to the recent created cluster
* `kubectl apply -f https://k8s.io/examples/security/example-baseline-pod.yaml`
  * Pod.yaml got it from this documentation
  * Output displays that the pod is started but with a warning


# Via minikube
* TODO: 
