---
title: Container Storage Interface (CSI)
id: csi
full_link: /docs/concepts/storage/volumes/#csi
short_description: >
    The Container Storage Interface (CSI) defines a standard interface to expose storage systems to containers.


aka: 
tags:
- storage 
---

* Container Storage Interface (CSI)
  * == standard interface / expose storage systems -- to -- containers
    * == container can use a persistence storage / ❌NOT know anything about it ❌
  * enable
    * vendors can create CUSTOM storage plugins / WITHOUT adding them | Kubernetes repository
  * CSI driver
    * [AVAILABLE ones](https://kubernetes-csi.github.io/docs/drivers.html) 
    * if you want to use a CSI driver | storage provider -> steps
      * [deploy the csi driver | your cluster](https://kubernetes-csi.github.io/docs/deploying.html)
      * create a [storage class](storage-class.md) / uses that CSI driver
  * [MORE](../../concepts/storage/volumes.md#csi)

```mermaid
flowchart LR
    subgraph KUBE["☁️ Kubernetes Cluster"]
        subgraph POD["🐳 Container"]
            APP["app\nreads/writes /data"]
            MOUNT["/data ← mounted volume\n(== local filesystem)"]
            APP --> MOUNT
        end

        subgraph CSI["💡 CSI Driver"]
            IFACE["standard interface\n❌ container does NOT know\nwhat's behind it"]
        end

        MOUNT -->|"I/O"| IFACE
    end

    subgraph VENDOR["🏢 Vendor"]
        PLUGIN["Custom Storage Plugin"]

        subgraph STORAGE["💾 actual Storage System"]
            S1["AWS EBS"]
            S2["GCP PD"]
            S3["NFS / Ceph..."]
        end

        PLUGIN --> S1
        PLUGIN --> S2
        PLUGIN --> S3
    end

    IFACE <--> PLUGIN

    style KUBE fill:#e8f4fd,stroke:#2196F3,stroke-width:2px
    style POD fill:#fff3e0,stroke:#FF9800
    style CSI fill:#ede7f6,stroke:#7B1FA2,stroke-width:2px
    style IFACE fill:#f3e5f5,stroke:#9C27B0
    style VENDOR fill:#e8f5e9,stroke:#388E3C,stroke-width:2px
    style PLUGIN fill:#c8e6c9,stroke:#2E7D32
    style STORAGE fill:#a5d6a7,stroke:#1B5E20
```
