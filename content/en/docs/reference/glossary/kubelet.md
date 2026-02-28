---
title: Kubelet
id: kubelet
full_link: /docs/reference/command-line-tools-reference/kubelet
short_description: >
  An agent that runs on each node in the cluster. It makes sure that containers are running in a pod.

aka:
tags:
- fundamental
---

- := node agent / 
  - runs | 👁️EACH cluster's node 👁
  - ensures that containers & pods are running & healthy
    - EXCEPTION: ⚠️containers / NOT created -- by -- Kubernetes⚠️
  - acts as bridge between master node < - > rest of nodes
  - TODO: fetches individual container statistics from the [`Container Runtime`](Kubernetes%20Components%20cc7b751f553341049dd7c054085c18f7.md) — via — Container Runtime Interface
      - if you use a CR / uses Linux cgroups & namespaces to implement containers + CR does NOT publish usage statistics → kubelet can look up the statistics directly via [cAdvisor](../../Reference/Glossary/Tool%20d2e30e74acaf4ac094ae1ae0be053a2d.md)

It can register the node with the apiserver using one of: the hostname; a flag to
override the hostname; or specific logic for a cloud provider.

The kubelet works in terms of a PodSpec. A PodSpec is a YAML or JSON object
that describes a pod. The kubelet takes a set of PodSpecs that are provided through
various mechanisms (primarily through the apiserver) and ensures that the containers
described in those PodSpecs are running and healthy. The kubelet doesn't manage
containers which were not created by Kubernetes.

Other than from an PodSpec from the apiserver, there are two ways that a container
manifest can be provided to the Kubelet.

File: Path passed as a flag on the command line. Files under this path will be monitored
periodically for updates. The monitoring period is 20s by default and is configurable
via a flag.

HTTP endpoint: HTTP endpoint passed as a parameter on the command line. This endpoint
is checked every 20 seconds (also configurable with a flag)
