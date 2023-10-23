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
* ConfigMapWithAllKeyValuePairs
  * `kubectl create -f configmap.yaml`
  * `kubectl create -f pod-allkeyvaluepairs-configmap.yaml`
    * `envFrom` is the key
  * `kubectl logs podName`
    * Check logs of the pod, displaying environment variables -- also the already including in the container and no passed by us --
* UseAsPodCommands
  * `kubectl create -f severalconfigmaps.yaml`
  * `kubectl create -f pod.yaml`
  * `kubectl logs pod-passingenvironmentvariables-ascommand`
    * Check the right job of echo command
* OptionalConfigMap
  * Cases
    * Case1: configMap no existing
      * `kubectl create -f podOptionalConfigMap.yaml`
        * Create the pod
      * `kubectl logs pod-optionalconfigmap`
        * Since configMap is optional, and it doesn't exist -> environment variable no added
    * Case2: configMap exists but not the key
      * `kubectl create configmap optional-configmap --from-literal=a=Paula`
        * Create a ConfigMap via kubectl from literal which it's a key-value pair, without matching with the required by the pod
      * `kubectl create -f podOptionalConfigMap.yaml`
        * Create the pod
      * `kubectl logs pod-optionalconfigmap`
        * Since configMap is optional, and no found by key -> environment variable no added
    * Case3: configMap exists and found the key
      * `kubectl create configmap optional-configmap --from-literal=akey=Paula`
      * `kubectl create -f podOptionalConfigMap.yaml`
      * `kubectl logs pod-optionalconfigmap`
        * Since configMap and key are found -> environment variable added
    * Case4: Non optional configMap
      * `kubectl create -f podNonOptionalConfigMap.yaml`
        * Since it's non-optional, and it doesn't exist -> not created the pod properly
    * Case5: configMap by reference, no existing
      * `kubectl create -f podOptionalConfigMapByReference.yaml`
      * `kubectl logs pod-optionalconfigmap-byreference`
        * Since configMap is optional, and it doesn't exist -> nothing displayed or existing in that path
    * Case6: configMap by reference exists
      * `kubectl create -f configmap.yaml`
      * `kubectl create -f podOptionalConfigMapByReference.yaml`
      * `kubectl logs pod-optionalconfigmap-byreference`
        * Since configMap exists -> resources existing in that path, displayed
* InvalidEnvironmentVariableNames
  * `kubectl create -f pod.yaml`
    * Pod started
  * `kubectl logs invalid-env-pod`
    * Display environment variables, skipping the invalid one
