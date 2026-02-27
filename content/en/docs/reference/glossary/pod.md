---
title: Pod
id: pod
full_link: /docs/concepts/workloads/pods/
short_description: >
  A Pod represents a set of running containers in your cluster.

aka: 
tags:
- core-object
- fundamental
---
 
* pod
  * == Kubernetes object
    * smallest
    * simplest
  * == set of running [containerS](container.md) | your cluster
    * containerS
      * NORMALLY, 1
      * if you want to run > 1 container == sidecar containers
        * _Example of use cases:_ logging
  * managed NORMALLY -- by a -- [deployment](deployment.md)
