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

* [kube-apiserver](../architecture/_index.md#kube-apiserver)
* [etcd](../architecture/_index.md#etcd)
* [kube-scheduler](../architecture/_index.md#kube-scheduler)
* [kube-controller-manager](../architecture/_index.md#kube-controller-manager)
* [cloud-controller-manager](../architecture/_index.md#cloud-controller-manager)
  * OPTIONAL

### Node Components

* [kubelet](../architecture/_index.md#kubelet)
* [kube-proxy](../architecture/_index.md#kube-proxy-kube-proxy)
  * OPTIONAL
* [container runtime](../architecture/_index.md#container-runtime)

* | your cluster,
  * you may require ADDITIONAL software | EACH node
    * _Example:_ [systemd](https://systemd.io/)

## Addons

* [addons](../architecture/_index.md#addons)
  * _Examples:_
    * [DNS](../architecture/_index.md#dns)
    * [Web UI (Dashboard)](../architecture/_index.md#web-ui-dashboard)
    * [Container Resource Monitoring](../architecture/_index.md#container-resource-monitoring)
    * [Cluster-level Logging](../architecture/_index.md#cluster-level-logging)

## Flexibility | Architecture

* flexibility | deploy & manage these components
  * _Example:_ [small development environments, large-scale production deployments]

* [here](../../architecture/_index.md#architecture-variations)
