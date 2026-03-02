---
reviewers:
- lavalamp
title: Kubernetes Components
content_type: concept
description: >
  An overview of the key components that make up a Kubernetes cluster.
weight: 10
card:
  title: Components of a cluster
  name: concepts
  weight: 20
---

* goal
  * Kubernetes cluster' essential components 

![](/static/images/docs/components-of-kubernetes.svg)

## Kubernetes core components

* 👀control plane OR master node + \>=1 worker nodes👀

### Control Plane Components

* [kube-apiserver](../../architecture/_index.md#kube-apiserver)

* [etcd](../../architecture/_index.md#etcd)

* [kube-scheduler](../../architecture/_index.md#kube-scheduler)

* [kube-controller-manager](../../architecture/_index.md#kube-controller-manager)

* [cloud-controller-manager](../../architecture/_index.md#cloud-controller-manager)

### Node Components

Run on every node, maintaining running pods and providing the Kubernetes runtime environment:

[kubelet](/docs/concepts/architecture/#kubelet)

[kube-proxy](/docs/concepts/architecture/#kube-proxy) (optional)
: Maintains network rules on nodes to implement {{< glossary_tooltip text="Services" term_id="service" >}}.

[Container runtime](/docs/concepts/architecture/#container-runtime)
: Software responsible for running containers. Read
  [Container Runtimes](/docs/setup/production-environment/container-runtimes/) to learn more.

{{% thirdparty-content single="true" %}}

Your cluster may require additional software on each node; for example, you might also
run [systemd](https://systemd.io/) on a Linux node to supervise local components.

## Addons

Addons extend the functionality of Kubernetes. A few important examples include:

[DNS](/docs/concepts/architecture/#dns)
: For cluster-wide DNS resolution.

[Web UI](/docs/concepts/architecture/#web-ui-dashboard) (Dashboard)
: For cluster management via a web interface.

[Container Resource Monitoring](/docs/concepts/architecture/#container-resource-monitoring)
: For collecting and storing container metrics.

[Cluster-level Logging](/docs/concepts/architecture/#cluster-level-logging)
: For saving container logs to a central log store.

## Flexibility in Architecture

Kubernetes allows for flexibility in how these components are deployed and managed.
The architecture can be adapted to various needs, from small development environments
to large-scale production deployments.

For more detailed information about each component and various ways to configure your
cluster architecture, see the [Cluster Architecture](/docs/concepts/architecture/) page.
