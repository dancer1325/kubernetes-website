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


# := diagnostic performed periodically on a container by kubelet
* `kubectl apply -f pods.yaml` & `docker exec -it ControlPlaneContainer journalctl -u kubelet | grep 'done-by-kubelet'`
  * check that it's watched
    * TODO: Check how to confirm that the diagnostic is indeed done by kubelet

# mechanisms available
* 1! mechanism / probe -- 'one-uniquemechanism-per-probe'
  * `kubectl apply -f pods.yaml` & check how we can NOT create that pod
## exec
* == execute a command within the container -- 'mechanism-exec' --
  * `kubectl apply -f pods.yaml` & `kubectl logs po/mechanism-exec`
    * Problems: "Why do NOT I see 'Read' as logs"
  * if the command exits with statusCode 0 == successful
    * Problems: "How to confirm that statusCode 0"
      * Attempt1: `kubectl get events --field-selector involvedObject.name=mechanism-exec`
      * Attempt2: `kubectl describe pod/mechanism-exec`
  * create / forks multiple processes
    * `kubectl exec -it pod/mechanism-exec sh` &
      * Attempt1: `ps aux` or `strace -f -e trace=execve echo`but tools NOT supported
      * Attempt2: `kubectl describe pod/mechanism-exec`
      * Solution: TODO:
## grpc
## httpGet
## tcpSocker