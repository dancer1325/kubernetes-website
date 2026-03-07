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

# Uses
* Network application running in a pod in your cluster → it's exposed
  * `kubectl apply -f simplePod.yaml`
  * `kubectl apply -f simpleService.yaml`
    * Check the manifest
    * Since this Service is type 'nodePort' -> you can reach outside the cluster!
      * `kubectl get nodes -o wide`
        * Get the ExternalIP of the nodes. If it's none -> 'localhost'
      * `kubectl describe svc/simpleservice`
        * Check the 'nodePort'
      * `curl http://NodeIP:NodePort`
        * Problems:
          * Problem1: TimeOut
            * Attempt1: Create 'simpleDeployment.yaml' & 'simpleService2.yaml' and try to run it 
            * Note1: I get this error running kind cluster and minikube
            * Note2: Check kube-proxy logs -- `kubectl get all -n kube-system` & `kubectl logs pod/kube-proxy-Name -n kube-system`
            * Solution: TODO: ?

# := Kubernetes Object
## Allows
### Losing coupling between dependent Pods
* Check 'kubernetesObject/allows/losingCouplingBetweenDependantPods/'
* `kubectl apply -f service.yaml` & `kubectl apply -f dependantPods.yaml`
  * Service handles the discovery mechanism, without hard coding IPs
### If pods die or are replicated -> no impact in the network application
* Check 'kubernetesObject/allows/podsCanDieOrReplicateWithoutImpact/'
* `kubectl apply -f service.yaml` & `kubectl apply -f deployment.yaml`
  * Service is a NodePort, but since I am having problems to access to it (Check TODO:) -> `kubectl port-forward service/my-nginx-service 8080:80`
    * Check that it works `curl localhost:8080`
* Although you delete pods -- `kubectl delete pod/PodName` --, Service distributes the traffic and you can still access to the pods' content -- `curl localhost:8080` --
### Incoming 'port' -- is mapped to -> 'targetPort'
* Check 'kubernetesObject/allows/incomingPortMappedToTargetPort/'
* `kubectl apply -f pod.yaml` & `kubectl apply -f services.yaml`
  * Services are NodePorts, but since I am having problems to access to it (Check TODO:) -> `kubectl port-forward service/my-nginx-service 8081:81`
    * Check that it works `curl localhost:8081`
      * 8081 -- port-forward -> 81 -- Service -> 80 -- Pod's containerPort --
* `kubectl describe svc/my-nginx-service` & `kubectl describe svc/my-nginx-service2`
  * Check that Service 'my-nginx-service2' targetPort = port, although it's not defined in the manifest
## spec.ports[x]
### spec.ports[x].appProtocol
* Check 'specPorts/appProtocol/' 
* UDP
  * `kubectl apply -f udpPod.yaml` & `kubectl apply -f udpService.yaml`
  * `kubectl get svc/my-udp-service`
    * Get the ClusterIP
  * `nc -u -v -z ClusterIP 53`
    * Check that the connection is succeeded
* SCTP
  * `kubectl apply -f sctpService.yaml`
  * `kubectl get svc/my-sctp-service`
    * Get the ClusterIP
  * TODO: How to hit SCTP requests?
    * Attempt1: `echo "Hello, SCTP!" | socat - SCTP:'10.100.190.6:5000'`
      * `brew install socat`
* 'kubernetes.io/h2c'
  * `kubectl apply -f h2cService.yaml`
  * `kubectl get svc/my-h2c-service`
    * Get the ClusterIP
  * Since I am having problems to access to it (Check TODO:) -> `kubectl port-forward service/my-h2c-service 8081:80`
    * Check that it works `curl -v --http2 http://localhost:8081'
      * 8081 -- port-forward -> 80 -- Service's port
* 'kubernetes.io/ws'
  * `kubectl apply -f wsService.yaml`
  * Since I am having problems to access to it (Check TODO:) -> `kubectl port-forward service/my-ws-service 8081:80`
    * Problems:
      * Problem1: It doesn't work
        * Solution: TODO: ?
  * Check that it works properly
    * Problems:
      * Problem1: TODO: How?
        * Attempt1: `websocat ws://ServiceIP:NodePort` & `websocat ws://ClusterIp:Port` 
* 'kubernetes.io/wss'
  * `kubectl apply -f wssService.yaml`
  * Since I am having problems to access to it (Check TODO:) -> `kubectl port-forward service/my-wss-service 8081:80`
    * Problems:
      * Problem1: It doesn't work
        * Solution: TODO: ?
  * Check that it works properly
    * Problems:
      * Problem1: TODO: How?
        * Attempt1: `websocat wss://ServiceIP:NodePort` & `websocat wss://ServiceIP:Port` & `websocat wss://NodeIP:Port` & `websocat wss://NodeIP:NodePort` 
