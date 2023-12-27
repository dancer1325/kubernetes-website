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
# Scheduler
* Check that a node / pod 
  * `kubectl apply -f pod.yaml` & `kubectl get pods -o wide` displaying the column 'NODE'
* If the scheduler doesn't find a feasible node → pod remains unscheduled in 'Pending' status
  * Check 'pod-withoutfeasiblenode'
  * `kubectl apply -f pod.yaml` & 
  * `kubectl get pods -o wide` displaying in the column 'NODE' empty and 'Pending' in the column 'STATUS' & 
  * `kubectl describe pod/pod-withoutfeasiblenode` to check the events

# kube-scheduler
* Check 'filtering' & 'scoring' processes logs
  * Attempt1: `kind create cluster --verbosity 6` -- Check [here](https://github.com/kubernetes-sigs/kind/blob/main/hack/ci/e2e-k8s.sh#L96)
  * Attempt2: `kind create cluster --config=cluster.yaml`
  * Solution: TODO:
* Check binding notification from 'kube-scheduler' to 'api server'
  * Solution: TODO: