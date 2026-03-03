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

- := network proxy /
  - runs | EACH cluster’s node (ALSO control plane)
  - maintains network rules | nodes
    - enable, | network sessions inside / outside the cluster, 
      - communicate -- with -- your pods
    - -> implement  [service](service.md)
  - how to be deployed | cluster?
    - -- as -- [container image](https://kubernetes.io/releases/download/#container-images)
  - if there is OS’ packet filtering layer -> uses it
    - packet filtering == examine data packets flowing -- through a -- network interface 
    - else -> forwards the traffic
  - ALLOWED modes / it can run
    - Iptables mode
      - == iptables rules
        -  👁️depends on node-ports & load-balancers 👁️
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
  * [source code](../command-line-tools-reference/kube-proxy)
