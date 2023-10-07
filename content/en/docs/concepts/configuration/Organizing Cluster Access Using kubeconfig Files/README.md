
# Content
* `kubectl cluster-info`
  * Check the current cluster state
  * Problems:
    * Problem1: "The connection to the server <server-name:port> was refused - did you specify the right host or port?"
      * Solution: Cluster is not running. Run it: 
        * If you Docker Desktop -> Open and run the Kubernetes cluster
        * If you Rancher Desktop -> Open and run the Kubernetes cluster
        * If you use Minikube -> `minikube start`
        * If you use Kind -> `kind create cluster`
      * Notes: Check the Kubeconfig file -- by default in `~/.kube/config` --
    * Problem2: Get response, but you can't access to the cluster
      * `kubectl cluster-info dump` / `kubectl cluster-info dump > out.yaml`
        * Check cluster configuration