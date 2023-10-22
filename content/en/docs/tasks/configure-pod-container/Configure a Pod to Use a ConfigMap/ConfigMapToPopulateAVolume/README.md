# Content
* Simple
  * `kubectl create -f configmap.yaml`
  * `kubectl create -f pod.yaml`
  * `kubectl logs pod-configmap-populate-volume`
    * Check the logs of the pod, displaying the files -- which are the configMap's data --
* Specify the path
  * `kubectl create -f configmap.yaml`
  * `kubectl create -f pod.yaml`
    * Problems
      * Problem1: If you run this command directly without running the first one previously -> ' is waiting to start: ContainerCreating'
        * Solution: Run the previous command
  * `kubectl logs pod-configmap-populate-volume-specifypath`
    * Check the value of the configMap's data in that path is displayed
* If you update a configMap mounted in a volume -> there is a delay to project to the pod
  * `kubectl apply -f configmap.yaml`
  * `kubectl apply -f pod.yaml`
  * `kubectl exec -it pod-withadelay-updatingavolume -- cat /etc/config/app.properties`
  * If you update the configMap's data & `kubectl apply -f configmap.yaml` & `kubectl exec -it pod-withadelay-updatingavolume -- cat /etc/config/app.properties`
    * if you are fast enough (< delay time) -> the change is not projected in the pod !!
    * if you leave a time (> delay time) -> the change is projected in the pod !!
* If you update a configMap mounted in a volume's subpath -> there is a delay to project to the pod
  * `kubectl apply -f configmap.yaml`
  * `kubectl apply -f pod.yaml`
  * `kubectl exec -it pod-withadelay-updatingavolume-subpath -- cat /etc/config/app.properties`
  * If you update the configMap's data & `kubectl apply -f configmap.yaml` & `kubectl exec -it pod-withadelay-updatingavolume-subpath -- cat /etc/config/app.properties`
    * Independently the time spent -> the change is not projected in the pod !!
* simpleViaDataAndBinaryData
  * `kubectl create -f configmap.yaml`
  * `kubectl create -f pod.yaml`
  * `kubectl logs pod-configmap-populate-volume-viadataandbinarydata`
    * Check the logs of the pod, displaying the files -- which are the configMap's data and binaryData --