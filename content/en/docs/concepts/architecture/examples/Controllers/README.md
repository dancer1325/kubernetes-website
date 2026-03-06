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


# Job controller
* `kubectl apply -f job.yaml`
  * Create a Kubernetes Job Controller based on the manifest (:= file)
* `kubectl get job example-job`
  * Check job status. After a time, the job is indicated as completed
* `kubectl get pods`
  * Check the pods created
  * After the job is scheduled (=== completed) -> pod is completed but not running currently
    * Reason: job's succeed is scheduling ===  pod matches with the node (!= running)
* `kubectl logs example-job-s8qkx`
  * Check the logs after job is completed
* `kubectl logs -n kube-system kube-controller-manager-minikube | grep example-job`
  * Display logs about the Job Controller, confirming that it runs inside the Control Plane's controller-manager

# Deployment controller
* `kubectl apply -f deployment.yaml`
  * Create a Kubernetes Deployment Controller based on the manifest (:= file)
* `kubectl get deployment example-deployment`
  * Check deployment status. After a time, the deployment is indicated as READY
* `kubectl get pods`
  * Check the pods created
  * After the deployment is completed (!= scheduled, just for job) -> pod is running currently
* `kubectl logs -n kube-system kube-controller-manager-minikube | grep example-deployment`
  * Display logs about the Deployment Controller, confirming that it runs inside the Control Plane's controller-manager

# Direct control
* Check https://github.com/kubernetes/autoscaler/tree/master
