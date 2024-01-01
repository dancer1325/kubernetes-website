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

# Quality of Service Classes (QoS)
* is assigned to each pod -- 'dummy' --
  * `kubectl apply -f pods.yaml` & `kubectl describe pod/dummy` and check that there is a field named 'QoS Class'
## types
### Guaranteed
* [Requirements] all pod's containers must have: 1.defined memory limit & memory request, 2. memory limit == memory request, 3. defined CPU limit & CPU request, 4. CPU limit == CPU request -- 'guaranteed' --
  * `kubectl apply -f pods.yaml` & `kubectl describe pod/guaranteed` to confirm that the QoS class is 'Guaranteed'
* can use static policy
  * TODO:
### Burstable
* [Requirements] 1. NOT meet the Guaranteed requirements, & at least 1 pod's container must have: 1.defined memory limit, or 2. defined memory request, or 3. defined CPU limit, or 4. defined CPU request -- 'burstable-1' & 'burstable-2' --
  * `kubectl apply -f pods.yaml` & `kubectl describe pod/burstable-1` & `kubectl describe pod/burstable-2` to confirm that the QoS class is 'Burstable'
* if a limit (memory or CPU) NOT specified → related to the capacity of the node -- 'burstable-2' --
  * Attempt1: `kubectl describe pod/burstable-2`
  * Attempt2: `kubectl top pod burstable-2` not enabled metrics API
  * Attempt3: `http://NodeIp:4194` but not enabled the CAdvisor UI
  * Solution: TODO:
* if there are containers without specifying requests or limits of resources -> try to use any amount of node resources
  * TODO:
### BestEffort
* [Requirements] 1. NOT meet the Guaranteed nor Burstable requirements, & none of the pod's container must have: 1.defined memory limit, nor 2. defined memory request, nor 3. defined CPU limit, nor 4. defined CPU request -- 'besteffort-1' --
  * `kubectl apply -f pods.yaml` & `kubectl describe pod/besteffort-1` to confirm that the QoS class is 'BestEffort'
* can use node resources which are NOT assigned specifically to pods in other QoS classes
  * TODO:
## about eviction
* if there are not enough resources on a node -> order to evict are 'BestEffort', 'Burstable' and 'Guaranteed' pods 
  * Attempt1: -- 'cause-evict' -- `kubectl apply -f podToCauseEviction.yaml`, but it doesn't kill 'BestEffort' pods to run this 'Burstable' one
  * TODO: 
* if there is a resource pressure → just pods exceeding resources requests are candidates for eviction
  * TODO:

# Memory QoS with cgroup v2
* TODO:

# Behaviors independent of the QoS
* if a container exceeds a resource limit →  1. killed, 2. restarted by kubelet without affecting other pod’s containers
  * TODO:
* if a container exceeds its resource request & node faces node resource pressure → 1. pod becomes a candidate for eviction, 2. all pod’s containers will be terminated, 3. Kubernetes may create a replacement pod, usually on a different node
  * TODO:
* pod's resource request = Sum all containers resource requests, pod's resource limit = Sum all containers resource limits
  * TODO:
* Pods to preempt are NOT based on QoS 
  * Check '../Scheduling, Preemption and Eviction/Pod Priority and Preemption'