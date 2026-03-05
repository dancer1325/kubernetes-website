---
title: "Workloads"
weight: 55
description: >
  Understand Pods, the smallest deployable compute object in Kubernetes, and the higher-level abstractions that help you to run them.
no_list: true
card:
  title: Workloads and Pods
  name: concepts
  weight: 60
---

* [definition](../../reference/glossary/workload.md)

## Workload placement

* [Workload API](workload-api/)

## {{% heading "whatsnext" %}}

TODO: There are two supporting concepts that provide backgrounds about how Kubernetes manages pods
for applications:
* [Garbage collection](/docs/concepts/architecture/garbage-collection/) tidies up objects
  from your cluster after their _owning resource_ has been removed.
* The [_time-to-live after finished_ controller](/docs/concepts/workloads/controllers/ttlafterfinished/)
  removes Jobs once a defined time has passed since they completed.
