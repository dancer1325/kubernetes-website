---
title: Controller
id: controller
full_link: /docs/concepts/architecture/controller/
short_description: >
  A control loop that watches the shared state of the cluster through the apiserver and makes changes attempting to move the current state towards the desired state.

aka: 
tags:
- architecture
- fundamental
---

* controller
  * == control loops / 
    * responsible for
      * watch -- , through the [kube-apiserver](kube-apiserver.md), -- the cluster's state
      * desired state == current state
  * == separate process / EACH controller
  * SOME run | [kube-controller-manager](kube-controller-manager.md)
    * _Examples:_ 
      * deployment controller
      * daemonset controller
      * namespace controller
      * persistent volume controller

* ALL controllers
  * are compiled | 1! binary
  * run | 1! process

- types of controllers
  - Node controller
    - if nodes go down → notice & respond
  - Job controller
    - watches [job](job.md) objects
  - EndpointSlice controller
    - populates [EndpointSlices](endpoint-slice.md)
  - ServiceAccount controller
    - create default [ServiceAccounts](service-account.md) / new namespaces
