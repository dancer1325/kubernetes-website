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

----

# Use cases
## ReplicaSet is rollout
* Check 'ReplicaSet is rollout/'
* `kubectl apply -f deploymentReplicaSet.yaml`
* `kubectl get all`
  * Check the basic pattern naming for pods and replicaset
* `kubectl get deployments`
  * 'READY' column
    * ready/desired # of replicas of the application are running in the cluster
  * 'UP-TO-DATE' column
    * number of replicas updated to the desired state
  * 'AVAILABLE' column
    * number of replicas available to your users
* `kubectl rollout status deployment/nginx-deployment`
  * monitor the rollout process of the deployment
    * if all was fine -> will display succeed
* `kubectl get rs`
  * Check the replicaset
  * name = DeploymentName-Hash
    * hash == pods' hash  created by the replicaset!!
* `kubectl get pods --show-labels` & `kubectl get rs --show-labels`
  * show labels
  * 'pod-template-hash'
    * is added by the Deployment controller to the pods and ReplicaSets created by it
    * -> child ReplicaSets of the deployment don't overlap
### Child ReplicaSets
* Update something in the Deployment manifest
* `kubectl get all`
  * take a look about the pods and ReplicaSets
* `kubectl apply -f deploymentReplicaSet.yaml`
* `kubectl rollout status deployment/nginx-deployment`
  * Check that it's succeed
* `kubectl get all`
  * Check that
    * new replicaset is created, although the previous one still exists
      * === child ReplicaSets
    * new pods are created with name following the pattern of the new replicaset

## Pods with a new state
* Check 'Podswithanewstate/'
* `kubectl apply -f deployment.yaml`
* `kubectl get all`
  * Check the pods name and replicaset existing
* Ways to update it
  * Update `spec.template` and `kubectl apply -f deployment.yaml` or 
  * `kubectl set image DeploymentName NameOfTheContainer=NewVersionImage` or
  * `kubectl edit DeploymentName`
* `kubectl rollout status DeploymentName`
* `kubectl get all`
  * Check that
    * new pods have been created
    * there are 2 ReplicaSets
* `kubectl describe DeploymentName`
  * Check the events happened
### Update label selectors
* `kubectl get all`
  * Run to check the previous state of the cluster
* Add
  * `spec.selector.matchLabels` is immutable!!
  * `spec.template.metadata.labels` can be added
  * `kubectl apply -f deployment.yaml`
  * `kubectl get all`
    * Check that 
      * new ReplicaSet is created and the previous one exists till
      * new pods are created based on the new ReplicaSet
* Update value
  * `spec.template.metadata.labels` can be modified
  * `kubectl apply -f deployment.yaml`
  * `kubectl get all`
    * Check that
      * new ReplicaSet is created and the previous one exists till
      * new pods are created based on the new ReplicaSet
* Remove an existing one but there are others matching
  * `kubectl apply -f deployment.yaml`
  * `kubectl get all`
    * Check that
      * create or roll back to ReplicaSet with that configuration
      * new pods are created based on this new ReplicaSet or previous one


## Rollback to an earlier version
* Check 'Rollback to an earlier revision/'
* Let's force a unstable deployment
  * `kubectl apply -f deployment.yaml`
  * `kubectl set image deployment/deployment-to-rollback nginx=nginx:1.161`
    * defined a non-existing image version
  * `kubectl get all`
    * Check that
      * new ReplicaSet is created
      * the corresponding pods aren't running
      * old pods are still running
  * `kubectl rollout status deployment.apps/deployment-to-rollback`
    * live rollout which it's stuck
