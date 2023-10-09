# Deploy an app
* `kubectl create deployment kubernetes-bootcamp --image=gcr.io/google-samples/kubernetes-bootcamp:v1`
  * If the image is hosted outside Docker Hub -> full repository url
    * Example: "gcr.io/google-samples/kubernetes-bootcamp:v1"
  * What was done behind?
    * search for a suitable node where to run an instance of the application
    * schedule the application to run in that node
    * configure the cluster to reschedule the instance on a new node when it's needed
* `kubectl get deployments`
  * Check the deployments'
    * status
    * number of instances of an application

# View the app
* `kubectl proxy [--port=PORT] [--www=static-dir] [--www-prefix=prefix] [--api-prefix=prefix]`
  * Create a proxy or application-level gateway between localhost ← → Kubernetes API server
  * `PORT`
    * (By default) 8001
      * `curl http://localhost:8001/version` Check that it's equal to use kubectl
  * 1 endpoint / pod is created
    * `export POD_NAME=$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` and `echo Name of the Pod: $POD_NAME`
      * Get the pod name and storing it in a variable
    * `curl http://localhost:8001/api/v1/namespaces/default/pods/$POD_NAME/`
      * Check pod information