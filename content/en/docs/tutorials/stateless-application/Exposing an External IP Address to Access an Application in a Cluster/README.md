# Pre requisites
* Locally 
  * Run some docker daemon
    * [Docker desktop](https://www.docker.com/products/docker-desktop/)
  * Install some local cluster tool
    * [minikube](https://minikube.sigs.k8s.io/docs/start/)
    * [kind](https://kind.sigs.k8s.io/)
  * curl
  * kubectl > v1.28
    * `kubectl get nodes -o=jsonpath=$'{range .items[*]}{@.metadata.name}: {@.status.nodeInfo.kubeletVersion}\n{end}'`
      * Check it in the current context's nodes
* Cloud provider to create
  * Kubernetes cluster
  * external Load Balancer

# Steps
* `kubectl apply -f deployment.yaml`
  * Create the deployment and replicaset
* `kubectl expose deployment hello-world --type=LoadBalancer --name=my-service`
  * Create a Service of type LoadBalancer to expose the deployment
  * `kubectl get services` or `kubectl describe service my-service`
    * Check the service created and take note about
      * 'EXTERNAL-IP' column or 'LoadBalancer Ingress' and
      * 'PORT' column with the format '<NodeIP>:<NodePort>' or 'Port' and 'NodePort'
    * If 'EXTERNAL-IP' = localhost == service hosted locally
    * If you want to provide a service hosted in a cloud provider -> TODO
* `curl http://LoadBalancerIngress:PortValue`
  * Check that you get the output of the endpoint of the application running in a cluster