---
title: EndpointSlice
id: endpoint-slice
full_link: /docs/concepts/services-networking/endpoint-slices/
short_description: >
  EndpointSlices track the IP addresses of Pods for Services.

aka:
tags:
- networking
---

* EndpointSlice
  * := [Kubernetes Object](object.md) /
    * == 💡service's slice/subset of backend network endpoints💡 
    * != service's sub group
  * associated (NORMALLY) -- with a -- [service](service.md)
  * enable
    * services
      * are linked -- to -- Pods
      * can scale -- to -- handle LARGE number of backends
    * cluster can update its list of healthy backends
  * ways to be tracked/updated
    * AUTOMATICALLY -- by -- the control plane
    * MANUALLY -- for -- [Services](service.md) / NO [selectors](selector.md) specified

* backend network endpoints
  * == `IP:port` / receives the traffic
    * (NORMALLY) [pods](pod.md)
    * ❌ != service's ClusterIP ❌

```mermaid
flowchart LR
    KUBECTL["🛠️ kubectl / manual\n(Services / NO selectors)"]

    subgraph CLUSTER["☁️ Kubernetes Cluster"]
        CP["⚙️ control plane\n(AUTOMATICALLY)"]

        SVC["📡 Service (frontend)\nClusterIP: 10.96.0.5"]

        subgraph ES["EndpointSlice"]
            EP1["10.0.1.1:8080"]
            EP2["10.0.1.2:8080"]
            EP3["10.0.2.1:8080"]
        end

        POD1["🐳 Pod 1\n10.0.1.1"]
        POD2["🐳 Pod 2\n10.0.1.2"]
        POD3["🐳 Pod 3\n10.0.2.1"]

        CP -->|"updates"| ES
        ES -.->|"associated to"| SVC
        EP1 -.->|"IP of"| POD1
        EP2 -.->|"IP of"| POD2
        EP3 -.->|"IP of"| POD3
    end

    KUBECTL -->|"updates"| ES

    style CLUSTER fill:#e8f4fd,stroke:#2196F3,stroke-width:2px
    style SVC fill:#e3f2fd,stroke:#1565C0,stroke-width:2px
    style ES fill:#e8f5e9,stroke:#388E3C,stroke-width:2px
    style CP fill:#f3e5f5,stroke:#9C27B0
    style KUBECTL fill:#fff8e1,stroke:#FFC107
    style EP1 fill:#c8e6c9,stroke:#2E7D32
    style EP2 fill:#c8e6c9,stroke:#2E7D32
    style EP3 fill:#c8e6c9,stroke:#2E7D32
    style POD1 fill:#fff3e0,stroke:#FF9800
    style POD2 fill:#fff3e0,stroke:#FF9800
    style POD3 fill:#fff3e0,stroke:#FF9800
```

