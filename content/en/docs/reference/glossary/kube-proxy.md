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
  * [overview](#overview)
  * [packet filtering layer](#packet-filtering)
  * [ALLOWED modes / it can run](#allowed-modes--it-can-run)
  * [source code](../command-line-tools-reference/kube-proxy)

## overview

* kube-proxy
  - runs | EACH cluster’s node (ALSO control plane)
    - == default one == built-in by Kubernetes
    - Reason:🧠ensure service API & associated behaviors are AVAILABLE | your cluster network🧠
    - EXCEPTION: ⚠️if you use network plugins / provide their own network proxy -> it can be replaced⚠️
  - maintains (write, read) network rules | nodes
    - enable, | network sessions inside / outside the cluster,
      - communicate -- with -- your pods
    - -> implement  [service](service.md)
  - how to be deployed | cluster?
    - -- as -- [container image](https://kubernetes.io/releases/download/#container-images)

* network sessions inside / outside the cluster
  * inside
    * cluster1's pod1 -- cluster1's service -- cluster1's pod1 
  * outside
    * external client -- cluster1's service (type = NodePort / LoadBalancer) -- cluster1's pod

```mermaid
flowchart TB
    NETPLUGIN["⚠️ Network Plugin\n(provides own proxy)\n→ kube-proxy NOT needed"]

    subgraph CLUSTER["☁️ Kubernetes Cluster"]
        SVC["📡 Service API\n(config only — no traffic)"]

        subgraph CPNODE["🖥️ Control Plane Node"]
            KP_CP["🔀 kube-proxy"]
            RULES_CP["📋 network rules\n(kernel)"]
            KP_CP -->|"maintains"| RULES_CP
        end
        subgraph NODE1["🖥️ Worker Node 1"]
            KP1["🔀 kube-proxy"]
            RULES1["📋 network rules\n(kernel)"]
            KP1 -->|"maintains"| RULES1
        end
        subgraph NODE2["🖥️ Worker Node 2"]
            KP2["🔀 kube-proxy"]
            RULES2["📋 network rules\n(kernel)"]
            KP2 -->|"maintains"| RULES2
        end

        SVC -.->|"watch (config)"| KP_CP & KP1 & KP2
    end

    NETPLUGIN -.->|"EXCEPTION:\nreplaces"| KP_CP & KP1 & KP2

    style CLUSTER fill:#e8f4fd,stroke:#2196F3,stroke-width:2px
    style CPNODE fill:#fce4ec,stroke:#E53935
    style NODE1 fill:#fff3e0,stroke:#FF9800
    style NODE2 fill:#fff3e0,stroke:#FF9800
    style KP_CP fill:#ede7f6,stroke:#7B1FA2,stroke-width:2px
    style KP1 fill:#ede7f6,stroke:#7B1FA2,stroke-width:2px
    style KP2 fill:#ede7f6,stroke:#7B1FA2,stroke-width:2px
    style RULES_CP fill:#fff8e1,stroke:#FFC107
    style RULES1 fill:#fff8e1,stroke:#FFC107
    style RULES2 fill:#fff8e1,stroke:#FFC107
    style NETPLUGIN fill:#fce4ec,stroke:#F44336,stroke-dasharray: 5 5
    style SVC fill:#e3f2fd,stroke:#1565C0
```

## packet filtering

TODO: 
- if there is OS’ packet filtering layer -> uses it
  - packet filtering == examine data packets flowing -- through a -- network interface
  - else -> forwards the traffic

```mermaid
flowchart TB
    PACKET["📦 data packet\n(inside / outside cluster)"]

    subgraph KUBE_PROXY["🔀 kube-proxy"]
        CHECK{"OS packet\nfiltering layer\navailable?"}
        FILTER["use OS packet\nfiltering layer\n(examine packets\nthrough network interface)"]
        FORWARD["forward\nthe traffic\ndirectly"]
    end

    PODS["🐳 Pods"]

    PACKET --> CHECK
    CHECK -->|"✅ yes"| FILTER
    CHECK -->|"❌ no"| FORWARD
    FILTER --> PODS
    FORWARD --> PODS

    style KUBE_PROXY fill:#ede7f6,stroke:#7B1FA2,stroke-width:2px
    style CHECK fill:#fff8e1,stroke:#FFC107
    style FILTER fill:#e8f5e9,stroke:#4CAF50
    style FORWARD fill:#e3f2fd,stroke:#1565C0
    style PODS fill:#e8f5e9,stroke:#4CAF50
```

## ALLOWED modes / it can run

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

```mermaid
flowchart TB
  subgraph IPTABLES["🔧 iptables mode"]
      direction TB
      subgraph SVC_PORT1["/ service's port"]
          R1["1 rule in KUBE-SERVICES"]
          R2["1 rule in KUBE-SVC-HASH"]
      end
      subgraph EP1["/ pod's endpoint"]
          R3["rules in KUBE-SVC-HASH"]
          R4["rules in KUBE-SEP-HASH"]
      end
      NOTE1["👁️ depends on\nnode-ports & load-balancers"]
  end

  subgraph IPVS["⚡ ipvs mode"]
      direction TB
      subgraph VS["virtual servers (created by kube-proxy)"]
          VS1["service's port"]
          VS2["NodePorts"]
          VS3["external IPs"]
          VS4["load-balancer IPs"]
      end
      subgraph RS["real servers (created by kube-proxy)"]
          RS1["pod's endpoint"]
      end
      VS --> RS
  end

  TRAFFIC["🌐 traffic"] --> IPTABLES
  TRAFFIC --> IPVS

  style IPTABLES fill:#fff3e0,stroke:#FF9800,stroke-width:2px
  style IPVS fill:#e8f4fd,stroke:#2196F3,stroke-width:2px
  style SVC_PORT1 fill:#fff8e1,stroke:#FFC107
  style EP1 fill:#fce4ec,stroke:#E53935
  style VS fill:#e3f2fd,stroke:#1565C0
  style RS fill:#e8f5e9,stroke:#4CAF50
  style NOTE1 fill:#f3e5f5,stroke:#9C27B0,stroke-dasharray: 5 5
```
