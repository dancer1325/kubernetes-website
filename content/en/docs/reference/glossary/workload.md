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
  * := application / running | pods

* workload resources
  * built-in
    * [DaemonSet](daemonset.md)
    * [Deployment](deployment.md)
    * [Job](job.md) & [CronJobs](cronjob.md)
    * [ReplicaSet](replica-set.md)
    * [StatefulSet](statefulset.md)
  * allows
    * managing on your behalf the pods
      * ⚠️OTHERWISE, you would need to MANUALLY manage EACH pod⚠️
  * configure a [controller](controller.md)
  * [ways to extend](../../concepts/extend-kubernetes)

* _Example:_ workload / has a web server + database
  * database could run | [StatefulSet](statefulset.md)
  * web server could run | [deployment](deployment.md)
