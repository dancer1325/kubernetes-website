---
title: Master
id: master
short_description: >
  Legacy term, used as synonym for nodes running the control plane.

aka:
tags:
- fundamental
---

* master/master node
  * == ⚠️legacy term ⚠️
    * ALTHOUGH STILL used by some provisioning tools
      * _Examples:_ [kubeadm](kubeadm.md), managed services
      * -- to --
        * label nodes -- with -- `kubernetes.io/role`
        * control control plane pods location 
  * := node / hosts the control plane
