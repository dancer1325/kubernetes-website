---
title: StatefulSet
id: statefulset
full_link: /docs/concepts/workloads/controllers/statefulset/
short_description: >
  A StatefulSet manages deployment and scaling of a set of Pods, with durable storage and persistent identifiers for each Pod.

aka: 
tags:
- fundamental
- core-object
- workload
- storage
---

* StatefulSet
  * allows
    * managing the deployment & scaling of a set of [pods](pod.md) /
      * guarantee ordering & uniqueness / EACH pod
      * created -- from the -- SAME spec
      * NOT interchangeable
        * Reason: 🧠tracks a state🧠
      * has a persistent identifier / EACH pod
        * maintained across any rescheduling
        * -> easy to replicate
  * use cases
    * workloads / provide persistence -- via -- storage volumes
      * == ALTHOUGH pod is killed -> state is persisted
  * ALTHOUGH individual pods fail -> you can match existing volumes -- , via persistent pod identifiers, -- to the new pods

```mermaid
flowchart TB
    subgraph CLUSTER["☁️ Kubernetes Cluster"]
        subgraph SS["StatefulSet"]
            SPEC["📋 SAME spec"]
        end

        subgraph PODS["pods"]
            direction LR
            subgraph POD1["pod-0"]
                ID1["🔖 persistent ID"]
                PV1["💾 Volume A"]
            end
            subgraph POD2["pod-1"]
                ID2["🔖 persistent ID"]
                PV2["💾 Volume B"]
            end
            subgraph POD3["pod-2"]
                ID3["🔖 persistent ID"]
                PV3["💾 Volume C"]
            end
        end

        SS -->|"scale up: pod-0 → pod-1 → pod-2"| PODS
        PODS -->|"scale down: pod-2 → pod-1 → pod-0"| SS
        SPEC -->|"same spec"| POD1
        SPEC -->|"same spec"| POD2
        SPEC -->|"same spec"| POD3

        POD1_NEW["pod-0 (rescheduled)"]
        POD1 -.->|"fails & rescheduled\n→ same ID reused"| POD1_NEW
        POD1_NEW -.->|"matched via\npersistent ID"| PV1
    end

    style CLUSTER fill:#e8f4fd,stroke:#2196F3,stroke-width:2px
    style SS fill:#f3e5f5,stroke:#9C27B0,stroke-width:2px
    style SPEC fill:#fff8e1,stroke:#FFC107
    style PODS fill:#e8f5e9,stroke:#388E3C,stroke-width:1px
    style POD1 fill:#e8f5e9,stroke:#4CAF50
    style POD2 fill:#e8f5e9,stroke:#4CAF50
    style POD3 fill:#e8f5e9,stroke:#4CAF50
    style PV1 fill:#fff3e0,stroke:#FF9800
    style PV2 fill:#fff3e0,stroke:#FF9800
    style PV3 fill:#fff3e0,stroke:#FF9800
    style ID1 fill:#c8e6c9,stroke:#2E7D32
    style ID2 fill:#c8e6c9,stroke:#2E7D32
    style ID3 fill:#c8e6c9,stroke:#2E7D32
    style POD1_NEW fill:#fce4ec,stroke:#E53935,stroke-dasharray: 5 5
```

* Deployment vs StatefulSet
  
  |                      | Deployment          | StatefulSet                                                                                |
  |----------------------|---------------------|--------------------------------------------------------------------------------------------|
  | Pod identity         | random, ephemeral   | UNIQUE <br/> &nbsp;&nbsp; SAME name <br/> &nbsp;&nbsp; SAME DNS <br/> &nbsp;&nbsp; SAME PV |
  | Pod creation order   | arbitrary           | sequential <br/> &nbsp;&nbsp; _Example:_ "pod-0", "pod-1", ...                             |
  | Pod deletion order   | arbitrary           | reverse sequential <br/> &nbsp;&nbsp; _Example:_ "pod-N", "pod-N-1", ...                   |
  | Interchangeable pods | ✅ yes              | ❌NO                                                                                        |
  | Volume               | 1 shared / ALL pods | 1 / EACH pod                                                                               |
  | Use case             | stateless workloads | stateful workloads <br/> &nbsp;&nbsp; _ExampleS:_ DBs, queues, ...                         |
