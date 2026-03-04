---
reviewers:
- bgrant0607
- mikedanese
title: "Overview"
description: >
  Kubernetes is a portable, extensible, open source platform for managing containerized workloads and services that facilitate both declarative configuration and automation. It has a large, rapidly growing ecosystem. Kubernetes services, support, and tools are widely available.
content_type: concept
weight: 20
card:
  name: concepts
  weight: 10
  anchors:
  - anchor: "#why-you-need-kubernetes-and-what-can-it-do"
    title: Why Kubernetes?
no_list: true
---

## Why you need Kubernetes & what it can do {#why-you-need-kubernetes-and-what-can-it-do}

* container management | production environment
  * BEFORE Kubernetes,  
    * PROBLEMS: ⚠️MANUALLY, guarantee NO downtime⚠️

* Kubernetes
  * == framework /
    * enable you, to
      * run DISTRIBUTED systems RESILIENTLY
  - Features
    - Automated rollouts & rollbacks
      - rolls out 
        - := gradual changes of
          - application
          - configuration
      - rollback
        - 👀️if something goes wrong | rollout → AUTOMATIC rollback👀
          - Reason: 🧠thanks to monitor application health /
            - ensure NOT ALL instances are killed | same time🧠
    - Service discovery & load balancing
      - own IP address / pod
      - 1! DNS / set of pods
      - load balance ACROSS the pods
      - service discovery provided natively
        - == ❌NO need to use an unfamiliar service discovery mechanism❌
      - [MORE](../services-networking)
    - Storage orchestration
      - you choose the storage system / mount 
        - local
        - public cloud provider
          - _Example:_ AWS, GCP
        - network
          - _Example:_ NFS, Cinder, Ceph, iSCSI
      - [MORE](../storage)
    - Self-healing
      - ==
        - if containers fail → restart
        - replace containers
        - once containers are ready to serve -> advertise the containers to clients 
        - if containers do NOT respond to user-defined health-check → kill containers
    - Secret & configuration management
      - are updated/deployed WITHOUT
        - rebuilding the image
        - exposing secrets
      - _Example of sensitive information:_ passwords, OAuth tokens, SSH keys
      - [MORE](../../reference/glossary/secret.md)
    - Automatic bin packing
      - based on your [resources](../../reference/glossary/infrastructure-resource.md), 
        - how to fit containers | your nodes
        - how to distribute resources | your containers
      - [MORE](../configuration/manage-resources-containers.md)
    - Batch execution
      - if a container fails → it’s replaced
      - [jobs](../../reference/glossary/job.md)
    - Horizontal scaling
      - Scale up or down the application 
        - based on
          - CPU usage
        - -- via --
          - MANUALLY (commands, UI)
          - AUTOMATICALLY
      - [Horizontal Pod Autoscaling](Kubernetes/Documentation/Tasks/Run%20Applications/Horizontal%20Pod%20Autoscaling%2009ea3f3035ba41db8d3feca844baa3f1.md)
    - IPv4/IPv6 dual-stack
      - allocate IPv4 & IPv6 addresses to
        - Pods &
        - Services
      - [MORE](../services-networking/dual-stack.md)
    - [extendable](../extend-kubernetes)

## What Kubernetes is not

- != ❌PaaS system❌
  - 👀ALTHOUGH provide SOME PaaS' feature👀
    - Reason:🧠Kubernetes 
      - operates | container level
      - NOT operate | hardware level🧠
    - _Examples:_ deployment, scaling, load balancing, integrate with logging + monitoring + alerting solutions
  - NO PaaS features
    - Kubernetes is NOT monolithic
      - Reason: 🧠[it's extendable](../extend-kubernetes)
  - _Examples of PaaS:_ Heroku

- ❌Kubernetes does NOT❌
  - limit the types of supported applications
    - [Workloads](../workloads)
  - deploy source code
  - build your application
  - provide 
    - application-level services
      - _Example:_ middleware, databases, data-processing frameworks, caches, ...
      - BUT, these application-level service can be
        - run | Kubernetes
        - accessed by applications running on Kubernetes
          - -- through -- portable mechanisms
            - _Example:_ [Open Service Broker](https://openservicebrokerapi.org/)
    - logging / monitoring / alerting solutions
      - ONLY some integrations -- as -- POC
  - force a configuration language-system
    - Reason: 🧠Kubernetes API can be declared -- by -- many forms👁️
  - provide / adopt systems of
    - maintenance
    - management
    - self-healing
  - orchestrate simply
    - orchestrator
      - := executor of a defined workflow
    - Reason:🧠== control processes / drive from current state -- to -- desired state🧠

## Historical context for Kubernetes {#going-back-in-time}

![Deployment evolution](/static/images/docs/Container_Evolution.svg)

### Traditional deployment era

- applications running | physical servers
- problems
  - ❌NO way to define applications' resource boundaries❌
    - -> resource allocation issues
    - _Example1:_ if >1 applications run | SAME physical server → some instances could take most of the resources → others would underperform
      - _Attempt1:_ 1 application / physical server → 
        - DIFFICULT to scale resources (expensive, maintenance)
        - underuse resources
      - _Solution1:_ virtualization

### Virtualized deployment era

- VM
  - := FULL machine
    - == run ALL components (⚠️even own OS⚠️) | virtual hardware

- virtualization
  - allows
    - set of physical resources == cluster of VM
    - MULTIPLE VM / 1 physical server
    - 👀VM1’s applicationX — is isolated from — VM2’s applicationY👀
      - -> application1’s information — can NOT be FREELY accessed by — application2’s information
    - better
      - use of physical server's resources 
        - == reduce hardware costs
        - set of physical resources — are presented as a — cluster of VM
      - scalability
        - Reason:🧠an application can be added OR updated EASILY🧠
  - cons
    - ❌NOT portable❌
      - Reason:🧠coupled -- , due to Hypervisor, to the -- underlying infrastructure🧠
        - _ExampleS:_ 
          - AWS hypervisor: Xen
          - Azure hypervisor: Hyper-V
          - ...

### Container deployment era

* [container](../../reference/glossary/container.md)
