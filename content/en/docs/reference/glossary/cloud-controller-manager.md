---
title: Cloud Controller Manager
id: cloud-controller-manager
full_link: /docs/concepts/architecture/cloud-controller/
short_description: >
  Control plane component that integrates Kubernetes with third-party cloud providers.
aka: 
tags:
- architecture
- operation
---

* == [Control plane](control-plane.md)'s component /
  * embeds cloud-specific control logic /
    * 
  * OPTIONAL
    * Reason:🧠ONLY run controllers / specific -- to -- your cloud provider🧠
      * -> ❌NO exist | 
        * premises Kubernetes cluster
        * learning environment | your own PC❌
  * allows
    * linking your cluster -- into -- your cloud provider's API /
      * components / interact with the cloud platform are separated -- from -- components / ONLY interact with your cluster
        * -> rhythm to release features by cloud providers can be != rhythm to release features by Kubernetes project 
  * if you want to improve performance -> scale horizontally
