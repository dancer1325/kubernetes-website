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

# TODO:

# Headless Services
* Without selectors
  * `kubectl apply -f Headless\ Services/withoutSelectors.yaml`
    * Create Headless Service
  * `kubectl get endpoints headlessservice-withoutselectors`
    * EndpointSlices not created
  * `kubectl apply -f Headless\ Services/dnslookuppod.yaml`
    * Deploy the dnslookup pod
    * `kubectl get pods`
      * Check the podName
    * `kubectl exec -it dns-lookup-pod -- sh`
      * commands to execute into this pod
      * `nslookup example.com`
        * check IPv4 address
        * since in the output.Addresses contains IP address -> A record is returned
      * `nslookup -query=AAAA example.com`
        * check IPv6 address
        * if there is output.Addresses -> AAAA record is returned
  * TODO: Build a sample, showing incompatibilities / errors adding targetPort != Port 
* With selectors
  * `kubectl apply -f Headless\ Services/withSelectors.yaml`
    * Create Headless Service
  * `kubectl get endpoints headlessservice-withselectors`
    * EndpointSlices created
  * `kubectl apply -f Headless\ Services/dnslookuppod.yaml`
    * Deploy the dnslookup pod
    * `kubectl get pods`
      * Check the podName
    * `kubectl exec -it dns-lookup-pod -- sh`
      * commands to execute into this pod
      * `nslookup example.com`
        * check IPv4 address
        * since in the output.Addresses contains IP address -> A record is returned
      * `nslookup -query=AAAA example.com`
        * check IPv6 address
        * if there is output.Addresses -> AAAA record is returned
* `kubectl get services`
  * Check that the 'CLUSTER-IP' column is 'NONE'
