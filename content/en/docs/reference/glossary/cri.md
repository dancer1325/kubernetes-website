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
  - allows
    - pulling the container image — from — a registry
    - unpacking the container
    - running the application
  - == [gRPC](https://grpc.io) protocol /
    - Kubelet uses to communicate -- with -- CR
  - implementations / supported by Kubernetes
    - [CRI-O](https://cri-o.io/#what-is-cri-o)
    - [containerd](https://containerd.io/docs/)
