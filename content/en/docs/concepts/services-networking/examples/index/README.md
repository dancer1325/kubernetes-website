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

# Goal
* Kubernetes Network Model

---

# IP-per-pod
* `kubectl apply -f deployment.yaml`
  * `kubectl describe PodName`
    * Check that
      * each of the pods has got it's own IP
      * IP addresses are into the same subnet

# Networking implementation
## Any node’s pod — can communicate without NAT with — any other node2’s pod
* `kubectl label nodes kind-worker  node-type=worker`
  * Add a label in one of the Cluster's nodes
* `kubectl apply -f testingPod.yaml`
  * Deploy the testing pod in one of the worker nodes
* `kubectl get pods -o wide`
  * Check the nodes in which the different pods are living
* `kubectl exec -it pod/busybox-pod -- sh`
  * go into the sh command
  * `ping PodIPLivingInAnotherNode`
    * Check that we receive packages
## Any node’s agent — can communicate with — any pod in that node
* `kubectl get pods`
  * If Status column is displaying information == kuelet can reach them -- since it checks the healthy of the pods --
* `docker ps`
  * Check the Control Plane's IP
* `kubectl exec -it pod/busybox-pod -- sh`
  * go into the sh command
  * `ping IPNodeLivingKubelet`
    * Check that we receive packages == from testing pod, you can reach the other way around


# Kubernetes networking addresses
## Pod1 ← can communicate via @Cluster Networking  with → pod2
* 2 pods in the cluster, so we try to get access from one to another via Service
* `kubectl apply -f networkingAddressPod1.yaml` & `kubectl apply -f networkingAddressPod2.yaml` & `kubectl apply -f networkingAddressService.yaml`
* `kubectl exec -it pod/networkingaddress2 -- sh`
  * `curl networkingaddress1IP:80`
    * Return the dummy nginx 'index.html' configuration
  * `curl networkingaddress`
    * Trying to reach the pod/networkingaddress1 -- via Kubernetes Service --
    * Problems:
      * Problem1: "curl: (7) Failed to connect to networkingaddress port 80 after 3 ms: Couldn't connect to server"
        * Note1: `nslookup networkingaddress` "server can't find networkingaddress.cluster.local: NXDOMAIN"
        * Note2: `kubectl get pods -n kube-system -l k8s-app=kube-dns` to check that CoreDNS is working fine
        * Note3: `kubectl logs -n kube-system CoreDNSPodName` to check that nothing weird is happening
        * Note4: `kubectl get endpoints networkingaddress` should show ENDPOINTS entries, matching with the specific pod 
        * Solution: Service matches the pod by labels!!!!

# Host ports
* `kubectl apply -f hostPortPod.yaml`
* `kubectl get pods -o wide`
  * Check the IP column value
  * `curl http://NodeIp:8080`
    * It should display the dummy value for nginx containers
    * Problems:
      * Problem1: Time out
        * Attempt1: Switch to another hostPort
        * Attempt2: `kubectl get nodes -o wide` check the INTERNAL-IP and the hostPort
        * Solution: TODO: ?
