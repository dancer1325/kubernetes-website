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
  - == 💡Control Plane's core component💡
    - == FE -- for -- Control Plane
  - responsible for
    - 👀exposes the Kubernetes API👀
      - == implement the [Kubernetes API](kubernetes-api.md)
    - handles the conversion between [API versions](../using-api/_index.md#api-versioning)
    - track _managed fields_ / ALL NEWLY created objects

- kube-apiserver
  - == Kubernetes API server's main implementation
  - 's design
    - scale horizontally (== deploy MORE instances)
      - if you run >1 kube-apiserver instance → balance traffic BETWEEN those
  - 👁️can be deployed --as -- [container image](image.md) | cluster 👁️
