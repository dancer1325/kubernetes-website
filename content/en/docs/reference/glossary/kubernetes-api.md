---
title: Kubernetes API
id: kubernetes-api
full_link: /docs/concepts/overview/kubernetes-api/
short_description: >
  The application that serves Kubernetes functionality through a RESTful interface and stores the state of the cluster.

aka: 
tags:
- fundamental
- architecture
---

* Kubernetes API
  * == 💡HTTP API💡 / 
    * 's core
      * flexible
      * extensible -- to -- support custom resources
    * serves -- , via REST interface, -- Kubernetes functionality
      * _Example:_ operations & communications BETWEEN 
        * Kubernetes components
        * Kubernetes components -- & -- external
    * stores the cluster's state
    * 👀implemented -- by the -- [Kubernetes API server](kube-apiserver.md)👀
    * 💡MOSTLY declarative💡
      * Reason: 🧠controller keeps desired state == cluster state🧠
      * MOSTLY
        * Reason: 🧠some are imperative
          * _Examples:_ `kubectl logs`, `kubectl exec`, ...🧠
  * ways to use it
    * DIRECTLY
      * == REST calls
    * -- through -- 
      * CL tools (_Example:_ [kubectl](kubectl.md), [kubeadm](kubeadm.md))
      * [client libraries](../using-api/client-libraries.md)
