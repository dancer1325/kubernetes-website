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


# Pod
## Any number of different types of volumes can be used
* `kubectl apply -f PodWithSeveralVolumes/configmap.yaml`
* `kubectl apply -f PodWithSeveralVolumes/persistentvolumeclaim.yaml `
* `kubeclt apply -f PodWithSeveralVolumes/pod.yaml `
  * Create a pod, using different types of volumes simultaneously
  * Although you run this command before, and not running till volumes are up, pod listens on

# Containers
## If you restart a container with volumes - volumes are preserved
* Update `spec.container[x].image` & `kubectl apply -f PodWithSeveralVolumes/pod.yaml`
  * pod is restarted
  * volumes are preserved -- `kubectl get pvc` & `kubectl get pods` --

# Volume
## Depending on type of volume
* Different directory's structure
  * `kubectl describe configmap/my-config-map` vs `kubectl describe pvc/my-pvc`
* Different medium which backs it
  * `kubectl get storageclass PVCStorageClassName`
* Content
  * `kubectl describe configmap/my-config-map` vs (`kubectl exec -it pod-multivolume -- /bin/sh` + `ls /var/data` or `kubectl cp pod-multivolume:/var/data`) 
