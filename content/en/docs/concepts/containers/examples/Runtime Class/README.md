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

# Configure CRI implementation on nodes
## Set up CRI runtimes
* Check '../productionEnvironment/CR'
* Execute / each Cluster's node -- `docker exec -it ContainerId sh` --
  * store the modules to load under '/etc/modules-load.d/k8s.conf'
```   
    cat <<EOF | tee /etc/modules-load.d/k8s.conf
    overlay
    br_netfilter
    EOF
```
  * load kernel modules into the running kernel
```
  modprobe overlay
  lsmod | grep overlay            # Check that the module is loaded
  modprobe br_netfilter
  lsmod | grep br_netfilter       # Check that the module is loaded
```
    
* Problems: The modules are not loaded
  * Solution: TODO:
  
  * establish IPv4 rules and iptables
```
  cat <<EOF | tee /etc/sysctl.d/k8s.conf
  net.bridge.bridge-nf-call-iptables  = 1
  net.bridge.bridge-nf-call-ip6tables = 1
  net.ipv4.ip_forward                 = 1
  EOF
```
  * reboot sysctl -> apply new params  -- `sysctl --system` --
    * check that the rules are applied -- `sysctl net.bridge.bridge-nf-call-iptables net.bridge.bridge-nf-call-ip6tables net.ipv4.ip_forward` --
## Configure Runtime handlers
* based on CR
  * containerd
    * `[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.${HANDLER_NAME}]` under '/etc/containerd/config.toml'
  * CRI-O
    * `[crio.runtime.runtimes.${HANDLER_NAME}] runtime_path = "${PATH_TO_BINARY}"` under '/etc/crio/crio.conf'
  * Docker Engine
    * not supported ? TODO:

# Create RuntimeClass
* There should exist 1 handler / RuntimeClass
* `kubectl apply -f RuntimeClass.yaml`
  * `kubectl get all` doesn't appear since it's not associated to a namespace

# Specify in a pod
* Previous steps must have been done
* `kubectl apply -f podUsingRunTimeClass.yaml`
  * `kubectl get pods/mypod`
    * if Runtime Class doesn't exist or CRI can not run the handler -> 'Fail' phase for the pod

# .scheduling
* Ensure pods with this RuntimeClass — are scheduled to — nodes which supports it
  * `kubectl apply -f node.yaml` & `kubectl apply -f RuntimeClass.yaml` & `kubectl apply -f podUsingRunTimeClass.yaml`
    * Problems:
      * Problem1: Node is "NotReady"
        * Solution: TODO:

# .overhead
* specifying the overhead of running pods / use this RuntimeClass
  * `kubectl apply -f RuntimeClass.yaml` & `kubectl apply -f podUsingRunTimeClass.yaml`
    * Problems:
      * Problem1: "no runtime for "example-runtime-handler" is configured"
        * Solution: TODO: