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

- Kubernetes object
  - := persistent entity | Kubernetes system
    - persistent
      - == | create the object -> Kubernetes works to ensure it’s existence
        - Reason: 🧠Kubernetes is declarative🧠
    - can be
      - expressed | ".yaml"
      - represented | Kubernetes API
        - if you want to operate with them -> use the Kubernetes API
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
  - 's fields
    - `.spec`
      - nested object 
      - if you want to create a Kubernetes object -> MANDATORY to specify
    - `.status`
      - nested object
    - `.apiVersion`
      - if you want to create a Kubernetes object -> MANDATORY to specify
    - `kind`
      - if you want to create a Kubernetes object -> MANDATORY to specify
    - `metadata`
      - nested object
      - if you want to create a Kubernetes object -> MANDATORY to specify
  - [good practices](/content/en/blog/_posts/2025/kubernetes-configuration-best-practices)

- `.spec`
  - == object
    - == 👀Kubernetes object's desired state👀
      - specified -- by the -- user
      - [control plane tries to reach this desired state always](control-plane.md)
    - specific / Kubernetes object
      - _ExampleS:_ [PodSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podspec-v1-core), [StatefulSetSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#statefulsetspec-v1-apps)

- `status`
  - == object
    - == Kubernetes object's current state
      - handled -- by -- Kubernetes
        - == ❌NOT specified -- by the -- user❌
      - specific / Kubernetes object
        - *Example:* [PodStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#podstatus-v1-core), [StatefulSetStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#statefulsetstatus-v1-apps)

- `apiVersion`
  - Kubernetes API Group / API Version

- `kind`
  - type of object

- `metadata`
  - := data -- to -- identify UNIQUELY the Kubernetes object
    - _Example:_ `.name`, `.namespace`

- [Kubernetes API Conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md)
