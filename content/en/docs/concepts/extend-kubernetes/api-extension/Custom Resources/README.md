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

# Custom Resources
* Structured data
  * _Example:_  For CRD, you have got all freedom given by OpenAPIv3 to define the schema to use in the CRD -- `spec.versions[x].schema.openAPIV3Schema`

# Ways to handle Custom Resources
## via Kubernetes API Aggregation Layer
* Check 'API Aggregation Layer/'
## via CRD
* `kubectl apply -f crd.yaml`
  * CRD is created without any programming 
  * `kubectl get crd` / `kubectl get crd -o wide` 
    * Check the creation
  * `kubectl apply -f instanceOfCRD.yaml`
    * Create an instance of the CRD
    * `kubectl get myapps` 
      * Check the instances created of the CRD
      * 'myapps' because it's the CRD's field `spec.names.plural`
* `kubectl get pods -n kube-system | grep 'apiserver'`
  * Just the already existing API servers -- no new one's as API Aggregation --
### use cases
* Imperative API
  * Check 'go.mod' and 'main.go'
  * `go build -o main`
  * `go run .`
  * `kubectl get pods` 
    * Check that a pod is created programmatically / imperative  
* kubectl support not required
  * Although, you have kubectl commands to interact with them
    * `kubectl get crd`
    * `kubectl describe crd myapps.example.com` -- CRDName --
    * `kubectl apply -f pathInWhichWeAddedCRD`
    * `kubectl describe myapps.example.com myapp-instance`
      * Describe the CRD's instance 
    * `kubectl edit myapps.example.com myapp-instance` / `kubectl delete myapps.example.com myapp-instance`
      * Edit / Delete CRD's instance
* Kubernetes UI support not required
  * Check '/tasks/accessApplicationCluster/deployAndAccessTheKubernetesDashboard' to enable Kubernetes UI
  * Once you login, you don't find anything about CRD or their instances in some namespace
* Already existing API
  * Check 'concepts/overview/TheKubernetesAPI/'
  * POST & GET -- '/apis/{group}/{version}/{plural}' -- for CRD
    * 'http://localhost:8001/apis/example.com/v1' / 'http://localhost:8001/apis/example.com/v1/myapps'
  * GET & PUT & PATCH & DELETE --'/apis/{group}/{version}/{plural}/{name}' -- for CRD's instances
    * `http://localhost:8001/apis/example.com/v1/myapps/myapp-instance` -- TODO: Why doesn't it work? --
### features
* Validations
  * OpenAPIv3 schema
    * `kubectl apply -f instanceOfCRDNotValidByOpenAPISchema.yaml` you get an error due to non-valid specified field
  * Validation Ratcheting
    * TODO: 
  * ValidatingAdmissionWebhook
    * TODO:
* Defaulting
  * via OpenAPI v3 schema default keyword in the CRD -- `spec.versions[x].schema.openAPIV3Schema.properties.spec.properties.propertyKey.default` --
  * via MutationAdmissionWebhook
    * TODO:
* Multiversioning
  * Check 'multiversioning/' folder for the samples
  * `kubectl apply -f crdMultiVersioning.yaml`, ` kubectl apply -f crdMultiversionV1.yaml` and `kubectl apply -f crdMultiversionV2.yaml`
  * `kubectl describe multiversioning.example.com`
    * TODO: Why multi-instance1 has 'API Version:  example.com/v1beta1' ?

