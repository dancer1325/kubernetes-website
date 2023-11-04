# no specified
* `kubectl apply -f podWithoutSpecifyingAnEmptyDir.yaml`
* `kubectl describe pods/podWithoutSpecifyingAnEmptyDir`
  * No emptyDir volume created automatically

# specifying
* `kubectl apply -f podSpecifyingAnEmptyDir.yaml`
  * Different pod's containers can share emptyDir volumes
* `kubectl logs pods/pod-emptydir-specified`
  * Check the logs, and no files is stored in volume files' path == empty

# Uses
## scratch space
* `kubectl apply -f uses/scratch\ space/pod.yaml`
  * Problems:
    * Problem1: Why it's started but suddenly "Back-off restarting failed container"
      * Solution: TODO
* `kubectl logs pods/scratch-space-pod -c container-2`
  * Check the logs of the container-2, displaying the temporary information contained in the emptyDir volume
## checkpoints for long computation processes
* `kubectl apply -f uses/checkpoints\ for\ long\ computation\ processes/pod.yaml`
  * Problems:
    * Problem1: "can't open file '//computation.py'"
      * Solution: TODO
    * Problem2: Why it's started but suddenly "Back-off restarting failed container"
      * Solution: TODO
## hold files
* TODO: Create an example