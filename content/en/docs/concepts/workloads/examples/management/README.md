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

# TODO:

# Updating your application WITHOUT an outage
* goal
  * how to create & update applications | deployments

* steps
  * `kubectl create deployment my-nginx --image=nginx:1.14.2`
  * `kubectl scale --replicas 1 deployments/my-nginx`
    * ensure there is 1! replica
  * `kubectl patch --type='merge' -p '{"spec":{"strategy":{"rollingUpdate":{"maxSurge": "100%" }}}}'`
    * add MORE temporary replicas | rollout
  * `kubectl edit deployment/my-nginx` & change the manifest
    * -> deployment update it PROGRESSIVELY AUTOMATICALLY

## how to manage rollouts?
### | deployment
* `kubectl apply -f my-deployment.yaml`
* `kubectl rollout status deployment/my-deployment --timeout 10m # 10 minute timeout`
  * wait for rollout PREVIOUSLY -- to -- finish

### | stateful
* `kubectl apply -f backing-stateful-component.yaml`
* `kubectl rollout status statefulsets/backing-stateful-component --watch=false`
  * NOT wait for rollout to finish, just check the status

# TODO:
