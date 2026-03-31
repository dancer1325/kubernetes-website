---
title: "Services, Load Balancing, and Networking"
weight: 60
description: >
  Concepts and resources behind networking in Kubernetes.
---

* goal
  * Kubernetes network model

* history
  * | older container systems,
    * ❌there was NO AUTOMATIC connectivity BETWEEN containers | DIFFERENT hosts❌
      * SOLUTION:
        * solution1: explicitly create links BETWEEN containers
        * solution2: map container ports -- to -- host ports
          * Reason:🧠make them reachable by containers | other hosts🧠
  * | Kubernetes
    * pods == VMs | physical hosts, about
      * port allocation
      * naming
      * service discovery
      * load balancing
      * application configuration
      * migration

- Kubernetes network model
  - characteristics
    - 💡[1! IP / EACH cluster's pod](#1-ip--each-clusters-pod)💡
    - [pod network OR cluster network](#pod-network-or-cluster-network)
    - [service proxying](#service-proxying)
    - [network policies](../../reference/glossary/network-policy.md)
    - [Gateway API](gateway)
    - [Service API](service.md)
    - [OWN private network namespace / EACH pod](#own-private-network-namespace--each-pod)
  - implementation
    - SOME are implemented by Kubernetes
    - ⚠️SOME is provided -- , through Kubernetes API, by -- other parts⚠️
  - networking implementation
    - AVAILABLE | DIFFERENT CR (Container Runtime)

## 1! IP / EACH cluster's pod

- ⚠️cluster-wide⚠️
  - != node-scope
- ❌️you do NOT need MANUALLY to❌
  - pod1 <- link -> pod2
    - == -- through -- IP
  - container’s ports <- map to -> host’s ports
  - Reason:🧠abstracted & managed by Kubernetes🧠
- 👀allocated -- by -- [network plugins](../extend-kubernetes/compute-storage-net/network-plugins.md)👀

```mermaid
flowchart TB
    subgraph CLUSTER["☁️ Kubernetes Cluster"]
        direction TB

        CNI["🔌 Network Plugin"]

        subgraph NODE1["🖥️ Node 1"]
            subgraph POD1["Pod A · IP: 10.0.1.1"]
                C1["container1"]
                C2["container2"]
                C1 <-->|"localhost"| C2
            end
            subgraph POD2["Pod B · IP: 10.0.1.2"]
                C3["container3"]
            end
        end

        subgraph NODE2["🖥️ Node 2"]
            subgraph POD3["Pod C · IP: 10.0.2.1"]
                C4["container4"]
            end
        end

        CNI -->|"assigns IP"| POD1
        CNI -->|"assigns IP"| POD2
        CNI -->|"assigns IP"| POD3

        POD1 <-->|"direct IP"| POD2
        POD1 <-->|"direct IP"| POD3
        POD2 <-->|"direct IP"| POD3
    end

    style CLUSTER fill:#e8f4fd,stroke:#2196F3,stroke-width:2px
    style CNI fill:#f3e5f5,stroke:#9C27B0,stroke-width:2px
    style NODE1 fill:#fff3e0,stroke:#FF9800
    style NODE2 fill:#fff3e0,stroke:#FF9800
    style POD1 fill:#e8f5e9,stroke:#4CAF50
    style POD2 fill:#e8f5e9,stroke:#4CAF50
    style POD3 fill:#e8f5e9,stroke:#4CAF50
```

## pod network OR cluster network

TODO: 
- Linux network namespace setup
  - system-level software implement the [CRI](../containers/cri.md)
- itself
  - managed -- by -- a [pod network implementation](../cluster-administration/addons.md#networking-and-network-policy)
    - | Linux, MOST container runtimes use  _CNI plugins_
- handle communication -- , WITHOUT using proxies OR NATs, -- BETWEEN pods /
  - node1's pod1 can communicate -- with -- node1's pod2
    - | Windows,
      - NOT valid | host-network pods
  - node1's pod1 can communicate -- with -- node2's pod2
    - | Windows,
      - NOT valid | host-network pods
- node’s agent can communicate -- with — ALL nod's pod 

```mermaid
flowchart TB
    subgraph CLUSTER["☁️ Pod Network / Cluster Network"]

        subgraph SETUP["🔧 Namespace Setup"]
            kubelet["kubelet"]
            cri["CRI\n(Container Runtime Interface)\nsystem-level software"]
            runtime["Container Runtime\n(containerd / CRI-O)"]
            ns["🐧 Linux Network Namespace\n(per pod)"]
            cni["CNI Plugin\n(pod network implementation)"]

            kubelet -->|"gRPC"| cri
            cri --> runtime
            runtime -->|"creates"| ns
            runtime -->|"calls"| cni
            cni -->|"assigns IP\nconfigures routes"| ns
        end

        subgraph NODE1["🖥️ Node 1"]
            subgraph POD1["Pod 1\nIP: 10.0.1.1"]
                C1["container"]
            end
            subgraph POD2["Pod 2\nIP: 10.0.1.2"]
                C2["container"]
            end
            AGENT1["🤖 node agent\n(kubelet)"]
        end

        subgraph NODE2["🖥️ Node 2"]
            subgraph POD3["Pod 3\nIP: 10.0.2.1"]
                C3["container"]
            end
            AGENT2["🤖 node agent\n(kubelet)"]
        end

        POD1 <-->|"✅ direct · no NAT · no proxy\n⚠️ Windows: NOT valid for host-network pods"| POD2
        POD1 <-->|"✅ direct · no NAT · no proxy\n⚠️ Windows: NOT valid for host-network pods"| POD3

        AGENT1 <-->|"✅ all pods on node"| POD1
        AGENT1 <-->|"✅ all pods on node"| POD2
        AGENT2 <-->|"✅ all pods on node"| POD3
    end

    style CLUSTER fill:#e8f4fd,stroke:#2196F3,stroke-width:2px
    style SETUP fill:#fff8e1,stroke:#FFC107
    style NODE1 fill:#fff3e0,stroke:#FF9800
    style NODE2 fill:#fff3e0,stroke:#FF9800
    style POD1 fill:#e8f5e9,stroke:#4CAF50
    style POD2 fill:#e8f5e9,stroke:#4CAF50
    style POD3 fill:#e8f5e9,stroke:#4CAF50
    style cni fill:#f3e5f5,stroke:#9C27B0
    style ns fill:#e3f2fd,stroke:#1565C0
```


## service proxying

- [kube-proxy](../../reference/glossary/kube-proxy.md)
  - default Kubernetes implementation
- SOME pod network implementations use their own service
  - Reason:🧠MORE integrated with the rest of the implementation🧠

## OWN private network namespace / EACH pod

- shared by ALL containers | pod
- pod1's container1 can communicate -- , through "localhost", with -- pod1's container2 