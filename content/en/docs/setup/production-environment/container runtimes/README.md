# Pre requisites
* Locally or cluster you fully control (=== configure control plane), where you can
  * Install some local cluster
    * tool
      * [minikube](https://minikube.sigs.k8s.io/docs/start/)
      * [kind](https://kind.sigs.k8s.io/)
  * Run a local cluster
    * [minikube]  `minikube start`
    * [kind] `kind create cluster`

----

# CR
* It's installed / cluster's node
  * Check running different CLI versions
    * `docker --version` / `ps aux | grep docker`
    * `containerd --version`
    * ...
* required configurations in the nodes -- Check '../concepts/containers/Runtime Class'
  * forward IPv4
  * let iptables
## supported by Kubernetes
### containerd
* support CRI > v1alpha2
  * TODO:
* [Install](https://github.com/containerd/containerd/blob/main/docs/getting-started.md)
  * Not possible to install in MAC OS
* with cgroup driver
  * systemd as cgroup driver
    * set the configurations and check that you can run containers
      * TODO:
* Override the sandbox (pause) image
  * What to check? TODO:
### CRI-O
* support CRI > v1alpha2
  * TODO:
* [Install](https://github.com/cri-o/cri-o/blob/main/install.md#readme)
  * Just possible in Linux machines
* with cgroup driver
  * systemd as cgroup driver
    * check that it's the default one
      * :TODO
  * cgroups as cgroup driver
    * adjust it TODO:
* Override the sandbox (pause) image
  * What to check? TODO:
### Docker Engine
* TODO:
### Mirantis CR / Docker Enterprise Edition
* TODO:
* Override the sandbox (pause) image
  * What to check? TODO:

# cgroup drivers
* Just available in Linux-OS
  * `docker exec -it NodeContainerId sh` & `cat /etc/os-releases`
* `cat /proc/cgroups` display the cgroup subsystems
## allows
* if kubelet’s cgroup driver (version & configuration) == CR’s cgroup driver (version & configuration)
  * Check
    * kubelet’s cgroup driver
      * `docker exec -it NodeId sh` & `cat /var/lib/kubelet/config.yaml` displaying the KubeletConfig
    * CR’s cgroup driver
      * Problems: How? TODO:
        * Attempt1: `cat /var/log/docker.log` OR `cat /var/log/containerd/containerd.log` nothing found
        * Attempt2: `kubectl describe node <node-name> | grep -i container-runtime` nothing found
        * Note: Could it be influence by running the cluster locally via kind?
  * → kubelet & CR — interface via their cgroup driver with —  cgroup to → enforce @Resource Management for Pods and Containers
## types
### cgroupfs driver
* TODO:
### systemd cgroup driver
* TODO:
