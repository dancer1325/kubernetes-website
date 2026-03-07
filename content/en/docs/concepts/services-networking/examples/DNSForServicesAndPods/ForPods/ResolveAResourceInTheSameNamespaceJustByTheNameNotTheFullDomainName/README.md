
# Steps
* `kubectl create namespace my-namespace`
  * Create own namespace
* `kubectl apply -f pod.yaml`
* `kubectl apply -f service.yaml`
* `kubectl exec -it my-pod -n my-namespace -- sh`
  * Get access to the pod's shell
  * `nslookup my-service`
    * Resolve the service just by the domain, without using the full domain name
    * DNS Server tries to resolve small variations of the name, but without succeed
  * `nslookup my-service.my-namespace.svc.cluster.local`
    * Resolve the service by the full domain name