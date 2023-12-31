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

---
# NodeSelector -- podspec.nodeselector — node.metadata.labels --
* Check all the required sample manifests in 'nodeSelector/'
* Check that the pods are assigned to the nodes, matching the labels -- 'cluster-for-nodeselector' --
  * `kind create cluster --config=cluster.yaml`
    * `kubectl get node cluster-for-nodeselector-worker -o json | jq '.metadata.labels'` to check the label added
  * `kubectl apply -f pod.yaml` & `kubectl get pod/nodeselector -o wide` to check that 'nodeselector' runs on the expected node
* Cases
  * If several nodes matching the nodeselector == same score -> kube-scheduler selects one randomly
    * `kind create cluster --config=clusterWithSeveralMatching.yaml`
      * `kubectl get node cluster-for-nodeselector-worker -o json | jq '.metadata.labels'` to check the labels added
    * `kubectl apply -f pod.yaml` & `kubectl get pod/nodeselector -o wide` to check that 'nodeselector' runs on the expected node
  * If there is no node matching the nodeselector -> kube-scheduler leaves on 'Pending' the pod
    * `kind create cluster` & `kubectl apply -f pod.yaml` & `kubectl get pods -o wide` & `kubectl describe pod/nodeselector` to confirm it

# Affinity & anti-affinity
* Check all the required sample manifests in 'Affinity & anti-affinity/'
## Node affinity
* Check in 'Node affinity/'
### Subtypes
* '.requiredDuringSchedulingIgnoredDuringExecution'
  * `kind create cluster --config=clusterRequired.yaml` & `kubectl apply -f pods.yaml`
  * Required fulfil the rule to assign the pod -- 'required-1' --
    * Check that it's assigned to the specific node
      * `kubectl get nodes --show-labels` to check the node name & `kubectl get pod/required-1 -o wide`
    * If the rule is NOT fulfilled by none of the nodes -> pod not assigned -- 'required-2' --
      * `kubectl get pod/required-2 -o wide` to check that it's 'Pending'
  * 'IgnoredDuringExecution' because if you change the node labels after assigning ('required-1') -> pod keeps on running  
    * `kubectl label node cluster-1-worker2 required.1/1.1-` to remove it & `kubectl get pod/required-1 -o wide` to check is assigned to that node
  * If there are several terms in `.nodeSelectorTerms` → just required to fulfil 1 to assign -- 'required-3' --
    * `kind create cluster --config=clusterRequired.yaml` & `kubectl apply -f pods.yaml`
    * `kubectl get node cluster-1 -o json | jq '.metadata.name'` to check the metadataName added
    * `kubectl get pod/required-3 -o wide` to check that it's assigned to first worker node
* '.preferredDuringSchedulingIgnoredDuringExecution'
  * If there is no node to fulfill the rule -> it's assigned to some node -- 'preferred-1' --
    * `kubectl apply -f pods.yaml` & `kubectl get pod/preferred-1 -o wide` to check that it's assigned to a node
  * If there are several preferred rules with weights ->  higher weigh more taken in account to assign to a node -- 'preferred-2' --
    * `kubectl apply -f pods.yaml` & `kubectl get pod/preferred-2 -o wide` to check that it's assigned to the worker 2
  * 'IgnoredDuringExecution' because if you change the node labels after assigning ('preferred-1') -> pod keeps on running
    * `kubectl label node cluster-1-worker2 preferred.2/1.2-` to remove the label & `kubectl get pod/preferred-2 -o wide` to check is assigned to that node
* Both subtypes
  * if there are several expressions in `.matchExpressions`  → required to fulfil ALL to assign
    * for '.requiredDuringSchedulingIgnoredDuringExecution' -- 'both-severalmatchexpressions-1' --
      * `kubectl apply -f pods.yaml` & `kubectl get pod/both-severalmatchexpressions-1 -o wide` to check that it's NOT assigned to any node and it's in 'Pending' STATUS
    * for '.preferredDuringSchedulingIgnoredDuringExecution' -- 'both-severalmatchexpressions-2' --
      * `kubectl apply -f pods.yaml` & `kubectl get pod/both-severalmatchexpressions-2 -o wide` to check that it's assigned to some node, since it's 'preferred'
  * can be defined at the same time -- 'both-subtypes' --
    * `kind create cluster --config=clusterBothSubTypes.yaml` & `kubectl apply -f pods.yaml` 
## Node anti-affinity
* Ways
  * via NOT operators
    * `NotIn`  -- 'anti-affinity-notin' --
      * `kubectl apply -f pods.yaml` & `kubectl get pod/anti-affinity-notin -o wide` to check that it's assigned to 'worker-1', the only one without the required label
    * `DoesNotExist`  -- 'anti-affinity-doesnotexist' --
      * `kubectl apply -f pods.yaml` & `kubectl get pod/anti-affinity-doesnotexist -o wide` to check that it's assigned to 'worker-1', the only one without the required label
  * via Taints and Tolerations
    * Check '../TainsAndTolerations/'
