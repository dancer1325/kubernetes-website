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

# Kubernetes Cluster
* `kind create cluster --config=kindClusterWithSeveralNodes.yaml`
  * Problems: 
    * Problem1: "failed to create cluster: node(s) already exist for a cluster with the name "kind"
      * Solution: `docker ps` and check the containerName of kind process. `docker kill ContainerNameOfKindContainer`. Or
  * `kubectl get nodes`
    * cluster is created with 3 nodes
      * 1 control plane == master node 
      * 2 worker nodes == worker node
    * Problems:
      * Problem1: Why the nodes declared as worker don't appear as worker
        * Attempt1: `kind create cluster --config=kindClusterWithoutSpecifyingImage.yaml`
        * Attempt2: `kind create cluster --config=kindClusterWithSeveralNodes.yaml`
        * Solution: `docker inspect ContainerId`
        * Note: [KindIssue](https://github.com/kubernetes-sigs/kind/issues/3421)
* Check 'tutorials/' to realize that any containerized application can be run there

# Control Plane / Master node components
* Using the previous cluster with 1 control plane + 2 worker nodes set up
* `kubectl get pods -n kube-system --field-selector spec.nodeName=kind-control-plane -o wide`
  * display the Control Plane's components
    * [CoreDNS](https://www.cncf.io/projects/coredns/)
      * All Kubernetes Cluster should have a cluster DNS, used as default one in [kind clusters](https://kind.sigs.k8s.io/)
      * Check that it's a Cluster DNS in fact
        * `kubectl describe pods/coredns-RandomName -n kube-system` and check that the image is of coreDNS
        * `kubectl get configmap coredns -n kube-system -o yaml` check that it's CoreDNS configuration
        * `kubectl port-forward -n kube-system <coredns_pod_name> 9153:9153` and open in your browser 'http://localhost:9153/metrics' checking the metrics exposed
          * 9153 is the default prometheus port
    * etcd
      * Check that it's an etcd really
        * `kubectl describe pods/etcd-kind-control-plane -n kube-system` and check that it contains a container with an etcd image
    * [kindnet](https://github.com/aojea/kindnet)
      * Network plugin for kind
    * kube-apiserver
      * Check that it's a kube-apisever in fact
        * `docker ps` and check that the 'control-plane' forwards '127.0.0.1:PortOfYourServer->6443/tcp', being 6443 the default port of kube-apisever
        * `cat .kube/config` and check that in your cluster's server == PortOfYourServer
      * Handles the conversion between API versions transparently
        * Select a resource in the Kubernetes API with several available version. _Example:_ https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#api-groups autoscaling
        * Check 'sameResourceHandledByDifferentAPIVersions/'
          * `kubectl apply -f autoscalingV1.yaml` & `kubectl apply -f autoscalingV2.yaml`
          * Problems:
            * Problem1: How to migrate from v1 to v2?
              * Attempt1: [With minikube] `minikube delete` & `minikube start --kubernetes-version=v1.28`
              * Solution: Check 'reference/APIOverview/DeprecatedAPIMigrationGuide/'
      * `docker exec -it ControlPlaneContainer cat /etc/kubernetes/manifests/kube-apiserver.yaml > apiserverPod.yaml` & `code apiserverPod.yaml`
        * be able to check the kube-apiserver podSpec
    * kube-controller-manager
      * Check that it's a kube-controller-manager really
        * `kubectl describe pods/kube-controller-manager-kind-control-plane -n kube-system` and check that it contains a container with a kube-controller-manager image
    * kube-proxy
      * Check that it's a kube-proxy really
        * `kubectl describe pods/kube-proxy-RandomNameSet -n kube-system` and check that it contains a container with a kube-proxy image
      * Check that it contains network rules
        * `kubectl logs pods/kube-proxy-RandomNameSet -n kube-system` checking references to iptables
        * `docker exec -it ControlPlaneContainer sh` and `iptables -t nat -L -n -v` to check the iptables rules
      * Check the kube-proxy configuration served in a configmap
        * `kubectl get configmap kube-proxy -n kube-system -o json | jq '.data' | jq -r '.["config.conf"]'`
          * If you want to extract the KubeProxyConfiguration
      * Check the logic about OS' packet filtering layer
        * TODO:
      * modes in which to run
        * `kubectl apply -f deployments.yaml` & `kubectl apply -f services.yaml`
        * Iptables mode
          * `docker exec -it ControlPlaneContainer sh` &
          * `iptables-save` to check that
            * / service’s port
              * 1 rule in KUBE-SERVICES -- Which one? TODO: --
              * 1 rule in KUBE-SVC-HASH -- Which one? TODO: --
            * / pod’s endpoint
              * small number rules in KUBE-SVC-HASH -- Which one? TODO: --
              * small number rules in KUBE-SEP-HASH -- Which one? TODO: --
        * ipvs mode
          * `docker exec -it ControlPlaneContainer sh` &
          * `ipvsadm -ln` to check that
            * Problems:
              * Problem1: 'ipvsadm' not found
                * Solution: `apt-get update` & `apt-get install ipvsadm`
              * Problem2: No result at all found for the different next cases
                * Solution: TODO:
            * virtual servers are created /
              * service’s port -- '80' from 'kubeproxy-modes'  --
              * NodePorts -- 'kubeproxy-modes-nodeport' --
              * external IPs -- 'kubeproxy-modes-externalip' --
              * load-balancer IPs -- 'kubeproxy-modes-loadbalancerip' --
            * real servers are created /
              * pod’s endpoint
      * check if kube-proxy is proxying -- 'kubeproxy-modes' --
        * `kubectl describe svc/kubeproxy-modes` to check the IP and port & `docker exec -it ControlPlaneContainer sh` &
        * `curl serviceIP:servicePort` getting result
    * kubelet
      * Check that this agent is included into the container
        * `docker exec ControlPlaneContainer ps aux | grep kubelet > kubelet.txt` and `code kubelet.txt` checking which contains '/usr/bin/kubelet'
      * Check that it's running
        * `docker exec -it ControlPlaneContainer sh` and `systemctl status kubelet` checking that it's running
      * Check all the logs
        * `docker exec -it ControlPlaneContainer journalctl -u kubelet`
      * Check its configuration
        * `systemctl status kubelet` finding 'config' path
        * `cat cat /var/lib/kubelet/config.yaml` displaying the configuration
      * acts as bridge between master node < - > rest of nodes -- How to check? TODO: --
      * fetches individual container statistics from the Container Runtime — via — Container Runtime Interface -- How to check? TODO: --
    * kube-scheduler
      * Check that it's a kube-scheduler really
        * `kubectl describe pods/kube-scheduler-kind-control-plane -n kube-system` and check that it contains a container with a kube-scheduler image
        * `kubectl logs pod/kube-scheduler-kind-control-plane -n kube-system` and check that 'PodDisruption' or 'watch' or .. are logged
      * Run commands -- `kubectl exec -it kube-scheduler-kind-control-plane -n kube-system COMMAND`
        * Solution: TODO:
    * Container Runtime
      * Check that there's a container runtime really
        * `docker exec -it ControlPlaneContainer sh` and check if there are some container runtime implementation
          * `ps aux | grep docker` and `systemctl status docker`
          * `ps aux | grep containerd` and `systemctl status containerd`
  * `--field-selector` is to filter the related to control plane node
  * 'cloud-controller-manager' is not present since we are running it locally

# Worker Node Components
* Using the previous cluster with 1 control plane + 2 worker nodes set up
* `kubectl get pods -n kube-system --field-selector spec.nodeName=kind-worker -o wide` / `kubectl get pods -n kube-system --field-selector spec.nodeName=kind-worker-2 -o wide`
  * display the Worker Node's components
    * [kindnet](https://github.com/aojea/kindnet)
      * Network plugin for kind
    * kube-proxy
      * Check that it's a kube-proxy really
        * `kubectl describe pods/kube-proxy-RandomNameSet -n kube-system` and check that it contains a container with a kube-proxy image
      * Check that it contains network rules
        * `kubectl logs pods/kube-proxy-RandomNameSet -n kube-system` checking references to iptables
        * `docker exec -it ControlPlaneContainer sh` and `iptables -t nat -L -n -v` to check the iptables rules
      * Check the kube-proxy configuration served in a configmap
        * `kubectl get configmap kube-proxy -n kube-system -o json | jq '.data' | jq -r '.["config.conf"]'`
          * If you want to extract the KubeProxyConfiguration
      * Check the logic about OS' packet filtering layer
        * TODO:
