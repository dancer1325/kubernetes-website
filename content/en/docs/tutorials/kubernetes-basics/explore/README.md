# Note
* Use the Deployment created in "/deploy-app"

# Content
* `kubectl get pods`
* `kubectl describe pods PodName`
  * Check pod's container details
    * IP address
    * ports used
    * List of events related to the lifecycle of the pod
* `kubectl logs PodName`
  * Check PodName pod's logs
* `kubectl exec it PodName -- SomeCommandToSpecify`
  * `SomeCommandToSpecify` can be
    * `bash`
      * Bash session is started in the pod