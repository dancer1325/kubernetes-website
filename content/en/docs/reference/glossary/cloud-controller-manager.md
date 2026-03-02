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

* OPTIONAL
* integrates -- with -- underlying cloud provider(s)


A Kubernetes {{< glossary_tooltip text="control plane" term_id="control-plane" >}} component
that embeds cloud-specific control logic. The cloud controller manager lets you link your
cluster into your cloud provider's API, and separates out the components that interact
with that cloud platform from components that only interact with your cluster.

<!--more-->

By decoupling the interoperability logic between Kubernetes and the underlying cloud
infrastructure, the cloud-controller-manager component enables cloud providers to release
features at a different pace compared to the main Kubernetes project.

The cloud-controller-manager only runs controllers that are specific to your cloud provider.
If you are running Kubernetes on your own premises, or in a learning environment inside your
own PC, the cluster does not have a cloud controller manager.

As with the kube-controller-manager, the cloud-controller-manager combines several logically
independent control loops into a single binary that you run as a single process. You can scale
horizontally (run more than one copy) to improve performance or to help tolerate failures.

The following controllers can have cloud provider dependencies:

- Node controller: For checking the cloud provider to determine if a node has been
  deleted in the cloud after it stops responding
- Route controller: For setting up routes in the underlying cloud infrastructure
- Service controller: For creating, updating and deleting cloud provider load balancers
