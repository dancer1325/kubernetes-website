# Note
* Use the Deployment created in "/deploy-app"

# Content
* `kubectl get pods` / `kubectl get services` 
  * If you use locally minikube -> a Service called 'Kubernetes' is created by default, if you start minikube cluster
  * `... -o json|yaml|name|go-template|...`
    * output format
    * [link](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#get)
* `kubectl expose deployment/DeploymentName --type="TypeOfService" --port PortNumber`
  * Create a Service to expose the deployment related objects
  * If TypeOfService == NodePort
    * unique clusterIP
    * <NodeIP>:<NodePort>
* `kubectl describe services/ServiceName`
  * Check detailed information about the service
* `export NODE_PORT="$(kubectl get services/ServiceName -o go-template='{{(index .spec.ports 0).nodePort}}')"`
  * Create an environment variable with the nodePort
* `curl http://"$(minikube ip):$NODE_PORT"`
  * Check application running in a container is exposed outside the cluster
  * Problems:
    * Problem1: Time Out
      * Note: Set up it minikube cluster with Docker Desktop as container driver
      * Solution: `minikube service kubernetes-bootcamp --url` and with the output `curl PreviousOutput`
        * Reason: Docker Desktop are isolated from your host computer -> you need a tunnel
* `kubectl describe deployment`
  * Check detailed information about the deployment
    * Default label is added
* `kubectl get pods,svc -l LabelAdded`  
  * Query pods and service by label
* `export POD_NAME="$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')"`
  * Create an environment variable with the podName
  * `kubectl get pods -o json > out.json`
    * Write in a file the output of the query
* `kubectl label pods PodName key=value`
  * Apply a label to the pod
  * `kubectl get pods -l key=value`
    * Check that the label has been applied
* `kubectl delete service -l key=value`
  * Delete a service, filtering by a label
  * `curl http://"$(minikube ip):$NODE_PORT"`
    * Check that the application is not reachable outside the cluster
  * `kubectl exec -ti $POD_NAME -- curl http://localhost:8080`
    * Check that the application is still running but only reachable inside the pod