* The values are mirrored by 'Endpoints' and 'EndpointSlices'
  * TODO: 
### spec.ports[x].targetPort
* Check 'specPorts/targetPort/'
* Pod’s fields to map
  * 'Pod.spec.containers[x].ports.name'
    * `kubectl apply -f pod.yaml`
    * `kubectl apply -f service.yaml`
    * Since I am having problems to access to it (Check TODO:) -> `kubectl port-forward service/nginx-service 8081:80`
      * Check that it works 'curl http://localhost:8081'
        * 8081 -- port-forward -> 80 -- Service's port -- map to -> Pod's port
  * 'Pod.spec.containers[x].ports.containerPort'  -- common one used --
    * `kubectl apply -f pod.yaml`
    * `kubectl apply -f service.yaml`
  * Since I am having problems to access to it (Check TODO:) -> `kubectl port-forward service/nginx-service 8081:80`
    * Check that it works 'curl http://localhost:8081'
      * 8081 -- port-forward -> 80 -- Service's port -- map to -> Pod's port
## IP address is assigned
* `kubectl apply -f simpleService.yaml`
* `kubectl describe svc/simpleservice`
  * Check that there is IP field
## Service Controller
* `kubectl apply -f simpleService.yaml`
* There is a Service Controller
  * `kubectl logs -n kube-system kube-proxy-Name | grep service`
    * Check that it provisions the Service Config Controller
  * `kubectl logs -n kube-system kube-controller-manager-Name | grep service`
    * Check that it provisions the Service Controller & Service Account Controller
* Others
  * `kubectl logs -n kube-system kube-apiserver-Name | grep simpleservice`
    * kube-api-server allocates the clusterIP
## Name convention following RFC1035
## == Abstraction of Logical Set of Endpoints + Policies
* Check 'kubernetesObject/abstraction/'
* `kubectl apply -f microServiceNetworkApplication.yaml` & `kubectl apply -f feService.yaml`
  * You have a set of entry points to the pod
  * Service is a NodePort, but since I am having problems to access to it (Check TODO:) -> `kubectl port-forward service/frontend-service 8084:82`
    * `curl localhost:8084`
      * Problems:
        * Problem1: "an error occurred forwarding 8084 -> 8080: error forwarding port 8080 to pod 4b0e53d79ccd6e335a02160214369000db88e4454563d8224662bb37739b038d, uid : exit status 1: 2023/12/10 16:27:05 socat[431289] E connect(5, AF=2 127.0.0.1:8080, 16): Connection refused"
          * Solution: TODO: ?

# Service without selectors
* Check 'ServiceWithoutSelectors/'
* Neither EndpointSlices nor Endpoints are created automatically
  * `kubectl apply -f service.yaml`
  * `kubectl get endpoints my-service` & `kubectl get endpointslices` -- Confirm it --
* Service -- adding an EndpointSlice manually, can be mapped to  -> network address and port
  * `kubectl apply -f endpointSlice.yaml`
  * `kubectl describe svc/my-service`
    * Check that there are endpoints associated
    * Problems:
      * Problem1: No endpoint associated
        * Attempt1: Create the EndpointSlice first
        * Note1: `kubectl get endpointslice` exists and it's created apparently fine
        * Solution: TODO: ?
## use cases
### different database cluster / environment
* Check 'differentDDBBClusterPerEnvironment/'
* `kubectl apply -f forProduction.yaml` & `kubectl apply -f forNonProduction.yaml`
* `kubectl get endpointslice`
  * Check that the EndpointSlices have been created
### service - points to -> another service in
#### another namespace
* `kubectl create namespace namespace-a` & `kubectl create namespace namespace-b` 
* `kubectl apply -f serviceNamespaceA.yaml`
* `kubectl apply -f serviceNamespaceB.yaml`
* `kubectl apply -f podInNamespaceB.yaml`
* `kubectl exec -it test-pod --namespace=namespace-b -- /bin/sh`
  * get access to the sh command to the pod
  * `wget -O - http://my-service-in-namespace-b.namespace-b.svc.cluster.local 80`
    * Problems:
      * Problem1: "wget: can't connect to remote host (10.96.179.231): Connection refused"
        * Note1: `nslookup my-service-in-namespace-b.namespace-b.svc.cluster.local` DNS lookup of the service-b works correctly
        * Note2" `nslookup my-service-in-namespace-a.namespace-a.svc.cluster.local` doesn't find, since there is no access to the other service directly
        * Attempt1: Trying to install curl via pod's command
        * Solution: TODO: ?
#### another cluster
TODO:
### migrate a workload to Kubernetes
TODO:

# Endpoints
## requirements to Service -- forward traffic to -> pods
* ALL Service's selectors -- should match with -- pod's labels
  * Check the pods 'nginx-nomatched' & 'nginx' and the services 'my-service' & 'my-service-2'