## Inter-pod affinity/anti-affinity
* Check in 'Inter-pod affinity & anti-affinity/'
### `topologyKey`
* If the topologyKey is NOT matched -> pod NOT assigned to any node -- 'topologykey-1' --
  * `kind create cluster --config=cluster.yaml` & `kubectl apply -f pods.yaml` & `kubectl get pod/topologykey-1 -o wide` to check that it's in 'Pending' status and not assigned
* If the topologyKey is matched to some node's labels -> pod assigned to a node -- 'topologykey-1-1' & 'topologykey-1' --
  * `kind create cluster --config=clusterMatchingTopologyKey.yaml` & `kubectl apply -f podMatchingFirst.yaml` & `kubectl get pod/topologykey-1-1 -o wide` to check that it's assigned the last worker node
  * `kubectl apply -f pods.yaml` & `kubectl get pod/topologykey-1 -o wide` to check that it's assigned to last worker node
* If different topologyKeyS for different podAffinity are matched to some node's labels -> pod assigned to a node -- 'topologykey-2' & 'topologykey-2-1' --
  * `kubectl apply -f podMatchingFirst.yaml` & `kubectl get pod/topologykey-2-1 -o wide` to check that it's assigned the last worker node
  * `kubectl apply -f pods.yaml` & `kubectl get pod/topologykey-2 -o wide` to check that it's assigned to last worker node
### Inter-pod anti-affinity
* Ways
  * via NOT operators
    * `NotIn`  -- 'anti-affinity-notin' & 'anti-affinity-notin-1' --
      * `kubectl apply -f podMatchingFirst.yaml` & `kubectl get pod/anti-affinity-notin-1 -o wide` to check that it's assigned the last worker node
      * `kubectl apply -f pods.yaml` & `kubectl get pod/anti-affinity-notin -o wide` to check that it's assigned to last worker node
    * `DoesNotExist`  -- 'anti-affinity-doesnotexist' & 'anti-affinity-doesnotexist-1' --
      * `kubectl apply -f podMatchingFirst.yaml` & `kubectl get pod/anti-affinity-doesnotexist-1 -o wide` to check that it's assigned the last worker node
      * `kubectl apply -f pods.yaml` & `kubectl get pod/anti-affinity-doesnotexist -o wide` to check that it's assigned to last worker node
### Subtypes
* 'requiredDuringSchedulingIgnoredDuringExecution' -- 'topologykey-1' & 'topologykey-2' & 'anti-affinity-notin' & 'anti-affinity-doesnotexist' --
* 'preferredDuringSchedulingIgnoredDuringExecution' -- 'topologykey-1' & 'topologykey-2' --
### `.namespaces`
* List of namespaces to match against -- 'namespace-1-1' & 'namespace-1' --
  * `kubectl apply -f namespace.yaml` 
  * `kubectl apply -f podMatchingFirstWithNamespace.yaml` & `kubectl get pod/namespace-1-1 -n interpod-namespace-2 -o wide` to check that it just exists in this namespace and running in the last worker node
  * `kubectl apply -f podsWithNamespace.yaml` & `kubectl get pod/namespace-1 -o wide` to check that it's created in the default namespace and run onto last worker node
* If it's empty / omitted -> namespace of the pod where the podAffinity is defined -- 'namespace-2' & 'namespace-1-1' --
  * `kubectl apply -f podMatchingFirstWithNamespace.yaml` & `kubectl get pod/namespace-1-1 -n interpod-namespace-2 -o wide` to check that it just exists in this namespace and running in the last worker node
  * `kubectl apply -f podsWithNamespace.yaml` & `kubectl get pod/namespace-2 -n interpod-namespace-2 -o wide` to check that it's created in the 'interpod-namespace-2' namespace and run onto last worker node
### `.namespaceSelector`
* Allows querying over the set of namespaces -- 'namespaceselector-1' & 'namespaceselector-1-1' --
  * `kubectl apply -f podMatchingFirstWithNamespace.yaml` & `kubectl get pod/namespaceselector-1-1 -n interpod-namespaceselector-1 -o wide` to check that it just exists in this namespace and running in the last worker node
  * `kubectl apply -f podsWithNamespace.yaml` & `kubectl get pod/namespaceselector-1 -o wide` to check that it's created in the default namespace and run onto last worker node
* If it's empty -> Checks for ALL namespaces -- 'namespaceselector-3' & 'namespace-1-1' & 'namespaceselector-1-1' --
  * `kind delete clusters cluster-3` & `kind create cluster --config=clusterMatchingTopologyKey.yaml`
  * `kubectl apply -f namespace.yaml` & `kubectl apply -f podMatchingFirstWithNamespace.yaml` & `kubectl get pods` to check that there are no pods in the default namespace
  * `kubectl apply -f podsWithNamespace.yaml` & `kubectl get pod/namespaceselector-3 -o json` is 'Running' and to check that the namespaceselector is empty 
