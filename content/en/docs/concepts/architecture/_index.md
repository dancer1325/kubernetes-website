---
title: "Cluster Architecture"
weight: 30
description: >
  The architectural concepts behind Kubernetes.
---

* [Kubernetes cluster](.././../reference/glossary/cluster.md)
* [Kubernetes components](../overview/Kubernetes%20Components/components.md)

The control plane manages the worker nodes and the Pods in the cluster
In production
environments, the control plane usually runs across multiple computers and a cluster
usually runs multiple nodes, providing fault-tolerance and high availability.

This document outlines the various components you need to have for a complete and working Kubernetes cluster.

{{< figure src="/images/docs/kubernetes-cluster-architecture.svg" alt="The control plane (kube-apiserver, etcd, kube-controller-manager, kube-scheduler) and several nodes. Each node is running a kubelet and kube-proxy." caption="Figure 1. Kubernetes cluster components." class="diagram-large" >}}

{{< details summary="About this architecture" >}}
The diagram in Figure 1 presents an example reference architecture for a Kubernetes cluster.
The actual distribution of components can vary based on specific cluster setups and requirements.

In the diagram, each node runs the [`kube-proxy`](#kube-proxy) component. You need a
network proxy component on each node to ensure that the
{{< glossary_tooltip text="Service" term_id="service">}} API and associated behaviors
are available on your cluster network. However, some network plugins provide their own,
third party implementation of proxying. When you use that kind of network plugin,
the node does not need to run `kube-proxy`.
{{< /details >}}

## Control plane components

The control plane's components make global decisions about the cluster (for example, scheduling),
as well as detecting and responding to cluster events (for example, starting up a new
{{< glossary_tooltip text="pod" term_id="pod">}} when a Deployment's
`{{< glossary_tooltip text="replicas" term_id="replica" >}}` field is unsatisfied).

Control plane components can be run on any machine in the cluster. However, for simplicity, setup scripts
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

TODO:
If you use a [network plugin](#network-plugins) that implements packet forwarding for Services
by itself, and providing equivalent behavior to kube-proxy, then you do not need to run
kube-proxy on the nodes in your cluster.

### Container runtime

* [here](../../reference/glossary/container-runtime.md)
* [here](../../setup/production-environment/container%20runtimes)

## Addons

TODO:
* Addons use Kubernetes resources ({{< glossary_tooltip term_id="daemonset" >}},
{{< glossary_tooltip term_id="deployment" >}}, etc) to implement cluster features.
Because these are providing cluster-level features, namespaced resources for
addons belong within the `kube-system` namespace.

Selected addons are described below; for an extended list of available addons,
please see [Addons](/docs/concepts/cluster-administration/addons/).

implement cluster-level features — via —  [Kubernetes resources](../Workloads/Workload%20Resources%20bedee4e8f78746bb9cd3dd82599906f2.md)
- → belong to `kube-system` namespace

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

TODO: A [cluster-level logging](/docs/concepts/cluster-administration/logging/) mechanism is responsible
for saving container logs to a central log store with a search/browsing interface.
For saving container logs to a central log store.

### Network plugins

[Network plugins](../../concepts/extend-kubernetes/compute-storage-net/network-plugins)
are software components that implement the container network interface (CNI) specification.
They are responsible for allocating IP addresses to pods and enabling them to communicate
with each other within the cluster.

## Architecture variations

While the core components of Kubernetes remain consistent, the way they are deployed and
managed can vary. Understanding these variations is crucial for designing and maintaining
Kubernetes clusters that meet specific operational needs.

### 👀ways to deploy the control plane components👀

* Traditional deployment
  * Control plane components run directly | dedicated machines OR VMs (often managed -- as -- systemd services)

* Static Pods
  * Control plane components are deployed -- as -- static Pods / 
    * managed by the kubelet | specific nodes
  * uses
    * by tools
      * _Example:_ kubeadm

* Self-hosted
: The control plane runs as Pods within the Kubernetes cluster itself, managed by Deployments
  and StatefulSets or other Kubernetes primitives.

* Managed Kubernetes services
: Cloud providers often abstract away the control plane, managing its components as part of their service offering.

### Workload placement considerations

The placement of workloads, including the control plane components, can vary based on cluster size,
performance requirements, and operational policies:

- In smaller or development clusters, control plane components and user workloads might run on the same nodes.
- Larger production clusters often dedicate specific nodes to control plane components,
  separating them from user workloads.
- Some organizations run critical add-ons or monitoring tools on control plane nodes.

### Cluster management tools

Tools like kubeadm, kops, and Kubespray offer different approaches to deploying and managing clusters,
each with its own method of component layout and management.

### Customization and extensibility

Kubernetes architecture allows for significant customization:

- Custom schedulers can be deployed to work alongside the default Kubernetes scheduler or to replace it entirely.
- API servers can be extended with CustomResourceDefinitions and API Aggregation.
- Cloud providers can integrate deeply with Kubernetes using the cloud-controller-manager.

The flexibility of Kubernetes architecture allows organizations to tailor their clusters to specific needs,
balancing factors such as operational complexity, performance, and management overhead.

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
