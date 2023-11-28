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

# Privileged mode
## Windows
* `kubectl apply -f podPrivilegedModeWindows.yaml`
  * Problems:
    * Problem1: "Failed to pull image "mcr.microsoft.com/windows/servercore:ltsc2022": no matching manifest for linux/arm64/v8 in the manifest list entries"
      * Attempt1: Switch to another release "ltsc2022"
      * Attempt2: [Switch to Windows containers](https://github.com/Sitecore/docker-images/issues/159) ? Try
      * Solution: TODO: ? 
## Linux
* `kubectl apply -f podPrivilegedModeLinux.yaml`
* Ways to check that it has got privileged rights
  * `kubectl exec -it pod/linux-pod sh`
    * Check that the container mounts host filesystems
      * `mount`
      * check that there is a volume '/dev/vda1' and then run `df -h /dev/vda1`
        * You check that it's mounted on the host system

---

# Probe