## invalid IPs
### loop back
* `kubectl apply -f loopbackEndpoints.yaml`
### link local
* `kubectl apply -f linkLocalEndpoints.yaml`
### clusterIPs
* `kubectl apply -f simpleClusterIPService.yaml` * `kubectl apply -f clusterIPEndpoints.yaml`
  * Problems:
    * Problem1: Why it's possible to create the Endpoints for 'cluster-ip-endpoints'?
      * Solution: TODO:
      * Attempt1: Specifying directly the clusterIP in the service
      * Attempt2: Use Service without selectors
### non-existing pod
* `kubectl apply -f nonExistingPod.yaml`
  * Problems:
    * Problem1: Why the Endpoint was created properly?
      * Solution: TODO:

# Multi-port services
* `kubectl apply -f service.yaml`

# Service type
* Check 'serviceType/'
## ClusterIp
* `kubectl apply -f clusterIPService.yaml`
* IP characteristics
  * 1!
    * `kubectl describe svc/simpleservice` and check IP entry
  * comes from the pool of IP addresses reserved by the cluster -- normally 10.96.0.0/12 --
    * Check that the one got is within that range
* How to reach the service?
  * Inside the cluster
    * `kubectl apply -f dummyPod.yaml` 
      * Run a dummy pod to test the connectivity inside the cluster
        * Problems:
          * Problem1: 'Failed to connect to simpleservice port 80'
            * Attempt1: `curl http://simpleservice:80` & `curl http:simpleservice:80`
            * Attempt2: `curl clusterIP:80` & `curl http://clusterIP:80` 
            * Solution: TODO:
  * Outside the cluster, by default, it's impossible to reach !! 
    * Attempt1: `kubectl port-forward svc/simpleservice 8080:80` invalid
    * Solution: `kubectl proxy --port=8001`
      * `curl -v http://localhost:8001/api/v1/namespaces/default/services/simpleservice:80/proxy/` to check that you can access to it
* `spec.clusterIP`
  * `kubectl get pod -n kube-system -l component=kube-apiserver -o jsonpath='{.items[0].spec.containers[0].command}' | grep service-cluster-ip-range > out.txt` & `code out.txt`
    * Look for 'service-cluster-ip-range' property defined
## NodePort
* Check the nested functionality from clusterIP
  * `kubectl apply -f nodePortService.yaml`
  * `kubectl get svc/nginx` checking that it has got clusterIP associated to
* Specify list of IP addresses to server NodePort Services
  * `kind create cluster --config=clusterWithSeveralNodes.yaml` & `kubectl cluster-info --context kind-kind`
    * Create a cluster with several nodes and use it
  * `kubectl get all -n kube-system -o wide | grep kube-proxy` & `kubectl describe pod/kube-proxy-NAME -n kube-system`
    * Check that there are several kube-proxy pods, all consuming from the same configMap 'kube-proxy'
  * `kubectl -n kube-system edit configmap kube-proxy` & edit to some valid IP block `nodePortAddresses: ["127.0.0.0/8"]`
  * `kubectl -n kube-system delete pod -l k8s-app=kube-proxy` delete them to be rebooted to pick the new configMap
  * `kubectl apply -f nodePortService.yaml` to deploy it again & `kubectl describe svc/my-service` should display loopback in some place
    * Problems:
      * Problem1: Why no reference about loopback is displayed?
        * Solution: TODO: ?
* Service is exposed on each Node’s IP at a static port
  * `kind create cluster --config=clusterWithSeveralNodes.yaml` & `kubectl cluster-info --context kind-kind`
    * Create a cluster with several nodes and use it
  * `kubectl apply -f nodePortService.yaml`
  * it should be reachable outside the cluster via `<NodeIP>:<NodePort>`, where '<NodeIp>' should be valid for all cluster's nodes
    * `kubectl get nodes -o wide` to check the NodeIP & `kubectl describe svc/my-service` to check the NodePort
    * `curl http://NodeIp:NodePort`
      * Problems:
        * Problem1: Not reached
          * Solution: TODO:
* Static port
  * is allocated by Control Plane
    * Go into the Control Plane node & check if there is a listener on that port
      * `docker exec -it ControlPlaneID bin/sh` & `ss -tlnp | grep NodePort`
        * Note: Docker because we are using kind cluster whose nodes are docker containers
        * Problem: 'Not found the NodePort listening'
          * Solution: TODO:
  * belongs to 'service-node-port-range'
    * `kubectl describe pod/kube-apiserver-NAME -n kube-system` and check the commands passed to the container
      * Problems: 'How to check that the default value is [30000, 32767] if it's not passed'?
        * Attempt1: `docker exec -it ControlPlaneId sh` $ `env` to check the environment variables
        * Solution: TODO:
    * check the 2 bands of that range
      * HOW? TODO:
    * check that dynamic port allocation vs static port allocation
      * HOW? TODO:
  * same port is allocated by all the cluster's nodes
    * `docker exec -it ClusterNode bin/sh` & `ss -tlnp | grep NodePort`
      * Note: Docker because we are using kind cluster whose nodes are docker containers
        * Problem: 'Not found the NodePort listening'
          * Solution: TODO:
