---
title: Preemption
id: preemption
full_link: /docs/concepts/scheduling-eviction/pod-priority-preemption/#preemption
short_description: >
  Preemption logic in Kubernetes helps a pending Pod to find a suitable Node by evicting low priority Pods existing on that Node.

aka:
tags:
- operation
---

* Preemption
  * := process of terminating Pods / have lower [priority](pod-priority.md)
    * == if a high priority pod needs to run -> can -- , by evicting pods / lower priority, -- run | that node
  * [MORE](../../concepts/scheduling-eviction/pod-priority-preemption.md#preemption)
