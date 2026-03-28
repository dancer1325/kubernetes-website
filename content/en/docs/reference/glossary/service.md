---
title: Service
id: service
full_link: /docs/concepts/services-networking/service/
short_description: >
  A way to expose an application running on a set of Pods as a network service.
tags:
- fundamental
- core-object
---

* Service
  * := method /
    * allows
      * exposing a network application / 
        * is running | >=1 cluster's podS
        * targeted set of pods is determined -- by a -- [selector](selector.md)
          * ⚠️if number of pods is changed -> set of pods / match de selector change⚠️
  * ALLOWED mechanisms
    
    | Mechanism                             | Service types                           |
    |---------------------------------------|-----------------------------------------|
    | IP networking (IPv4 / IPv6 / both)    | `ClusterIP`, `NodePort`, `LoadBalancer` |
    | External name reference -- via -- DNS | `ExternalName`                          |

  * enables: [Ingress](ingress.md) & [Gateway](gateway.md)
    * == ⚠️require a service⚠️
  * 👀default service👀
    * name: kubernetes
    * type: ClusterIP
    * | default namespace
    * created | install Kubernetes

```mermaid
flowchart TB
    subgraph CLUSTER["☁️ Kubernetes Cluster"]
        subgraph DEFAULT_NS["default namespace"]
            DEFSVC["Service: kubernetes\ntype: ClusterIP\n(created | install Kubernetes)"]
        end

        subgraph OTHER_NS["my-namespace"]
            INGRESS["Ingress"]
            GATEWAY["Gateway"]

            P1["Pod1 ✅\napp: my-app"]
            P2["Pod2 ✅\napp: my-app"]
            P3["Pod3 ✅\napp: my-app"]

            subgraph SVC["📡 Service"]
                SELECTOR["selector\napp: my-app"]

                subgraph NETWORKING["ALLOWED mechanisms"]
                    direction LR
                    IP["IP networking\n (IPv4 / IPv6)"]
                    DNS["external name reference \n -- via -- \n DNS"]
                end
            end
        end
    end

    INGRESS -->|"requires Service\nas backend"| SVC
    GATEWAY -->|"requires Service\nas backend"| SVC

    SELECTOR -->|"determines targeted pods"| P1
    SELECTOR -->|"determines targeted pods"| P2
    SELECTOR -->|"determines targeted pods"| P3

    style CLUSTER fill:#e8f4fd,stroke:#2196F3,stroke-width:2px
    style DEFAULT_NS fill:#fce4ec,stroke:#E53935
    style DEFSVC fill:#fce4ec,stroke:#E53935
    style OTHER_NS fill:#fff3e0,stroke:#FF9800
    style SVC fill:#ede7f6,stroke:#7B1FA2,stroke-width:2px
    style SELECTOR fill:#fff8e1,stroke:#FFC107
    style NETWORKING fill:#e3f2fd,stroke:#1565C0
    style IP fill:#e8f5e9,stroke:#4CAF50
    style DNS fill:#e8f5e9,stroke:#4CAF50
    style INGRESS fill:#fff3e0,stroke:#FF9800
    style GATEWAY fill:#fff3e0,stroke:#FF9800
```
