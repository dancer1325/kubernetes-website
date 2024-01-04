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
  * Steps to install 'metrics-server'
    * `kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml`
    * `kubectl patch -n kube-system deployment metrics-server --type=json \
      -p '[{"op":"add","path":"/spec/template/spec/containers/0/args/-","value":"--kubelet-insecure-tls"}]'`

---

# Ways to collect monitoring statistics
## Resource Metrics pipeline
* metrics related to
  * `kubectl apply -f pods.yaml`
  * HPA
    * TODO:
  * `kubectl top`
    * Problem: "error: Metrics API not available"
      * Attempt1: `kubectl top pods`
      * Solution: Check how to make run 'metrics-server'
    * `kubectl top pods` to confirm that it returns CPU and Memory of pods running in 'default' namespace
      * Note: It could take some time to extract statistics of pods
* Collected by 'metrics-server'
  * Follow the initial steps to make it run!!
  * How to check that it's collected by it? TODO:
  * allows
    * discovering all cluster's nodes
      * `kind create cluster --config=cluster.yaml` & Set 'metrics-server' & `kubectl apply -f pods.yaml`  & `kubectl get pods -o wide` displaying pods in different cluster's nodes
      * `kubectl top pods` displaying metrics from pods living in all cluster's nodes
        * Problems: Why doesn't it display statistics from pods running in control plane?
          * Solution: Apparently, now it's working fine
    * querying all node’s kubelet
      * `docker exec -it NodeId journalctl -u kubelet | grep 'metrics-server'` to confirm that metrics-server query to all nodes' kubelet
* exposed via 'Metrics' API in `metrics.k8s.io`
  * `kubectl get --raw /apis/metrics.k8s.io/v1beta1` to confirm that it returns an object
## Full Metrics pipeline
* TODO: 

# Kubernetes
* designed to work with OpenMetrics
  * How to check? TODO:
