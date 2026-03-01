---
title: API server
id: kube-apiserver
full_link: /docs/concepts/architecture/#kube-apiserver
short_description: >
  Control plane component that serves the Kubernetes API.

aka:
- kube-apiserver
tags:
- architecture
- fundamental
---

- Kubernetes API server
  - == 💡Control Plane's core component💡 /
    - exposes the Kubernetes API
    - == FE -- for -- Control Plane
    - TODO: handles the conversion between [API versions](../../Reference/API%20Overview%2088a30423709544f8b4d00b65b49f0e64.md) transparently
      - 👁️ different API versions === different representations of the same persisted data 👁️
        - **Note:** 👁️ You can check the [groups available / versions in the documentation](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#api-groups) 👁️
        - *Example:* If same resource has `v1` and `v1beta1`, and you create a resource with `v1` → it can be operated with `v1beta1`

- kube-apiserver
  - == Kubernetes API server's main implementation
  - 's design
    - scale horizontally (== deploy MORE instances)
      - if you run >1 kube-apiserver instance → balance traffic BETWEEN those
  - TODO: Check [here](../../Reference/Component%20tools/kube-apiserver%200441e85c0c8a4d9f8efbdb0a7d21587c.md)
  - 👁️ can be deployed as [CI](https://kubernetes.io/releases/download/#container-images) within the cluster 👁️