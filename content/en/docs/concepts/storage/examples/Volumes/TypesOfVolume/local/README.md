# How to use a local volume?
* Just possible via Persistent Volume
  * `kubeclt apply -f persistentVolume.yaml`
  * `kubeclt apply -f persistentVolumeClaim.yaml`
  * `kubectl apply -f pod.yaml`
    * Problems:
      * Problem1: "1 node(s) had untolerated taint {node-role.kubernetes.io/control-plane: }. preemption: 0/2 nodes are available: 1 No preemption victims found for incoming pod, 1 Preemption is not helpful for scheduling."
        * Solution1: Switch nodeAffinity to 'kind-worker'
      * 