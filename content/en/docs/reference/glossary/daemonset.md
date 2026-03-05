---
title: DaemonSet
id: daemonset
full_link: /docs/concepts/workloads/controllers/daemonset
short_description: >
  Ensures a copy of a Pod is running across a set of nodes in a cluster.

aka: 
tags:
- fundamental
- core-object
- workload
---

* ensures
  * a copy of a pod is running ACROSS a set of cluster's nodes
    * -> if you add a node & fits Daemonset's specification -> control plane schedules NEW pod | this node
* uses
  * deploy system daemons / normally run | every node
    * _Example:_ log collectors, monitoring agents
* use cases
  * fundamental operations | your cluster
    * _Example:_ [cluster networking](../../concepts/cluster-administration/networking.md#how-to-implement-the-kubernetes-network-model)
  * node management
    * _Example:_ Fluentd
  * optional behavior / enhances the container platform
    * _Example:_ GPU plugins
* job / perform a Daemonset's pod == job / perform a system daemon | classic Unix / POSIX server
