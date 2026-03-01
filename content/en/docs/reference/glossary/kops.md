---
title: kOps (Kubernetes Operations)
id: kops
full_link: /docs/setup/production-environment/kops/
short_description: >
  kOps will not only help you create, destroy, upgrade and maintain production-grade, highly available, Kubernetes cluster, but it will also provision the necessary cloud infrastructure.

aka: 
tags:
- tool
- operation
---

- `kOps`/Kubernetes Operations
  - == automated provisioning system /
    - fully automated installation
    * identify clusters -- via -- DNS
    * self-healing
      * == everything runs | Auto-Scaling Groups
    * MULTIPLE OS support (Amazon Linux, Debian, Flatcar, RHEL, Rocky and Ubuntu)
    * High-Availability support
    * DIRECTLY provision OR generate terraform manifests
  - enable you,
    - about Kubernetes cluster,
      - create
      - destroy
      - upgrade
      - maintain
        - production-grade
        - highly available 
    - provision the necessary cloud infrastructure
  - supported ones
    - officially
      - AWS
    - beta
      - DigitalOcean,
      - GCE
      - OpenStack
    - alpha
      - Azure
