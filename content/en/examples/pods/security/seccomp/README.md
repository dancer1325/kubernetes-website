# Notes
* Kind
  * runs Kubernetes in Docker ->
    * each cluster's node == container
      * load files onto a node == mount files in the container's file system 

# Run kind cluster with seccomp profiles
* `kind create cluster --config=kind.yaml`
  * Create the cluster with this configuration
* `docker ps`
  * Identify the Docker container running the cluster's node
* `docker exec -it ContainerName ls /var/lib/kubelet/seccomp/profiles`
  * Check that the "profiles/" directory has been successfully loaded into the seccomp path of the kubelet
    * === seccomp profiles are available to the kubelet running within kind

# Create a Pod that uses the container runtime default seccomp profile
* Most of CR provide a default syscalls
  * syscalls
    * [Reference link](https://man7.org/linux/man-pages/man2/syscalls.2.html)
* It can be got, setting `securityContext.seccompProfile` = `RuntimeDefault`
* `kubectl apply -f https://k8s.io/examples/pods/security/seccomp/ga/default-pod.yaml`
  * Problems:
    * Problem1: "runtime: failed to create new OS thread (have 2 already; errno=22)"
      * Solution: Update image to 1.0, and run here `kubectl apply -f ga/default-pod.yaml`


# Create a Pod with a seccomp profile for syscall auditing
* `kubectl apply -f https://k8s.io/examples/pods/security/seccomp/ga/audit-pod.yaml`
  * Problems:
    * Problem1: "runtime: failed to create new OS thread (have 2 already; errno=22)"
      * Attempt1: Update image to 1.0
      * Solution: ?
* `kubectl expose pod audit-pod --type NodePort --port 5678`
  * Create a NodePort service to access inside the control plane container
  * `kubectl get service audit-pod`
    * Check in the column PORT, with the structure <NodeIP>:<NodePortNumber>
* `docker exec -it DockerNameOfTheKindClusterControlPlane curl localhost:NodePortNumber`
  * Make curl request into the cluster to reach the service
  * The outputs should match with the ones done in your local computer, since the syscalls belong to your own computer
    * `tail -f /var/log/syslog | grep 'http-echo'`
      * Run in our own computer that it matches with the previous outputs

# Create a Pod with a seccomp profile for syscall violated
* `kubectl apply -f https://k8s.io/examples/pods/security/seccomp/ga/violation-pod.yaml`
