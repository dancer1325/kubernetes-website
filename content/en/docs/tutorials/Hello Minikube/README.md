
# Create a deployment
* `kubectl create deployment hello-node --image=registry.k8s.io/e2e-test-images/agnhost:2.39 -- /agnhost netexec --http-port=8080`
  * `kubectl create deployment DeploymentName`
    * create a deployment with name "DeploymentName"
  * "registry.k8s.io"  TODO: Which one is the concrete url for ?
  * "e2e-test-images/agnhost:2.39"
    * [DockerHub](https://hub.docker.com/layers/opsdockerimage/e2e-test-images-agnhost/2.39/images/sha256-93c166faf53dba3c9c4227e2663ec1247e2a9a193d7b59eddd15244a3e331c3e)
      * === application packaged in a Docker container
    * [agnhost](https://pkg.go.dev/k8s.io/kubernetes/test/images/agnhost#section-readme)
      * agnostic - host
      * Extendable CLI which outputs and behaves same in any underlying OS
      * `netexec`
        * start a HTTP server with the [following endpoints](https://pkg.go.dev/k8s.io/kubernetes/test/images/agnhost#readme-netexec)
* `kubectl get deployments`
  * Check the deployments
* `kubectl get pods`
  * Check the pods
* `kubectl get events`
  * Check cluster events
* `kubectl config view`
  * Check Kubernetes config file
* `kubectl logs PodName`
  * Check application logs for the pod’s containers

# Create a service
* Pod’s accessibility
  * (By default) it’s — only accessible via — internal IP address within Kubernetes cluster
* `kubectl expose deployment DeploymentName --type=LoadBalancer --port=8080`
  * Create a service of type LoadBalancer, exposing the port in which the application agnhost is listening
    * Since no load balancer is provided -> no external IP address for the service
* `kubectl get service`
  * Check the existing services
* `minikube service DeploymentName`
  * Get access to the services of LoadBalancer type

# Enabled addons
* `minikube addons list`
  * supported Minikube addons
* `minikube addons enable AddonName`
  * enable an addon
  * Pod and services are bootstrap in the kube-system namespace!!
    * `kubectl get pod,svc -n kube-system`
      * Check pods and services in kube-system namespace
* `minikube addons disable AddonName`
  * disable an addon

# Clean up
* `kubectl delete service hello-node`
* `kubectl delete deployment hello-node`