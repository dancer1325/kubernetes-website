# Pre requisites
* Locally or cluster you fully control (=== configure control plane), where you can
  * Run some docker daemon
    * [Docker desktop](https://www.docker.com/products/docker-desktop/)
  * Install some local cluster tool
    * [minikube](https://minikube.sigs.k8s.io/docs/start/)
    * [kind](https://kind.sigs.k8s.io/)
  * kubectl > v1.4
    * `kubectl get nodes -o=jsonpath=$'{range .items[*]}{@.metadata.name}: {@.status.nodeInfo.kubeletVersion}\n{end}'`
      * Check it in the current context's nodes
  * Linux environment
    * Reason: AppArmor is just valid in Linux

# Steps
* `kubectl get nodes -o wide`
  * Identify the node
  * If you use Kind cluster -> docker containers are created to be used as Kubernetes nodes
    * `docker ps -a`
      * Identify the related to the Kind node
* `docker exec -it DockerKindName sh`
  * We will run sh commands into the docker container, which is a Linux OS
  * `apt update` & `apt install apparmor`
    * Update apt dependencies and install apparmor
  * `zgrep -i apparmor /proc/config.gz`
    * Check if the kernel accepts apparmor
    * If "CONFIG_SECURITY_APPARMOR=y" === AppArmor is supported
    * If "# CONFIG_SECURITY_APPARMOR is not set" === AppArmor is not supported
      * Note: Current set up. MacOs Ventura v13.1, arm64, kind v0.20.0
      * Attempt1: `apt install apparmor-utils`
      * Attempt2: `lsmod | grep apparmor` or `modprobe apparmor`
      * Solution: TODO 
  * `cat /sys/module/apparmor/parameters/enabled`
    * Check and or enable AppArmor
  * Check that your CR (Container Runtime) supports AppArmor
    * Docker
    * CRI-O
    * containerd


