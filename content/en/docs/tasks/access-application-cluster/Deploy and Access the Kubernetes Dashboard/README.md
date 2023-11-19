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


# Set up dashboard UI
* `kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml`
  * create all the resources in an own namespace
* `kubectl apply -f serviceAccountDashboard.yaml`
* `kubectl apply -f clusterRoleBindingDashboard.yaml`
* `kubectl -n kubernetes-dashboard create token admin-user`
  * Create a token that it's printed
  * Copy the value and access to the Dashboard in 'http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/'