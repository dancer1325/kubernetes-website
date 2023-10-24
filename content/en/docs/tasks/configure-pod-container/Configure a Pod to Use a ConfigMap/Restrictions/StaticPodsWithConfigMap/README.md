# Pre requirements
* Check 'Create static pods/'

# How to run locally?
* Connect to your cluster and create the static pod
  * For
    * minikube -> `minikube ssh`
  * Rest of command run inside the cluster
    * `mkdir -p /etc/kubernetes/manifests/`
      * folder monitored by kubelet on the node
    * `cd /etc/kubernetes/manifests/`
    * `sudo chown -r $USER:$USER .`
      * Give rights to your current user
    * Code line by line of 'commands.txt', 1.
  * Run outside the cluster
    * `kubectl get pods`
      * Check that 'static-nginx' pod exist
* Outside the cluster, create the configmap
  * ` kubectl create -f configmap.yaml`
* Create another static pod as first one, but trying to consume the configmap
  * For
    * minikube -> `minikube ssh`
  * Rest of command run inside the cluster
    * `cd /etc/kubernetes/manifests/`
    * Code line by line of 'commands.txt', 2.
  * Run outside the cluster
    * `kubectl get pods`
      * Check that 'static-configmap-nginx' pod NOT exist