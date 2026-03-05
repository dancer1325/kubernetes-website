---
title: Deployment
id: deployment
full_link: /docs/concepts/workloads/controllers/deployment/
short_description: >
  Manages a replicated application on your cluster.

aka: 
tags:
- fundamental
- core-object
- workload
---

* Deployment
  * == API object / 
    * manages -- , through ReplicaSets, -- stateless replicated pods
      * if you need local state -> use [StatefulSet](statefulset.md)
  * == workload resource
  * replacement of [ReplicationController](replication-controller.md)

* replicated pods
  * == SAME container spec 
  * -> any pod can be interchanged
