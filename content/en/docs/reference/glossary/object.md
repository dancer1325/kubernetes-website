---
title: Object
id: object
full_link: /docs/concepts/overview/working-with-objects/#kubernetes-objects
short_description: >
   An entity in the Kubernetes system, representing part of the state of your cluster.
aka:
tags:
- architecture
- fundamental
---

- object
  - := persistent entity | Kubernetes system
    - persistent
      - === once you create the object → Kubernetes works to ensure it’s existence
        - Reason: 🧠Kubernetes is declarative🧠
    - expressed | `.yaml`
    - if you want to operate with them -> use the [Kubernetes API](kubernetes-api.md)
    - uses
      - represent the cluster’s state
        - _Examples:_ 
          - containerized applications
          - available application’s resources
          - application’s policies
            - restart policies
            - upgrades
            - fault-tolerance
            - …

- nested objects  
  - ALMOST ALL Kubernetes object
  - `.spec`
    - NORMALLY MANDATORY
    - == desired state
      - specified -- by the -- user
      - specific / Kubernetes object
        - *ExampleS:* [PodSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podspec-v1-core), [StatefulSetSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#statefulsetspec-v1-apps)
  - `status`
    - == current state
      - handled -- by -- Kubernetes
      - ❌NOT specified -- by the -- user❌
      - specific / Kubernetes object
        - *Example:* [PodStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podstatus-v1-core), [StatefulSetStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#statefulsetstatus-v1-apps)
  - `apiVersion`
    - Kubernetes API Group / API Version
  - `kind`
    - type of object
  - `metadata`
    - := data -- to -- identify UNIQUELY the Kubernetes object
  - [Kubernetes API Conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md)
