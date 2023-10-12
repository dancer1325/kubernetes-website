# Note
* Use the Deployment created in "/deploy-app"

# Content
* `kubectl get deployments`
* `kubectl get replicaset` / `kubectl get rs`
* `kubectl scale deployments/DeploymentName —replicas=NumberOfReplicasToScale`
  * Scale the deployment to the desired number of replicas
  * `kubectl get all`
    * Check that there are the number of pods indicated
  * `kubectl get pods -o wide`
    * Get more detailed information about the existing pods
    * Each pod has different IP address !!
  * `kubectl describe deployments/kubernetes-bootcamp`
    * If you check 'events' information -> scaling happened
* `kubectl expose deployment/DeploymentName --type="TypeOfService" --port PortNumber`
  * Create a service to expose the deployment
  * Check '/expose' folder
* `curl http://"$(minikube ip):$NODE_PORT"`
  * Ping the applications from outside the cluster, through the service
  * Check '/expose' folder
    * Pay attention to the problem found for certain set up
  * Load balancer is working
    * Different pod is hit / request
* `kubectl scale deployments/DeploymentName --replicas=NumberOfReplicasDesired`
  * Adjust the number of replicas
  