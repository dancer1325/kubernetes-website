---
title: kube-scheduler
id: kube-scheduler
full_link: /docs/reference/command-line-tools-reference/kube-scheduler/
short_description: >
  Control plane component that watches for newly created pods with no assigned node, and selects a node for them to run on.

aka: 
tags:
- architecture
---

* == Control plane's component /
  * responsible for
    * watch for NEWLY created pods / NO assigned node
    * selects a node | run the pods -- based on --
      * individual & collective [resource](infrastructure-resource.md) requirements
      * hardware/software/policy constraints
      * affinity & anti-affinity specifications
      * data locality
      * inter-workload interference
      * deadlines
