# Content
* PodAndConfigmapInDifferentNamespaces
  * `kubectl apply -f namespaceA.yaml`
  * `kubectl apply -f namespaceB.yaml`
  * `kubectl create -f configmap.yaml`
  * `kubectl create -f pod.yaml`
    * Apparently no error as output
  * `kubectl get pods -n namespace-b`
    * Check the pods in namespace-b
    * You realize that the previous pod is in "CreateContainerConfigError" status
* StaticPodsWithConfigMap
  * Check the own README.md