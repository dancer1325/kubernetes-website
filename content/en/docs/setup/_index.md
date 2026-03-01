---
reviewers:
- brendandburns
- erictune
- mikedanese
title: Getting started
main_menu: true
weight: 20
content_type: concept
no_list: true
card:
  name: setup
  weight: 20
  anchors:
  - anchor: "#learning-environment"
    title: Learning environment
  - anchor: "#production-environment"
    title: Production environment
---

* goal
  * ways to set up & run Kubernetes

* Kubernetes type installation
  * choose -- based on --
    * ease of maintenance
    * security
    * control
    * available resources
    * expertise -- to -- operate & manage a cluster
  * manual
  * managed service
    * [certified platforms](production-environment/turnkey-solutions.md)

* some [Kubernetes components](/content/en/docs/concepts/overview/components/) can be deployed -- as -- container images | cluster
  * _Example:_ "kube-apiserver", "kube-proxy" 
  * NOT valid: kubelet
    * Reason: 🧠it's a node agent🧠

TODO: 
It is **recommended** to run Kubernetes components as container images wherever
that is possible, and to have Kubernetes manage those components.
Components that run containers - notably, the kubelet - can't be included in this category.

## Production environment

When evaluating a solution for a
[production environment](/docs/setup/production-environment/), consider which aspects of
operating a Kubernetes cluster (or _abstractions_) you want to manage yourself and which you
prefer to hand off to a provider.

For a cluster you're managing yourself, the officially supported tool
for deploying Kubernetes is [kubeadm](/docs/setup/production-environment/tools/kubeadm/).

## {{% heading "whatsnext" %}}

- [Download Kubernetes](/releases/download/)
- Download and [install tools](/docs/tasks/tools/) including `kubectl`
- Select a [container runtime](/docs/setup/production-environment/container-runtimes/) for your new cluster
- Learn about [best practices](/docs/setup/best-practices/) for cluster setup

- Learn to [set up clusters with Windows nodes](/docs/concepts/windows/)
