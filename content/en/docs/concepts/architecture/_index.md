---
title: "Cluster Architecture"
weight: 30
description: >
  The architectural concepts behind Kubernetes.
---

* goal
  * Kubernetes cluster's components / cluster is
    * complete 
    * working 

* [Kubernetes cluster](.././../reference/glossary/cluster.md)
* [Kubernetes components](../overview/components.md)

![](/static/images/docs/kubernetes-cluster-architecture.svg)

## Control plane components

The control plane's components make global decisions about the cluster (for example, scheduling),
as well as detecting and responding to cluster events (for example, starting up a new
{{< glossary_tooltip text="pod" term_id="pod">}} when a Deployment's
`{{< glossary_tooltip text="replicas" term_id="replica" >}}` field is unsatisfied).

Control plane components can be run on any machine in the cluster
* However, for simplicity, setup scripts
typically start all control plane components on the same machine, and do not run user containers on this machine.
See [Creating Highly Available clusters with kubeadm](/docs/setup/production-environment/tools/kubeadm/high-availability/)
for an example control plane setup that runs across multiple machines.

### kube-apiserver

* [here](../../reference/glossary/kube-apiserver.md)

### etcd

* [here](../../reference/glossary/etcd.md)

### kube-scheduler

* [here](../../reference/glossary/kube-scheduler.md)

### kube-controller-manager

* [here](../../reference/glossary/kube-controller-manager.md)

### cloud-controller-manager

* [here](../../reference/glossary/cloud-controller-manager.md)

## Node components

* run | EVERY node

### kubelet

* [here](../../reference/glossary/kubelet.md)

### kube-proxy {#kube-proxy}

* [here](../../reference/glossary/kube-proxy.md)

### Container runtime

* [here](../../reference/glossary/container-runtime.md)
* [here](../../setup/production-environment/container%20runtimes)

## Addons

* [here](../../reference/glossary/addons.md)
* [full list of Addons](../../concepts/cluster-administration/addons.md)

### DNS

* [cluster DNS](../services-networking/dns-pod-service.md)
  * == DNS server /
    * | your environment, you could have MULTIPLE
    * serves DNS records -- for -- Kubernetes services
    * if containers started by Kubernetes -> included AUTOMATICALLY | DNS searches
  * required by ALMOST ALL Kubernetes clusters

### Web UI (Dashboard)

* [here](../../reference/tools/_index.md#dashboard)

### Container resource monitoring

* [Container Resource Monitoring](../../tasks/debug/debug-cluster/resource-usage-monitoring.md)

### Cluster-level Logging

* [cluster-level logging](../../concepts/cluster-administration/logging.md)

### Network plugins

* [here](../../concepts/extend-kubernetes/compute-storage-net/network-plugins)

## Architecture variations

* Kubernetes core components
  * remain consistent
  * way to be deployed & managed: can vary

### 👀ways to deploy the control plane components👀

* Traditional deployment
  * Control plane components run directly | dedicated machines OR VMs (often managed -- as -- systemd services)

* [Static Pods ](../../reference/glossary/static-pod.md)

* Self-hosted
  * control plane runs as Pods | Kubernetes cluster itself /
    * managed by: Deployments, StatefulSets or other Kubernetes primitives

* Managed Kubernetes services
  * _Example:_ cloud providers
  * abstract away the control plane,
    * == manage control plane components -- via -- some offered service

### how to place workloads +  control plane components?

* depend on
  * cluster size
  * performance requirements
  * operational policies

* recommendations
  - | smaller OR development clusters,
    - run control plane components & user workloads | SAME nodes
  - | larger production clusters,
    - control plane components | specific nodes
    - user workloads | other nodes 
* other approaches
  * critical add-ons or monitoring tools | control plane nodes

### Cluster management tools

* _Example:_ kubeadm, kops, and Kubespray
* DIFFERENT approach to deploy & manage clusters / EACH cluster management tool

### Customization and extensibility

* type of customizations
  - custom schedulers /
    - work alongside the default Kubernetes scheduler
    - replace the default Kubernetes scheduler
  - extend API servers -- with -- CustomResourceDefinitions & API Aggregation
  - integrate cloud providers -- , via cloud-controller-manager, with -- Kubernetes 

## {{% heading "whatsnext" %}}

Learn more about the following:

- [Nodes](/docs/concepts/architecture/nodes/) and
  [their communication](/docs/concepts/architecture/control-plane-node-communication/)
  with the control plane.
- Kubernetes [controllers](/docs/concepts/architecture/controller/).
- [kube-scheduler](/docs/concepts/scheduling-eviction/kube-scheduler/) which is the default scheduler for Kubernetes.
- Etcd's official [documentation](https://etcd.io/docs/).
- Several [container runtimes](/docs/setup/production-environment/container-runtimes/) in Kubernetes.
- Integrating with cloud providers using [cloud-controller-manager](/docs/concepts/architecture/cloud-controller/).
- [kubectl](/docs/reference/generated/kubectl/kubectl-commands) commands.
