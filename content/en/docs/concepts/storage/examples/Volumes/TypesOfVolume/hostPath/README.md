# simple Pod
* Simple pod to check hostPath functionality
* `kubectl apply -f simplePod.yaml`
  * Problems:
    * Problem1: "runtime: failed to create new OS thread (have 2 already; errno=22)
      fatal error: runtime.newosproc"
      * Attempt1: Specify `resources.limits.cpu` & `resources.requests.cpu`
      * Attempt2: Restart docker daemon
      * Solution: Switch to another image
    * Problem2: It doesn't work since the container is continuously restarting
      * Attempt1: Add readiness and liveness probes
      * Attempt2: Adjust resources
      * Attempt3: Switch to `kubectl create -f`
      * Attempt4: Add ports
      * Attempt5: Add specific sha256 code to the image to pull
      * Solution: Add `command ["sleep", "3600"]`
* `kubectl exec -it test-pod -c test-container -- sh`
  * get access into the container
  * `ls test-pd/`
    * You should see all the host path's content
    * Problems:
      * Problem1: Why no content of the hostPath is displayed there?
        * Attempt1: `readOnly: true`
        * Attempt2: `kubectl exec -it test-pod -c test-container sudo -- sh` to get access
        * Solution: TODO ?

# Cluster with several nodes on different hosts
* Problems:
  * Problem1: How to create cluster with different nodes on different hosts?
    * Attempt1: Cluster with several nodes `kind create cluster --config=clusterWithSeveralNodes/kindCluster.yaml`, but on the same host

# Scope to specific files or directories
* TODO

# Type field
## empty
## DirectoryOrCreate
* `kubectl apply -f type\ field/directoryorcreatepod.yaml`
  * Problems:
    * Problem1: That directory should have been created in node's host ?
      * Attempt1: `readOnly: false`
* `kubectl describe pods/pod-directoryorcreate`
  * You can check in the volumes the host path
## Directory
* `kubectl apply -f type\ field/directorypod.yaml`
  * Not properly started up, because the directory in the host doesn't exist
## FileOrCreate
* `kubectl apply -f type\ field/fileorcreatepod.yaml`
* `kubectl exec -it pod-fileorcreate -c container-fileorcreate -- sh`
  * `cat test-pd`
    * Checking that it's an empty file
* `kubectl describe pods/pod-directoryorcreate`
  * You can check in the volumes the host path
## File
* `kubectl apply -f type\ field/filepod.yaml`
  * Not properly started up, because the file in the host doesn't exist
## Socket
* `kubectl apply -f type\ field/socketpod.yaml`
  * Not properly started up, because the socket in the host doesn't exist
## CharDevice
* `kubectl apply -f type\ field/chardevicepod.yaml`
  * Not properly started up, because the char device in the host doesn't exist
## BlockDevice
* `kubectl apply -f type\ field/blockdevicepod.yaml`
  * Not properly started up, because the block device in the host doesn't exist
