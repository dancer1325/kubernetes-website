# Set all
* `kubectl create namespace namespace-a`
* `kubectl apply -f podA.yaml`
* `kubectl apply -f serviceA.yaml`
* `kubectl create namespace namespace-b`
* `kubectl apply -f podB.yaml`
* `kubectl apply -f serviceB.yaml`


# Steps
## Without specifying namespace in the DNS query
* `kubectl exec -it pod-a -n namespace-a -- sh`
  * Get access to the pod's shell
  * `nslookup service-b`
    * Namespace not specified in the dns query -> own pod's namespace (namespace-a) -> not resolved service-b
    * DNS Server tries to resolve small variations of the name, but without succeed


## Specifying namespace in the DNS query
* `kubectl exec -it pod-a -n namespace-a -- sh`
  * Get access to the pod's shell
  * `nslookup service-b.namespace-b.svc.cluster.local`
    * Using the full domain name == specify the namespace -> resolved service-b


## Specifying resolv.conf
* `search domain1.com domain2.com domain3.com`
  * Structure of search field
*  `kubectl exec -it pod-a -n namespace-a -- sh`
  * `cd /etc`
  * adjust search row
    * `grep -v 'search namespace-a.svc.cluster.local svc.cluster.local cluster.local' resolv.conf | tee resolv.co
      nf`
      * delete search row
    * `echo "search namespace-a.svc.cluster.local svc.cluster.local cluster.local namespace-b.svc.cluster.local"
      | tee -a resolv.conf`
      * add search entry complete, including namespace-b domain
  * `nslookup service-b`
    * other service is resolved, since the default namespace-b is added 