---
title: "Scheduling, Preemption and Eviction"
weight: 95
content_type: concept
no_list: true
---

## Scheduling

* Scheduling
  * := process of guarantee that [Pods are matched -- with -- Nodes](assign-pod-node)
    * -> kubelet can run the pods
    * == serie of stages

* [Kubernetes Scheduler](kube-scheduler)
* [Pod Overhead](pod-overhead)
* [Pod Topology Spread Constraints](topology-spread-constraints)
* [Taints and Tolerations](taint-and-toleration)
* [Scheduling Framework](scheduling-framework)
* [Dynamic Resource Allocation](dynamic-resource-allocation)
* [Scheduler Performance Tuning](scheduler-perf-tuning)
* [Resource Bin Packing for Extended Resources](resource-bin-packing)
* [Pod Scheduling Readiness](pod-scheduling-readiness)
* [Gang Scheduling](gang-scheduling)
* [Descheduler](https://github.com/kubernetes-sigs/descheduler#descheduler-for-kubernetes)
* [Node Declared Features](node-declared-features)

## Pod Disruption

- [preemption](../../reference/glossary/preemption.md)
- [eviction](../../reference/glossary/eviction.md)
- [Pod Disruption](../../reference/glossary/pod-disruption.md)

* [Pod Priority and Preemption](pod-priority-preemption)
* [Node-pressure Eviction](node-pressure-eviction)
* [API-initiated Eviction](api-eviction)
