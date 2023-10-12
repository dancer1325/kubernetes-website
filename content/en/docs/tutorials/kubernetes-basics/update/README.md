# Note
* Use the Deployment created in "/deploy-app"

# Content
* `kubectl describe pods`
  * Check detailed information about the pods
    * Pay attention to the 'image' field
* `kubectl set image deployments/DeploymentName containerName=NewImageName:vNumberWithTheNewVersion`
  * Update the container's image version
    * Example:
      * [jocatalin/kubernetes-bootcamp:v2](https://hub.docker.com/layers/jocatalin/kubernetes-bootcamp/v2/images/sha256-fb1a3ced00cecfc1f83f18ab5cd14199e30adc1b49aa4244f5d65ad3f5feb2a5?context=explore)
  * If the 'NewImageName:vNumberWithTheNewVersion' in the registry -> Roll update
    * Example:
      * gcr.io/google-samples/kubernetes-bootcamp:v10
    * `kubectl rollout undo deployments/DeploymentName`
      * Revert the deployment
  * `kubectl describe pods`
    * Check that the 'image' field has been changed
  * `kubectl expose deployment/DeploymentName --type="TypeOfService" --port PortNumber`
    * Create a service to expose the deployment
    * Check '/expose' folder
  * `curl http://"$(minikube ip):$NODE_PORT"`
    * Ping the applications from outside the cluster, through the service
    * Check '/expose' folder
      * Pay attention to the problem found for certain set up
    * Load balancer is working
      * Different pod is hit / request
* `kubectl rollout status deployments/DeploymentName`
  * Display the rollout status
  