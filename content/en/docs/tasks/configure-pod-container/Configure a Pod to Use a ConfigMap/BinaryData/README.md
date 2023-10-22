# Content
* 'UsingDataNeitherASCIINorUTF8/'
  * `kubectl apply -f configmap.yaml`
  * `kubectl describe configmaps configMapName`
* 'NeitherASCIINorUTF8/'
  * `kubectl apply -f configmap.yaml`
  * `kubectl get configmap -o jsonpath='{.binaryData}' configMapName`
    * Check binaryData's entries