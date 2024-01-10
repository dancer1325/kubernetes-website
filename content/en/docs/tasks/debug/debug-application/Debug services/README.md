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
# Debug Services
## Pod is NOT running or restarting continuously
* Check '../Debug Pods'
* Check if you can access directly to the pods -- 'deployment-1' --
  * `kubectl apply -f deployments.yaml`
  * Run a dummy busybox to go into it, to get access directly to the pod
    * `kubectl run -it --rm --restart=Never busybox --image=gcr.io/google-containers/busybox sh` & `kubectl get pods -l app=nginx -o wide` to check the IPs
    * `wget -qO- 10.244.0.8:80`
      * '10.244.0.8' is the IP of one of the pods
      * '80' because it's the containerPort
## There are NOT enough endpoints / service
* Number of endpoints != Number of pods expected being member of your service  
* valid case -- 'deployment-1' & 'service-1' --
  * `kubectl apply -f deployments.yaml` & `kubectl apply -f services.yaml`
  * `kubectl get endpoints service-1` to check that the number of endpoints is fine & `kubectl get pods -o wide` checking that match with the Pod's ips
  * `kubectl get pods --selector=app=nginx` to check that there are >=1 pod associated with the service's selector
  * `kubectl run -it --rm --restart=Never busybox --image=gcr.io/google-containers/busybox sh` & `nslookup service-1`
    * run a dummy busybox
    * lookup by DNS name the service
* service’s selector does NOT match with a Pod -- 'selectors-mismatch' pod & service --
  * `kubectl apply -f pods.yaml` & `kubectl apply -f services.yaml`
  * `kubectl get endpoints selectors-mismatch` to check that there are NOT endpoints associated to it
  * `kubectl get pods --selector=selectors=mismatcH` to check that there is NOT pod associated with the service's selector

## Service is NOT working properly
* `podSpec.containers[x].containerPort` ≠ `serviceSpec.ports.targetPort` -- 'ports-not-match' pod & service --
  * `kubectl apply -f pods.yaml` & `kubectl apply -f services.yaml`
  * kubectl get endpoints ports-not-match` to check that there are endpoints associated to it
  * TODO:
* `serviceSpec.ports.targetPort` as string -- 'targetport-by-name-missing' pod & service --
  * `kubectl apply -f services.yaml` & `kubectl apply -f pods.yaml`
  * `kubectl get endpoints targetport-by-name-missing` to confirm that there are NOT endpoints, since it's not matching with a pod

## Service NOT created
* You forget to create the service -- 'service-not-created' --
  * `kubectl apply -f deployments.yaml` & `kubectl get pods -l case=service-not-created` to confirm that the pods are created
  * `nslookup service-not-created` or `wget -O- service-not-created` (Linux machines) to confirm that you are NOT able to resolve by DNS name 

## Service + Network Policy Ingress rules
* You add Network Policy Ingress rules, which restrict the incoming traffic to the pods
  * Check '../Concepts/Services,Load balancing, and networking/Network Policies'

## Service works by DNS name
* if pod lives in a namespace == service namespace -- 'deployment-1' & 'service-1' --
  * Check up, first section
* if pod lives in a namespace ≠ service namespace -- 'deployment-different-namespace' & 'service-different-namespace' --
  * `kubectl apply -f namespaces.yaml` & `kubectl apply -f deployments.yaml` & `kubectl apply -f services.yaml`
  * `kubectl get svc/service-different-namespace -n different-namespace` to check that the service lives in another namespace != default
  * Run a dummy pod to look up service by DNS
    * `kubectl run -it --rm --restart=Never busybox --image=gcr.io/google-containers/busybox sh` which runs in the namespace
    * `nslookup service-different-namespace` is NOT able to resolve the service
    * ways to resolve 
      * `nslookup service-different-namespace.different-namespace` or 
      * `nslookup service-different-namespace.different-namespace.svc.cluster.local` or
        * 'svc' since it's a service
        * 'cluster.local' because it's the cluster domain `kubectl get configmap coredns -n kube-system -o yaml`
      * `nslookup service-different-namespace.different-namespace.svc.cluster.local 10.96.0.10`
        * '10.96.0.10' is the clusterDNS Service IP `kubectl get svc/kube-dns -n kube-system` checking the 'CLUSTER-IP' column 
* wrongly configuration of ‘/etc/resolv.conf’ in your pod
  * Run a dummy pod to make checks
    * `kubectl run -it --rm --restart=Never busybox --image=gcr.io/google-containers/busybox sh` which runs in the default namespace
    * `cat /etc/resolv.conf` to check the DNS resolutions
      * 'nameserver ClusterDNSServiceIP' is passed into the kubelet via '--cluster-dns' flag
      * 'search suffixesToFindTheServices' is passed into the kubelet via '--cluster-domain' flag
        * 'default.svc.cluster.local' -- services in the local domain & default namespace --
        * 'svc.cluster.local' -- services in ALL namespaces --
        * 'cluster.local' -- names in the cluster --
* if you are NOT able to resolve Kubernetes master Service
  * Run a dummy pod to check if you are able to lookup the Kubernetes master default 
    * `kubectl run -it --rm --restart=Never busybox --image=gcr.io/google-containers/busybox sh` 
    * `nslookup kubernetes.default`

## Service works by IP
* check if you reach the service internally from the cluster -- 'deployment-1' & 'service-1' --
  * Run a dummy pod to check if you are able to reach the service internally
    * `kubectl run -it --rm --restart=Never busybox --image=gcr.io/google-containers/busybox sh`
    * `wget -qO- 10.96.103.177:80`
      * '10.96.103.177' is the 'CLUSTER-IP' of the service -- `kubectl get svc/service-1` to find --
      * '80' is the service-1's port

## Check kube-proxy
* Check if you use it in your cluster
  * `kubectl get pods -n kube-system | grep 'kube-proxy'`
* Check logs
  * `kubectl logs pod/kube-proxy-RandomNameSet -n kube-system`
* Check '../concepts/overview/KubernetesComponents' kube-proxy section

## Check network is properly configured
* Check that
  * “hairpin” traffic +
    * `docker exec -it ControlPlaneContainer sh` & 
      * `ps auxw | grep kubelet` finding the flag '--hairpin-mode=promiscuous-bridge'
      * check in the kubelet logs ('../concepts/overview/KubernetesComponents'), that you find 'Hairpin mode set to ' 
  * `kube-proxy` running in 'iptables' mode + 
    * check in the kubelet logs ('../concepts/overview/KubernetesComponents')
  * pods connected via bridge network
    * check kubelet has permission to operate in '/sys' node -- How? TODO: --
    * check kubelet has permission to manipulate linux bridge on node -- How? TODO: --

## Common approach for all cases
* Run a dummy busybox pod in the cluster and go into the shell
  * `kubectl run -it --rm --restart=Never busybox --image=gcr.io/google-containers/busybox sh`
* Go into a cluster's node
  * ``

