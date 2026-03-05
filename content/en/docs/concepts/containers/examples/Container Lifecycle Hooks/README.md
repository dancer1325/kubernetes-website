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

----

# Container hooks
* `kubectl apply -f containerhooks.yaml`
  * Problems:
    * Problem1: 'Sidecar container is crashing continuously'
      * Attempt1: `command: [ "/bin/sh", "-c", "cp /var/tmp/prestop-executed.txt /shared-volume/", "sleep 3600" ]`
      * Solution: TODO:
  * Check that 'PostStart' hook handler has been executed
    * Attempt1: `kubectl describe pod/pod-containerhooks-1` -- no appear as pod's event --
    * Attempt2: `kubectl logs pod-containerhooks-1 -c containerhooks-1` 
    * Solution: `kubectl exec -it pod-containerhooks-1 -c containerhooks-1 -- /bin/sh` & `cat /usr/share/message`
  * Check that 'PreStop' hook handler has been executed
    * Force to kill the container
      * Attempt1: `kubectl exec -it pod-containerhooks-1  -c containerhooks-1  -- /bin/kill -s KILL 1`
      * Attempt2: `kubectl exec -it pod-containerhooks-1  -c containerhooks-1  -- pkill -9 -f init`
      * Attempt3: `kubectl exec -it pod-containerhooks-1  -c containerhooks-1  -- kill -9 1`
      * Attempt4: `kubectl exec -it pod-containerhooks-1  -c containerhooks-1  -- sh -c "kill -9 1"`
      * Solution: TODO:
    * Check the volume logs
      * TODO:


# Terminate a container
## ways
* API request
  * `kubectl delete pod/pod-terminate-1` 
  * There is no specific for container
* Management event
  * Failure of readinessProbe -- 'pod-terminate-2', Running the pod, but nothing READY --
  * Failure of livenessProbe -- 'pod-terminate-3', CrashLoopBackOff the pod, and nothing READY --
## `PodSpec.terminationGracePeriodSeconds`


# Hook handlers
* sync calls
  * Check 'pod-sync' & run fastly once you run the manifest `kubectl get pod/pod-sync -o wide`
* logs for hooks are NOT exposed in pod's events -- Check first section of this file --
* if a handler fails → it broadcasts an event
  * for PostStart → FailedPostStartHook event
    * Check 'pod-failedpoststarthook' & `kubectl describe pod/pod-failedpoststarthook`
  * for PreStop → FailedPreStopHook event
    * Check 'pod-failedprestophook', but how to kill to check afterward the event?
## types
### Exec
* Check 'pod-containerhooks-1' and 'pod-exec'
### HTTP
* Check 'pod-http'
### Sleep
* Requirements
  * kubectl v1.29
  * 'PodLifecycleSleepAction' enabled
    * `kind create cluster --config=clusterWithFeatureGates.yaml`
* Check 'pod-sleep'

# Container phases
## Waiting
* Check 'pod-waiting' & `kubectl get pod/pod-waiting -o json`
  * Problems: Find a correct pod sample
    * Solution: TODO:
## Running
* Command one, after all fine
## Terminated
* Check 'pod-terminating' & `kubectl get pod/pod-terminating -o json`

# Hook
* at least once
  * Check 'pod-failedpoststarthook'  -- `kubectl describe pod/pod-failedpoststarthook` the events --
