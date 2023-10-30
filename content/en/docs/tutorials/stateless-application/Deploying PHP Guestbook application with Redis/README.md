# Pre requisites
* Locally
  * Run some docker daemon
    * [Docker desktop](https://www.docker.com/products/docker-desktop/)
  * Install some local cluster
    * tool
      * [minikube](https://minikube.sigs.k8s.io/docs/start/)
      * [kind](https://kind.sigs.k8s.io/)
    * recommendations
      * 2 nodes or more
  * curl
  * kubectl > v1.14
    * `kubectl get nodes -o=jsonpath=$'{range .items[*]}{@.metadata.name}: {@.status.nodeInfo.kubeletVersion}\n{end}'`
      * Check it in the current context's nodes


# Steps
* `kubectl apply -f https://k8s.io/examples/application/guestbook/redis-leader-deployment.yaml`
  * It's the same placed under 'examples/' in this repo 
* `kubectl apply -f https://k8s.io/examples/application/guestbook/redis-leader-service.yaml`
  * It's the same placed under 'examples/' in this repo 
  * Service to proxy traffic from Guestbook application -> Redis leader pod
  * Type isn't specified -> ClusterIP type -> only accessible within the Kubernetes cluster!! 
* `kubectl apply -f https://k8s.io/examples/application/guestbook/redis-follower-deployment.yaml`
  * It's the same placed under 'examples/' in this repo
  * Add pod followers to the leader pod
* `kubectl apply -f https://k8s.io/examples/application/guestbook/redis-follower-service.yaml`
  * It's the same placed under 'examples/' in this repo
  * Service to proxy traffic from Guestbook application -> Redis follower pod
  * Type isn't specified -> ClusterIP type -> only accessible within the Kubernetes cluster!!
* `kubectl apply -f https://k8s.io/examples/application/guestbook/frontend-deployment.yaml`
  * It's the same placed under 'examples/' in this repo
  * Guestbook php web server, configured to communicate with Redis 
    * if READ request -- communicat with -> Redis follower
    * if WRITE request -- communicat with -> Redis leader
  * `kubectl get pods -l app=guestbook -l tier=frontend`
    * Check that all 3 pods are running
* Create the frontend service
  * ClusterIP type
    * `kubectl apply -f https://k8s.io/examples/application/guestbook/frontend-service.yaml`
      * It's the same placed under 'examples/' in this repo
      * Type isn't specified -> ClusterIP type -> only accessible within the Kubernetes cluster!!
      * Ways to expose a service of ClusterIP type outside
        * Kubernetes Gateway API
        * Ingress
        * `kubectl port-forward serviceName localPort:RemotePortInThePod` 
          * expose the service outside the cluster
          * `kubectl port-forward svc/frontend 8080:80`
          * Open in your browser 'http://localhost:8080/' to check that you get access to the application running in a cluster
  * LoadBalancer type
    * Go to the 'examples/' and uncomment the line to specify the type
    * `kubectl apply -f frontend-service.yaml`
    * `kubectl get services/frontend`
    * Open in your browser 'http://localhost' to check that you get access to the application running in a cluster
* `kubectl scale deployment frontend --replicas=5`
  * Scale up/down the pods, since you use a Deployment controller