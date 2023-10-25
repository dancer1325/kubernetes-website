# Pre requisites
* Locally
  * Run some docker daemon
    * [Docker desktop](https://www.docker.com/products/docker-desktop/)
  * Install some local cluster tool
    * [minikube](https://minikube.sigs.k8s.io/docs/start/)
    * [kind](https://kind.sigs.k8s.io/)
  * kubectl > v1.14
* Kubernetes playground
  * [killerCoda](https://killercoda.com/playgrounds/scenario/kubernetes)
  * [PlayWithKubernetes](https://labs.play-with-k8s.com/)

# How to run it?
* `kubectl apply -f configmap.yaml`
  * You can play adjusting the configMap file
  * `kubectl describe configmaps example-redis-config`
    * Confirm, once you run apply again, that the configMap has been updated. BUT THE POD HAS NOT UPDATED THE CONFIGMAP -> you should
      * `kubectl delete pod redis`
      * `kubectl apply -f https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/pods/config/redis-pod.yaml`
* `kubectl apply -f https://raw.githubusercontent.com/kubernetes/website/main/content/en/examples/pods/config/redis-pod.yaml`
  * Create sample pod placed also in this documentation, which can be reached through internet
* `kubectl get pod/PodName configmap/ConfigMapName`
  * Check pod and configmap
  * `kubectl get pod/redis configmap/example-redis-config`
* `kubectl exec -it redis -- sh` & `cd /redis-master` & `ls` `cat redis.conf`
  * Commands to get into the pod via sh, and checking that the empty file of the configMap exist
* `kubectl exec -it redis -- redis-cli`
  * Commands to get into the pod via redis-cli
  * `CONFIG GET maxmemory` & `CONFIG GET maxmemory-policy`
    * Check [eviction policies](https://redis.io/docs/reference/eviction/)