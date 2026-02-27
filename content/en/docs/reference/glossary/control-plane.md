---
title: Control Plane
id: control-plane
full_link:
short_description: >
  The container orchestration layer that exposes the API and interfaces to define, deploy, and manage the lifecycle of containers.

aka:
tags:
- fundamental
---

* control plane
  * == container orchestration layer /
    * exposes, about lifecycle of containers (define + deploy + manage), 
      * API 
      * interfaces
  * == MULTIPLE components
    * [etcd](etcd.md)
    * [API Server](kube-apiserver.md)
    * [Scheduler](kube-scheduler.md)
    * [Controller Manager](kube-controller-manager.md)
    * [Cloud Controller Manager](cloud-controller-manager.md)
    * ...
  * | production environments, 
    * NORMALLY runs ACROSS MULTIPLE computers

* ways / control plane's components can be run
  * daemons OR
  * containers

* [master](master.md)
  * := hosts / run control plane's components
