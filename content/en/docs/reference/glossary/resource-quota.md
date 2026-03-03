---
title: ResourceQuota
id: resource-quota
full_link: /docs/concepts/policy/resource-quotas/
short_description: >
  Provides constraints that limit aggregate resource consumption per namespace.

aka: 
tags:
- fundamental
- operation
- architecture
---
Object that constrains aggregate resource
consumption, per {{< glossary_tooltip term_id="namespace" >}}.

<!--more-->

A ResourceQuota can either limit the quantity of {{< glossary_tooltip text="API resources" term_id="api-resource" >}}
that can be created in a namespace by type, or it can set a limit on the total amount of
{{< glossary_tooltip text="infrastructure resources" term_id="infrastructure-resource" >}}
that may be consumed on behalf of the namespace (and the objects within it).

- [`limits.cpu`](../../Concepts/Configuration/Resource%20Management%20for%20Pods%20and%20Containers%20dec00252d8ee41369f43063968c9d90d.md)
   - enforces container’s cpu < `limits.cpu`
   - / [scheduling interval](../../Concepts/Scheduling,%20Preemption%20and%20Eviction%2011a13fe1aaf54bcfa09ecf21179846c1.md)
     - linux kernel checks if `limits.cpu` is exceed
       - if it’s exceeded → cgroup resume [eviction](../../Concepts/Scheduling,%20Preemption%20and%20Eviction%2011a13fe1aaf54bcfa09ecf21179846c1.md)
         **Note:** linux kernel waits before it
 - [`limits.memory`](../../Concepts/Configuration/Resource%20Management%20for%20Pods%20and%20Containers%20dec00252d8ee41369f43063968c9d90d.md)
   - defined / cgroup
   - — can apply also to — memory backed volumes’ pages
     *Example:* emptyDir
 - [`requests.cpu`](../../Concepts/Configuration/Resource%20Management%20for%20Pods%20and%20Containers%20dec00252d8ee41369f43063968c9d90d.md)
   - if several cgroups (→ containers) want to run →
     CPU time allocated for [workloads](../../Concepts/Workloads%207f133c71ced8439a97de8ca6b4c88026.md) with large `requests.cpu` > CPU time allocated for [workloads](../../Concepts/Workloads%207f133c71ced8439a97de8ca6b4c88026.md) with small `requests.cpu`
     **Note:** 👁️ CPU time := time spent for a CPU, executing instructions / specific process or container 👁️
 - [`requests.memory`](../../Concepts/Configuration/Resource%20Management%20for%20Pods%20and%20Containers%20dec00252d8ee41369f43063968c9d90d.md)
   - if node with cgroup v2 → CR use it as hint to set
     **Note:** Check [article](https://kubernetes.io/blog/2023/05/05/qos-memory-resources/)
     - [`memory.min`](../../Concepts/Workloads/Pods/Pod%20Quality%20of%20Service%20Classes%20(QoS)%20662174e7cdc74aba91bfb18ad0ae38ad.md)
     - `memory.low`