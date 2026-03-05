---
title: Static Pod
id: static-pod
full_link: /docs/tasks/configure-pod-container/static-pod/
short_description: >
  A pod managed directly by the kubelet daemon on a specific node.

aka: 
tags:
- fundamental
---

* static pod
  * == pod / ⚠️managed directly -- by the -- [kubelet](kubelet.md) daemon | specific node⚠️  
    * -> ⚠️ALWAYS bound -- to -- 1 Kubelet | specific node⚠️
    * WITHOUT being observed -- by the -- API server
    * if the kubelet finds a static pod | its configuration -> AUTOMATICALLY tries to create a Pod object | Kubernetes API server
      * -> visible | API server
        * ❌BUT NOT controller | API server❌
  * ❌NOT support [ephemeral containers](ephemeral-container.md)❌
  * vs non-static pods
    * pods are managed -- by the -- [control plane](control-plane.md)
  * uses
    * by tools
      * _Example:_ kubeadm
  * use cases
    * deploy Control plane components 
