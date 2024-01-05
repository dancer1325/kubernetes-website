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

---
# Debug Pods
## Check state of the pod & pod's containers
* If pod stays in 'Pending' -- 'pending-1' & 'pending-2' --
  * `kubectl apply -f pods.yaml` 
  * Possible reasons  
    * insufficient resources      -- 'pending-1' --
      * `kubectl describe pod/pending-1` specifying 'Insufficient cpu' in this case
    * you are binding a pod → hostPort    -- 'pending-2-1' & 'pending-2-2' & 'pending-2-3' --
      * `kind create cluster --config=cluster.yaml`
      * `kubectl describe pod/pending-2-1` & `kubectl describe pod/pending-2-2` are 'Running', placing hostPort 8080  & `kubectl describe pod/pending-2-3` is 'Pending'
        * Problems: Why no pod 'pending-2' are running? Not all -1 should be running
          * Solution: TODO:
        * Events in 'pending-2-3' specify 'didn't have free ports for the requested pod ports'
        * number of pods to schedule = number of nodes
* If pod stays in 'Waiting' -- 'waiting-1' --
  * `kubectl apply -f pods.yaml`
  * Possible reasons
    * failure to pull the image -- 'waiting-1' --
      * `kubectl describe pod/waiting-1` to check the pod's Status
        * Problems: pod's status displays ' Pending', but containers' state is 'Waiting' with Reason 'ErrImagePull'
          * Solution: That's correct? TODO:
* If pod stays in 'Terminating'  -- 'terminating-'
  * `kubectl apply -f pods.yaml`
  * Possible reasons
    * check if the cluster has
      * ValidatingWebhookConfiguration -- '' --
        * TODO:
      * MutatingWebhookConfiguration  -- '' --
        * TODO:
    * Check if you use a third-party webhook that -- '' --
      * TODO:
    * Check if you use your own webhook that -- '' --
      * TODO:
* If pod is crashing or unhealthy  
  * Check 'Debug Running Pods/'
* If pods are running, but not doing what I told it -- 'running-not-doing-what-i-told' --
  * `kubectl apply -f pods.yaml`
  * Possible reasons
    * errors in the podSpec & ignored during pod creation
      * Such as? TODO:
## Check events
* `kubectl describe pod/PodName`

# Debug Replication Controller
* if they can NOT create pods → Check debug pods      -- 'replication-controller' --
  * `kubectl apply -f replicationController.yaml` & `kubectl describe rc replication-controller` but not relevant message detailing the root cause
  * `kubectl describe pod/PodName` as common debug pod to get detailed information about the root cause
