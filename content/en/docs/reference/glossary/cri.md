---
title: Container Runtime Interface (CRI)
id: cri
full_link: /docs/concepts/architecture/cri
short_description: >
  Protocol for communication between the kubelet and the local container runtime.

aka:
tags:
  - fundamental
---

- [CRI/Container Runtime Interface](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md)
  - == 💡plugin interface 💡/
    - enable
      - [Kubelet](kubelet.md) can use -- , WITHOUT recompiling the cluster components, -- MULTIPLE CR
  - == 💡[gRPC](https://grpc.io) API + protocol buffers💡
  - requirements
    - [CR](container-runtime.md) | EACH cluster's node
      - Reason:🧠kubelet can launch pod + pod's containers🧠
  - allows
    - pulling the container image — from — a registry
    - unpacking the container
    - running the application
  - implementations / supported by Kubernetes
    - [CRI-O](https://cri-o.io/#what-is-cri-o)
    - [containerd](https://containerd.io/docs/)

```mermaid
flowchart TB
    subgraph CLUSTER["☁️ Kubernetes Cluster — CR required on EACH node"]
        subgraph N1["🖥️ Node 1"]
            K1["🤖 kubelet"]
            CR1["CR ✅"]
            K1 --- CR1
        end
        subgraph N2["🖥️ Node 2"]
            K2["🤖 kubelet"]
            CR2["CR ✅"]
            K2 --- CR2
        end
        subgraph N3["🖥️ Node 3"]
            K3["🤖 kubelet"]
            CR3["CR ✅"]
            K3 --- CR3
        end
    end

    subgraph SINGLE["1 Cluster's node"]
        kubelet["🤖 kubelet\n(client)"]

        subgraph CRI["💡 CRI — plugin interface + gRPC protocol"]
            GRPC["gRPC calls\n(RunPodSandbox, CreateContainer...)"]
            subgraph IMPLS["Implementations"]
                CRIO["CRI-O"]
                CONTAINERD["containerd"]
            end
        end

        subgraph CR_OPS["Container Runtime — allows"]
            PULL["🐳 pull image\nfrom registry"]
            UNPACK["📦 unpack\ncontainer"]
            RUN["🚀 run\napplication"]
            PULL --> UNPACK --> RUN
        end

        kubelet -->|"gRPC\n(WITHOUT recompiling cluster)"| GRPC
        GRPC --> CRIO
        GRPC --> CONTAINERD
        CRIO --> CR_OPS
        CONTAINERD --> CR_OPS
    end

    style CLUSTER fill:#e8f4fd,stroke:#2196F3,stroke-width:2px
    style N1 fill:#fff3e0,stroke:#FF9800
    style N2 fill:#fff3e0,stroke:#FF9800
    style N3 fill:#fff3e0,stroke:#FF9800
    style SINGLE fill:#f9f9f9,stroke:#9E9E9E,stroke-width:2px
    style CRI fill:#ede7f6,stroke:#7B1FA2,stroke-width:2px
    style IMPLS fill:#f3e5f5,stroke:#9C27B0
    style CR_OPS fill:#e8f5e9,stroke:#4CAF50
    style kubelet fill:#e3f2fd,stroke:#1565C0,stroke-width:2px
```