* If it's null -> Checks for namespace of the pod defined the affinity -- 'namespaceselector-2' & 'namespace-1-1' & 'namespaceselector-1-1' --
  * `kind delete clusters cluster-3` & `kind create cluster --config=clusterMatchingTopologyKey.yaml`
  * `kubectl apply -f namespace.yaml` & `kubectl apply -f podMatchingFirstWithNamespace.yaml` & `kubectl get pods` to check that there are no pods in the default namespace
  * `kubectl apply -f podsWithNamespace.yaml` & `kubectl get pod/namespaceselector-2` is 'Pending' since there is no Pod matching the rules in the default namespace
* If you select either `.namespaces` or `.namespaceSelector` -> all namespaces indicated will be used to filter in -- 'namespaceselector-4' & 'namespaceselector-1-1' --
  * `kind delete clusters cluster-3` & `kind create cluster --config=clusterMatchingTopologyKey.yaml`
  * `kubectl apply -f namespace.yaml` & `kubectl apply -f podMatchingFirstWithNamespace.yaml` & `kubectl get pods` to check that there are no pods in the default namespace
  * `kubectl apply -f podsWithNamespace.yaml` & `kubectl get pod/namespaceselector-4 -o json` is 'Running' and to check that all namespaces indicated are used to filter in
### `.matchLabelKeys`
* Requirements
  * Kubernetes v1.29 &
  * enable the feature gate MatchLabelKeysInPodAffinity
* `kind create cluster --config=clusterMatchingTopologyKeyAndEnableFeatureGates.yaml` & `kubectl describe pod/kube-apiserver-cluster-4-control-plane -n kube-system` to check that it's passed via command the feature gate
* once pod affinity rules are satisfied → specifying keys for the labels  -- 'matchlabelkeys-1' & 'matchlabelkeys-1-1' --
  * `kubectl apply -f namespace.yaml` & `kubectl apply -f podMatchingFirstWithNamespace.yaml`
  * `kubectl apply -f podsWithMatchLabelKeys.yaml` & `kubectl get pod/matchlabelkeys-1 -o wide` to check that it's 'Running' and onto last worker node
* If you try to add a 'matchLabelKeys' entry which doesn't exist after all pods with affinity -> error through API -- 'matchlabelkeys-2' --
  * `kubectl apply -f podsWithMatchLabelKeys.yaml`
### Use cases
* Co-locate set of workload resources in the same defined topology
  * Check 'useCase/'
  * `kind create cluster --config=cluster.yaml` & `kubectl get nodes --show-labels` to check that all have got 'kubernetes.io/hostname' label
    * Cluster with >3 nodes
  * `kubectl apply -f firstDeployment.yaml` & `kubectl get pods -o wide` to check that all pods about the firstDeployment live in different cluster's nodes
  * `kubectl apply -f secondDeployment.yaml` & `kubectl get pods -o wide` to check that all pods about the secondDeployment live in same cluster's nodes that firstDeployment
## Operators
* Valid for Node Affinity & Inter-pod affinity
  * `In`
  * `NotIn`
  * `Exists`
  * `DoesNotExist`
* Valid for Node Affinity
  * `Gt`  -- 'operator-gt' --
    * `kind create cluster --config=clusterRequired.yaml` & `kubectl get nodes --show-labels` to confirm all the labels added
    * `kubectl apply -f pods.yaml` & `kubectl get pod/operator-gt -o wide` to confirm that it's running on the node 'other-na'
  * `Lt` -- 'operator-lt' --
    * `kubectl apply -f pods.yaml` & `kubectl get pod/operator-lt -o wide` to confirm that it's running on the last worker node

# NodeName -- podspec.nodeName - node.name -- 
* Check 'nodeName/'
* `podspec.nodeName` - node.name  -- 'cluster-namingnode' --
  * `kind create cluster --config=cluster.yaml` & `kubectl get nodes` to check the specific name to the node
  * `kubectl apply -f pods.yaml` & `kubectl get pod/nodename -o wide` to confirm that it's assigned to 'other-na' node
* If it’s not empty →
  * override NodeSelector & Affinity Rules -- 'nodename-override-nodeselectorandaffinity' --
    * `kubectl apply -f pods.yaml` & `kubectl get pod/nodename-override-nodeselectorandaffinity -o wide` to confirm that it's assigned to 'other-na' node, but it's NOT 'Running'
  * & named node does NOT exist → pod will NOT run -- 'nodename-notexist' --
    * `kubectl apply -f pods.yaml` & `kubectl get pod/nodename-notexist -o wide` to confirm that it's in 'Pending' because not found any node
  * & node has NOT got the enough resources → pod fails  -- 'nodename-notenoughresources' --
    * `kubectl describe node other-na` to check the allocated resources for the node
    * `kubectl apply -f pods.yaml` & `kubectl get pod/nodename-notenoughresources -o wide` to confirm that it's in 'OutOfcpu' because the node has NOT got enough resources to allocate it
    * 

# Pod topology spread constraints
* Check '../Pod Topology Spread Constraints'