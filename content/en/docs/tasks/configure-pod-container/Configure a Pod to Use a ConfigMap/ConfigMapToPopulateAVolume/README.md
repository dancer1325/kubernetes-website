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