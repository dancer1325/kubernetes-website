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


# Ways to interact with the Kubernetes API
## kubectl
* `kubectl version`
* Check 'Reference/kubectl'
## kubeadm
* Check 'GettingStarted/ProductionEnvironment/InstallingKubernetesWithDeploymentTools/BootstrappingClustersWithKubeadm'
## REST calls
* Ways
  * Via a proxy
    * `kubectl proxy --port=8001` & `curl -k http://localhost:8001/` or open in the browser 'http://localhost:8001'
      * checking all the available paths
  * Via API server directly
    * `curl -k -X GET https://kube-api-serverIP`
      * Problems: Time out
        * Solution: TODO: ?


# OpenAPI
* You can access via any of the previous one's, better via REST calls
* v2
  * Via a proxy
    * `kubectl proxy --port=8001`
    * `curl -k http://localhost:8001/openapi/v2` or open in the browser 'http://localhost:8001/openapi/v2'
      * checking the v2
  * Via API server directly
    * `curl -k -X GET https://kube-api-serverIP/openapi/v2`
      * Problems: Time out
        * Solution: TODO: ?
  * request's headers which can be customized
    * 'Accept'
      * `curl -k -H "Accept:application/json" http://localhost:8001/openapi/v2`
        * 'application/json' is the default value
      * `curl -k -H "Accept:application/com.github.proto-openapi.spec.v2@v1.0+protobuf" http://localhost:8001/openapi/v2`
        * 'application/com.github.proto-openapi.spec.v2@v1.0+protobuf' for intra-cluster use
      * `curl -k -H "Accept:*" http://localhost:8001/api`
        * '*' serves application/json
    * 'Accept-Encoding'
      * `curl -k -H "Accept-Encoding:gzip" http://localhost:8001/openapi/v2 > out.json`
        * Encoded the response
* v3
  * Via a proxy
    * `kubectl proxy --port=8001` & `curl -k http://localhost:8001/openapi/v3` or open in the browser 'http://localhost:8001/openapi/v3'
      * checking the v3 
  * Via API server directly
    * `curl -k -X GET https://kube-api-serverIP/openapi/v3`
      * Problems: Time out
        * Solution: TODO: ?
  * same request's headers can be customized


# API Discovery
* Check API groups and versioning
  * `/api` and `/apis`
* Check list of resources per group and version
  * `/apis/<group>/<version>`
## Aggregated Discovery
* Requirements
  * kubernetes v1.27+
* All resources are accessed filtering by group and version passing by header and hitting `/api` or `/apis`
  * `curl -k -H "Accept: application/json;v=v2beta1;g=apidiscovery.k8s.io;as=APIGroupDiscoveryList" http://localhost:8001/api` 
    * You can select any other group and version
    * '8001' is because I am using kube proxy
