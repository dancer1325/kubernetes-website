# Steps
* `kubectl get storageclass`
  * Check the available storage classes in your cluster. Based on that update storageClassName in the PVC
  * If you don't find a storageClass with name 'standard' -> `kubectl apply -f storageclass.yaml`
* `kubectl apply -f pvc.yaml`
* `kubectl apply -f pod.yaml`
* `kubectl exec -it pods/my-pod -- sh`
  * `cd /data`
  * `pwd`
    * Check that this volume is mounted in this path 
