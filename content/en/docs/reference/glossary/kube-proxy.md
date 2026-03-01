---
title: kube-proxy
id: kube-proxy
full_link: /docs/reference/command-line-tools-reference/kube-proxy/
short_description: >
  `kube-proxy` is a network proxy that runs on each node in the cluster.

aka:
tags:
- fundamental
- networking
---

[kube-proxy](/docs/reference/command-line-tools-reference/kube-proxy/)
These network rules allow network
communication to your Pods from network sessions inside or outside of
your cluster.

kube-proxy uses the operating system packet filtering layer if there is one
and it's available. Otherwise, kube-proxy forwards the traffic itself.

- := network proxy /
  - runs / EACH cluster’s node (ALSO control plane)
  - implements part of the [service](service.md) concept
  - maintains network rules | nodes
    - **Note:** 👀 allow from network sessions inside / outside the cluster — communicate to → your Pods 👀
- if there is OS’ packet filtering layer → uses it, else → forwards the traffic
**Note:** 👁️ packet filtering == examine data packets flowing through a network interface👁️
- 👁️ can be deployed as [CI](https://kubernetes.io/releases/download/#container-images) within the cluster 👁️
- modes in which it can run
  - Iptables mode
    - iptables rules
      - **Note:** 👁️ depends on node-ports & load-balancers 👁️
      - / service’s port
        - 1 rule in `KUBE-SERVICES`
        - 1 rule in `KUBE-SVC-HASH`
      - / pod’s endpoint
        - small number rules in `KUBE-SVC-HASH`
        - small number rules in `KUBE-SEP-HASH`
  - ipvs mode
    - virtual servers are created by `kube-proxy` /
      - service’s port
      - NodePorts
      - external IPs
      - load-balancer IPs
    - real servers are created by `kube-proxy` /
      - pod’s endpoint
- `kube-proxy` must be proxying
