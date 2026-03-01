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
  * == application / 
    * 's core
      * flexible
      * extensible -- to -- support custom resources
    * serves -- , via REST interface, -- Kubernetes functionality
      * _Example:_ operations & communications BETWEEN 
        * Kubernetes components
        * Kubernetes components -- & -- external
    * stores the cluster's state
  * ways to use it
    * DIRECTLY
    * -- through -- tools (_Example:_ [kubectl](kubectl.md))
