---
title: ReplicaSet
id: replica-set
full_link: /docs/concepts/workloads/controllers/replicaset/
short_description: >
 ReplicaSet ensures that a specified number of Pod replicas are running at one time

aka: 
tags:
- fundamental
- core-object
- workload
---

* ReplicaSet
  * == workload resource
  * goal
    * maintain a set of replica pods running | ANY given time
  * uses
    * by [deployment](deployment.md)

* Child ReplicaSets
  * := ReplicaSets / created by a deployment 
    * | rollout a deployment -> NEW ReplicaSet
