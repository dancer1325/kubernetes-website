# Pre requisites
* Locally or cluster you fully control (=== configure control plane), where you can
  * Run some docker daemon
    * [Docker desktop](https://www.docker.com/products/docker-desktop/)
  * Install some local cluster
    * tool
      * [minikube](https://minikube.sigs.k8s.io/docs/start/)
      * [kind](https://kind.sigs.k8s.io/)
  * Run a local cluster
    * [minikube]  `minikube start`
    * [kind] `kind create cluster`
  
# Sample1
* Check 'sample1/'
* Projected volume with volumes
  * secret
  * downward api
    * defined itself in the pod template
  * configmap
* `kubectl apply -f secret.yaml`
* `kubectl apply -f configMap.yaml`
* `kubectl apply -f pod.yaml`
* `kubectl exec -it pods/volume-test -- sh`
  * `cd projected-volume/`
  * `ls`
    * Check that 
      * secret.items[0].path is specified
  * About downwardAPI
    * `cat labels`
      * Confirm that it matches with metadata.labels
    * `cat cpu_limit`
      * Confirm that it matches with containers[x].resources.limits.cpu
      * Problems:
        * Problem1: Why does it return '1' instead of the configured one?
  * About configMap
    * `cat my-group/my-config`
      * matches with the value configured in the configMap
  * About secret
    * `my-group/my-username`
      * matches with the value configured in the secret
      * Problems:
        * Problem1: Why it's not displayed the value?

# Sample2
* Check 'sample2/'
* Projected volume with volumes
  * secret with non-default permission mode set


# Notes
* File permissions in Unix-like OS
  * Owners'permissions Group'permissions Others'permissions
    * X'spermissions can be represented in
      * [mode](https://www.nexcess.net/help/what-is-chmod/) 
        * octal or
        * integer
      * string
* `kubectl apply -f secret1.yaml`
* `kubectl apply -f secret2.yaml`
* `kubectl apply -f pod.yaml`
* `kubectl exec -it pods/volume-test-2 -- sh`
  * `cd /projected-volume/my-group`
  * `cat my-username`
    * Problems:
      * Problem1: Why it's not displayed the value?
  * `cat my-password`
    * Problems:
      * Problem1: Why it's not displayed the value?