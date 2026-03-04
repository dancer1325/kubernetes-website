---
title: Container Runtime
id: container-runtime
full_link: /docs/setup/production-environment/container-runtimes
short_description: >
 The container runtime is the software that is responsible for running containers.

aka:
tags:
- fundamental
- workload
---

- container runtime
  - := Kubernetes' component /
    - manages, about containers,
      - execution
      - lifecycle
    - 1 is configured by default | ANY Kubernetes cluster
  - if you need to use >1 CR | Kubernetes cluster -> specify [RuntimeClass](../../concepts/containers/runtime-class.md)
