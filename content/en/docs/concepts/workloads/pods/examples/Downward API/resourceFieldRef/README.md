All container's fields can be accessed via both mechanisms

# Via environment variables
* `kubectl apply -f podViaEnvironmentVariables.yaml`
  * Problems:
    * Problem1: Why it's started but suddenly "Back-off restarting failed container"
    * Solution: TODO
* `kubectl logs pods/pod-resourcefieldref-viaenv`
  * Check the logs, displaying our desired environment variables

# Via downward API volume files
* `kubectl apply -f podViaDownwardAPIVolumeFile.yaml`
  * Problems:
    * Problem1: Why it's started but suddenly "Back-off restarting failed container"
    * Solution: TODO
* `kubectl logs pods/pod-resourcefieldref-viadownwardvolumefiles`
  * Check the logs, displaying our desired environment variables