* Kubernetes additionally allocates another port
  * How to check? TODO:
* Uses
  * set up a custom load balancer 
    * What? TODO:
* `ServiceNodePortStaticSubrange`
  * Check TODO:
## LoadBalancer
* Check the nested functionality from NodePort
  * `kubectl apply -f loadBalancerService.yaml` & `kubectl describe svc/load-balancer`
    * Check that it contains also an assigned NodePort and a clusterIp
* Service is exposed externally — via — an external load balancer
  * Create an external Load Balancer (Cloud provider, local, ..)
    * TODO: 
* `spec.loadBalancerIP`
  * Create an external load balancer with an IP
    * TODO:
* `spec.allocateLoadBalancerNodePorts`
  * enable the nodePort allocation
    * `kubectl describe svc/load-balancer-disable-nodeports` check that there are no nodePorts allocated
* 'MixedProtocolLBService'
  * By default, it's true
    * Problems: 'How to check?'
      * Attempt1: `kubectl logs pod/kube-apiserver-kind-control-plane -n kube-system | grep feature-gate`
      * Attempt1: `kubectl describe pod/kube-apiserver-kind-control-plane -n kube-system -o=jsonpath='{.items[0].spec.containers[0].command}' | grep feature-gate`
      * Solution: TODO:
  * Check that we can deploy 'load-balancer-severalports'
* `.spec.loadBalancerClass`
  * Check 'load-balancer-class' Service
  * Create an external load balancer -- TODO:
### uses
* Mixed environments
  * TODO:
* Split-horizon DNS environment
  * TODO:
## ExternalName
* ≠ Superset of previous service types
  * `kubectl apply -f externalNameService.yaml` & `kubectl get svc` & `kubectl describe svc/external-name`
    * Check that it has not got clusterIp nor NodePort
* Requirements
  * CoreDNS ≥ v0.0.8
    * `kubectl get pods -n kube-system | grep coredns` & `kubectl describe pod/coredns-5dd5756b68-l8xqt -n kube-system`
      * Check the containers.coredns.image
  * kube-dns ≥ v1.7 has been replaced by the previous one
    * `kubectl get pods -n kube-system` and check that you can't find it
* Service — is mapped to a → DNS name
  * within the cluster
    * `kubectl exec -it curl-pod -- sh` & `nslookup my.database.example.com` should be able to identify it
      * Problems: 'don't identify it, based on DNS name'
        * Solution: TODO:
  * outside the cluster
    * `nslookup my.database.example.com` should be able to identify it
* `spec.externalName`
  * values
    * DNS name      -- Check external-name --
    * IPv4 address   -- Check external-name-passingasipv --
* == special case of Service without selectors + DNS names ->
  * no associated pod  -- although you add a label -- 
    * Check 'external-name-without-selectors' and `kubectl get endpoints` without getting anything
  * no proxying available -- `kubectl proxy external-name-without-selectors` or `kubectl proxy external-name-passingasipv` --
    * Problems: I can do it, without error
      * Solution: TODO:
* hostname used by clients inside the cluster ≠ name referenced by the ExternalName
  * `kubectl exec -it curl-pod -- sh` & `curl http://external-name.default.svc.cluster.local` should reach it, using the service's name
    * Problems: 'Not reach it'
      * Solution: TODO:
  * Example1: HTTP requests with Host: header → origin server doesn’t recognize it
    * TODO:
  * Example2: client provides a hostname → TLS servers not able to provide a certificate, matching with that hostname
    * TODO:

# Label selectors
* Check 'labelSelectors/'
* If you don’t specify it → Endpoints objects not created
  * Check 'ServiceType/externalName' and `kubectl get endpoints` checking that there are no endpoint objects for externalName services
  * -> you need to map manually
    * `kubectl apply -f serviceWithoutSelectors.yaml` & `kubectl apply -f endpoints.yaml` and you can check that it's mapped via `kubectl describe svc/service-without-selectors`

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

# Service Discovery Mechanisms
## modes supported by Kubernetes
### environment variables
* kubelet adds default variables to the pods
  * `kubectl apply -f simplePod.yaml` & `env` to check it 
### DNS
* Check 'overview/KubernetesComponents' & '../DNSForServicesAndPods'

# External IPs
* Check 'externalIPs/'
* `kubectl apply -f service.yaml`
* Based on `.spec.externalIPs[]` & `.spec.ports[].port` -> you can access to it
  * In this case -- "198.51.100.32:80"