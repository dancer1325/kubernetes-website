---
title: Workload
id: workload
full_link: /docs/concepts/workloads/
short_description: >
   A workload is an application running on Kubernetes.

aka: 
tags:
- fundamental
---

* workload
  * := application / running | Kubernetes
  * 's parts
    * DaemonSet
    * Deployment
    * Job
    * ReplicaSet
    * StatefulSet

* _Example:_ workload / has a web server + database
  * database could run | [StatefulSet](statefulset.md)
  * web server could run | [deployment](deployment.md)
