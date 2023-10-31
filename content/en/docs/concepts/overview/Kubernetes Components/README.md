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