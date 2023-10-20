# Content
* 1! ConfigMap data
  * `kubectl create configmap 1-configmap --from-literal=special.how=very`
    * Create a ConfigMap via kubectl from literal which it's a key-value pair
  * [busybox](https://hub.docker.com/_/busybox)
  * `kubectl create -f .`
  * `kubectl logs podName`
    * Check logs of the pod, displaying environment variables -- also the already including in the container and no passed by us -- 
* Several ConfigMaps
  * `kubectl create -f severalconfigmaps.yaml`
    * Create the configmaps
  * [busybox](https://hub.docker.com/_/busybox)
  * `kubectl create -f pod-several-configmap-envvariable.yaml`
  * `kubectl logs podName`
    * Check logs of the pod, displaying environment variables -- also the already including in the container and no passed by us -- 