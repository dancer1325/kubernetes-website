# Steps
* `kubectl apply -f secret.yaml`
* `kubectl apply -f pod.yaml`
* `kubectl exec -it pods/my-app-pod -- sh`
  * `cd etc/secrets`
  * `ls`
    * Check that this path contain files with the secrets' data entries 