* `kubectl rollout history deployment.apps/deployment-to-rollback`
  * 'REVISION' column
    * represent the deployment's version
  * 'CHANGE-CAUSE' column
    * copied from the deployment's annotation `kubernetes.io/change-cause`
      * `kubectl describe deployment.apps/deployment-to-rollback` and check if it matches with annotations field
      * you can
        * add the annotation to check `kubectl annotate deployment/deployment-to-rollback kubernetes.io/change-cause="image updated to 1.16.1"`
        * `kubectl rollout history deployment.apps/deployment-to-rollback`
  * `kubectl rollout history deployment.apps/deployment-to-rollback --revision=1`
    * Check details from a specific revision
* `kubectl rollout undo deployment/deployment-to-rollback`
  * Rollback to the previous version
  * `kubectl get all`
    * Check that stuck ReplicaSet is stopped
  * `kubectl describe deployment.apps/deployment-to-rollback`
    * Check the events triggered


## Scale up the deployment
* Check 'Scale up the deployment/'
* `kubectl apply -f deployment.yaml`
* `kubectl scale deployment/deployment-to-scaleup --replicas=10`
  * `kubectl get all`
    * Check that the numbers of replicas is 10
* If HPA is enabled in your cluster
  * `kubectl api-versions | grep autoscaling/v1`
    * if you get an output === HPA is enabled
  * `kubectl autoscale deployment/deployment-to-scaleup --min=10 --max=15 --cpu-percent=80`
    * set up an Autoscaler for your deployment
    * `kubectl get hpa`
      * Check that HPA has been created


## Deployment’s rollout is paused and resumed
* Check 'Deployment’s rollout is paused/'
* `kubectl apply -f deployment.yaml`
* `kubectl rollout pause deployment/deployment-to-rolloutpauseandresume`
* `kubectl set image deployment/deployment-to-rolloutpauseandresume nginx=nginx:1.16.1`
* `kubectl rollout history deployment/deployment-to-rolloutpauseandresume`
  * No new rollout is started
* `kubectl get rs`
  * Confirm that no new ReplicaSet has been created, since the rollout is paused
* `kubectl set resources deployment/deployment-to-rolloutpauseandresume -c=nginx --limits=cpu=200m,memory=512Mi`
  * Make another adjustment
* `kubectl get all` / `kubectl get rs`
  * No updates have been done since the rollout is paused
* `kubectl rollout resume deployment/deployment-to-rolloutpauseandresume`
  * Resume the rollout, applying all the pending adjustments
  * `kubectl rollout history deployment/deployment-to-rolloutpauseandresume` & `kubectl get all`
    * Just 1 new version, merging all the pending changes has been done


## Deployment’s status
* Check 'Deployment’s status/'
* `kubectl apply -f deployment.yaml`
### Progressing
* `kubectl set image deployment/deployment-status nginx=nginx:1.16.1` & run immediately `kubectl get deployment.apps/deployment-status -o yaml`
  * Check that `status.conditions` section contains an entry with ``type: Progressing; status: "True"; reason: ReplicaSetUpdated``
### Complete
* `kubectl get deployment deployment-status -o yaml`
  * Run after the first deployment
  * You can check that `status.conditions` section contains an entry with ``type: Progressing; status: "True"; reason: NewReplicaSetAvailable``
### Fail to progress
* `.spec.progressDeadlineSeconds`
  * time wait for the Deployment Controller to mark the Deployment as stuck
  * `kubectl patch deployment/deployment-status -p '{"spec":{"progressDeadlineSeconds":600}}'`
    * update fields of a resource
#### Insufficient quota
* `kubectl apply -f deploymentInsufficientQuota.yaml`
* `kubectl rollout status deployment.apps/deployment-status-insufficientquota`
  * return that the progress deadline was exceeded
* `kubectl get deployment.apps/deployment-status-insufficientquota -o yaml`
  * You can check that `status.conditions` section contains an entry with ``type: Progressing; status: "False"; reason: ProgressDeadlineExceeded``
#### Readiness probe fails
* `kubectl apply -f deploymentReadinessProbeFail.yaml`
* `kubectl get all`
  * Check that there are not pods AVAILABLE
