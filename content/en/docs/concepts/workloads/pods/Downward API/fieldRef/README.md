Certain pod's fields can just be accessed via a determined way

# Via environment variables + downward API volume files
* Via environment variables
  * `kubectl create -f Via\ environment\ variables\ +\ downward\ API\ volume\ files/podViaEnvironmentVariables.yaml`
  * `kubectl logs pods/podfields-bothfields-viaenv -c container-bothfields-viaenv`
    * Check that the fields are added as environment variables
* Via downward API volume files
  * `kubectl create -f Via\ environment\ variables\ +\ downward\ API\ volume\ files/podViaDownwardAPIVolumeFile.yaml`
  * `kubectl logs pods/podfields-bothfields-viadownwardvolumefiles -c container-bothfields-viadownwardvolumefiles`
    * Check that the fields are added as volume files

# Via environment variables
* `kubectl apply -f Via\ environment\ variables/pod.yaml`
  * Problems: 
    * Problem1: Why it's started but suddenly "Back-off restarting failed container"
* `kubectl exec podfields-viaenv -c container-viaenv -- env`
  * Get into the pod, and run a command `env` directly in the specific container
  * Problems:
    * Problem1: It doesn't work since the container is continuously restarting
* `kubectl logs pods/podfields-viaenv`
  * Check the logs of the pod, with the env displayed due to the command configuration. You can check, our desired environment variables

# Via downward API volume files
* `kubectl apply -f Via\ downward\ API\ volume\ files/pod.yaml`
  * Problems:
    * Problem1: Why it's started but suddenly "Back-off restarting failed container"
    * Solution: TODO
* `kubectl exec podfields-viadownwardvolumefiles -c container-viadownwardvolumefiles -- env`
  * Get into the pod, and run a command `env` directly in the specific container
  * Problems:
    * Problem1: It doesn't work since the container is continuously restarting
* `kubectl logs pods/podfields-viadownwardvolumefiles -c container-viadownwardvolumefiles`
  * Check the logs of the container, with the added by the entry command of the container