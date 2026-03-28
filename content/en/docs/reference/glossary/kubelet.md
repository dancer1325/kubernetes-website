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
  - responsible for
    - ensures that: containers & pods are running & healthy
      - -- based on the -- PodSpec / specified -- via --
        - API server
        - File path
          - `kubelet ... --pod-manifest-path filePath`
        - HTTP endpoint
          - `kubelet ... --manifest-url HTTPEndpoint`
      - EXCEPTION: ⚠️containers / NOT created -- by -- Kubernetes⚠️
    - fetches individual container statistics — , via CRI, from the -- [`Container Runtime`](container-runtime.md)
      - if you use a CR / implement containers -- via -- Linux cgroups & namespaces + CR does NOT publish usage statistics → kubelet look up the statistics -- via -- [cAdvisor](cadvisor.md)
    - registering it's OWN node | apiserver, -- through --
      - hostname
      - flag / override the hostname
      - specific logic -- for a -- cloud provider
    - acts as bridge BETWEEN [master node (== control plane)](control-plane.md) -- & -- rest of nodes

* PodSpec
  * == YAML OR JSON object /
    * describes a pod

```mermaid
flowchart TB
    subgraph CLUSTER["☁️ Kubernetes Cluster"]
        subgraph CPNODE["🖥️ Control Plane Node\n(runs API Server, etcd, scheduler...)"]
            APISERVER["🗂️ API Server"]
            KCP["🤖 kubelet"]
        end

        subgraph NODE1["🖥️ Worker Node 1"]
            K1["🤖 kubelet"]
        end

        subgraph NODE2["🖥️ Worker Node 2"]
            K2["🤖 kubelet"]
        end
    end

    APISERVER <-->|"🌉 bridge"| K1
    APISERVER <-->|"🌉 bridge"| K2
    KCP <-->|"🌉 bridge\n(kubelet also runs here)"| APISERVER

    subgraph DETAIL[" 1 Cluster's node"]
        direction TB

        subgraph PODSPEC_SOURCES["📋 PodSpec Sources"]
            SRC_API["API Server (primary)"]
            SRC_FILE["--pod-manifest-path (file)"]
            SRC_HTTP["--manifest-url (HTTP, poll 20s)"]
        end

        kubelet["🤖 kubelet"]

        subgraph PODS["Pods & Containers"]
            POD1["Pod ✅ running & healthy"]
            SKIP["⚠️ container NOT created\nby Kubernetes → IGNORED"]
        end

        CR["CR (Container Runtime)"]

        subgraph STATS["📊 Container Stats"]
            CADVISOR["cAdvisor"]
            CGROUPS["Linux cgroups\n& namespaces"]
            CADVISOR -->|"reads"| CGROUPS
        end

        SRC_API --> kubelet
        SRC_FILE --> kubelet
        SRC_HTTP --> kubelet
        kubelet -->|"ensures"| POD1
        kubelet -.->|"ignores"| SKIP
        kubelet -->|"1️⃣ fetch stats via CRI"| CR
        CR -->|"✅ publishes stats"| kubelet
        CR -.->|"❌ does NOT publish stats\n(uses cgroups & namespaces)"| CADVISOR
        CADVISOR -->|"2️⃣ fallback: kubelet\nlooks up stats here"| kubelet

        subgraph REGISTER["📝 register OWN node | apiserver — through"]
            direction LR
            HN["hostname\n(default)"]
            FLAG["--hostname-override\n(flag)"]
            CLOUD["specific logic\n(cloud provider)"]
        end

        DETAIL_API["🗂️ API Server"]

        kubelet -->|"self-registers"| REGISTER
        REGISTER -->|"POST Node object"| DETAIL_API
    end

    style CLUSTER fill:#e8f4fd,stroke:#2196F3,stroke-width:2px
    style CPNODE fill:#fce4ec,stroke:#E53935,stroke-width:2px
    style NODE1 fill:#fff3e0,stroke:#FF9800
    style NODE2 fill:#fff3e0,stroke:#FF9800
    style DETAIL fill:#f9f9f9,stroke:#9E9E9E
    style PODSPEC_SOURCES fill:#fff8e1,stroke:#FFC107
    style PODS fill:#e8f5e9,stroke:#4CAF50
    style REGISTER fill:#e8f5e9,stroke:#388E3C
    style kubelet fill:#e3f2fd,stroke:#1565C0,stroke-width:2px
    style SKIP fill:#fce4ec,stroke:#F44336,stroke-dasharray: 5 5
    style CADVISOR fill:#f3e5f5,stroke:#9C27B0
    style CRI fill:#ede7f6,stroke:#7B1FA2
```