* `kubectl get deployment.apps/deployment-status-readinessprobefail -o yaml`
  * You can check that `status.conditions` section contains an entry with ``type: Progressing; status: "False"; reason: ProgressDeadlineExceeded``
#### Image pull errors
* `kubectl apply -f deploymentImagePullError.yaml`
* `kubectl get all`
  * Check that there are not pods READY
* `kubectl rollout status deployment.apps/deployment-status-imagepullerror`
  * Waiting for if the 'spec.progressDeadlineSeconds' is exceeded
* `kubectl get deployment.apps/deployment-status-imagepullerror -o yaml`
  * You can check that `status.conditions` section contains an entry with ``type: Progressing; status: "False"; reason: ProgressDeadlineExceeded``
#### Insufficient rights
* TODO: 
#### Limit ranges
* TODO:
#### Application runtime missconfiguration
* TODO:


## Old ReplicaSets are cleaned up
* Check 'OldReplicaSetsAreCleanedUp/'
* `.spec.revisionHistoryLimit`
  * field to specify the number of ReplicaSets to retain for a deployment
  * (by default) 10
* `kubectl apply -f deployment.yaml`
* `kubectl get rs`
  * Pay attention to the existing ReplicaSets
* Apply several changes
  * `kubectl set image deployment/deployment-oldreplicasetscleanup nginx=nginx:1.16.1`
  * `kubectl set image deployment/deployment-oldreplicasetscleanup nginx=nginx:1.15`
  * `kubectl set image deployment/deployment-oldreplicasetscleanup nginx=nginx:1.12`
* `kubectl get rs`
  * You can check that there are non all the generated ReplicaSets
* `kubectl rollout history deployment/deployment-oldreplicasetscleanup`
  * You can check that REVISION 1 is not existing anymore

---

# Strategy
* Check 'strategy/'
## Recreate deployment
* `kubectl apply -f deploymentRecreate.yaml`
* `kubectl get all`
  * Check the existing resources
* `kubectl set image deployment/strategy-recreate nginx=nginx:latest` & run fastly `kubectl get all`
  * You can check that new pods are being created, but previous have been already terminated!!! 
## Rolling update deployment
* `kubectl apply -f deploymentRollingUpdate.yaml`
* `kubectl get all`
  * Check the existing resources
* `kubectl set image deployment/strategy-rollingupdate nginx=nginx:latest` & run fastly `kubectl get all` 
  * You can check that new pods are being created, but previous ones keep on existing!!!
#### maxUnavailable
* `kubectl apply -f deploymentRollingUpdateMaxUnavailable.yaml`
* `kubectl get all`
  * Check the existing resources
* `kubectl set image deployment/strategy-rollingupdate-maxunavailable nginx=nginx:latest` & run fastly `kubectl get all`
  * You can check that new pods are being created, but previous ones keep on existing!!!
#### maxSurge
* `kubectl apply -f deploymentRollingUpdateMaxSurge.yaml`
* `kubectl get all`
  * Check the existing resources
* `kubectl set image deployment/strategy-rollingupdate-maxsurge nginx=nginx:latest` & run fastly `kubectl get all`
  * You can check that new pods are being created, but previous ones keep on existing!!!
#### maxUnavailable + maxSurge 
* `kubectl apply -f deploymentRollingUpdateHybrid.yaml`
* `kubectl get all`
  * Check the existing resources
* `kubectl set image deployment/strategy-rollingupdate-hybrid nginx=nginx:latest` & run fastly `kubectl get all`
  * You can check that new pods are being created, but previous ones keep on existing!!!

---

# Probes
* Check 'tasks/ConfiguredPodsAndContainers/ConfigureLivenessReadinessAndStartupProbes'

---

# HPA
* Check 'tasks/RunApplications/HPA'
* Check 'HPA/' here
* `kubectl apply -f hpa.yaml`
* `kubectl apply -f deployment.yaml`
