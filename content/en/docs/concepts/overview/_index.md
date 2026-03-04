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
  - Storage orchestration
    - you choose the storage system / mount 
      - local
      - public cloud provider
        - _Example:_ AWS, GCP
      - network
        - _Example:_ NFS, Cinder, Ceph, iSCSI
  - Self-healing
    - about
      - if containers fail → restart
      - if nodes die → replace / reschedule containers
      - if containers do NOT respond to user-defined health-check → kill
  - Secret and configuration management
    - are updated/deployed WITHOUT
      - rebuilding the image
      - exposing secrets
  - Automatic bin packing
    - TODO: Containers are placed automatically based on
      **Note:** without sacrificing availability
      - resource requirements &
      - other constraints
    - [**Resource Management for Pods and Containers**](Kubernetes/Documentation/Concepts/Configuration/Resource%20Management%20for%20Pods%20and%20Containers%20dec00252d8ee41369f43063968c9d90d.md)
  - Batch execution
    - Batch and CI workloads
      - if a container fails → it’s replaced
    - [Jobs](Kubernetes/Documentation/Concepts/Workloads/Workload%20Resources/Jobs%20f854f3a6b1dc48ad8ee77a7623b1d785.md)
  - Horizontal scaling
    - Scale up or down the application based on
      - commands
      - UI
      - CPU usage
    - [Horizontal Pod Autoscaling](Kubernetes/Documentation/Tasks/Run%20Applications/Horizontal%20Pod%20Autoscaling%2009ea3f3035ba41db8d3feca844baa3f1.md)
  - IPv4/IPv6 dual-stack
    - allocate IPv4 & IPv6 addresses to
      - Pods &
      - Services
    - [**IPv4/IPv6 dual-stack**](Kubernetes/Documentation/Concepts/Services,%20Load%20Balancing%20and%20Networking/IPv4%20IPv6%20dual-stack%20d9b71fc134c84c4bbae1a430fc1d5f77.md)
  - [extendable](../extend-kubernetes)

Kubernetes provides you with:

* **Storage orchestration**
  
* **Automatic bin packing**
  You provide Kubernetes with a cluster of nodes that it can use to run containerized tasks.
  You tell Kubernetes how much CPU and memory (RAM) each container needs. Kubernetes can fit
  containers onto your nodes to make the best use of your resources.
* **Self-healing**
  Kubernetes restarts containers that fail, replaces containers, kills containers that don't
  respond to your user-defined health check, and doesn't advertise them to clients until they
  are ready to serve.
* **Secret and configuration management**
  Kubernetes lets you store and manage sensitive information, such as passwords, OAuth tokens,
  and SSH keys. You can deploy and update secrets and application configuration without
  rebuilding your container images, and without exposing secrets in your stack configuration.
* **Batch execution**
  In addition to services, Kubernetes can manage your batch and CI workloads, replacing containers that fail, if desired.
* **Horizontal scaling**
  Scale your application up and down with a simple command, with a UI, or automatically based on CPU usage.
* **IPv4/IPv6 dual-stack**
  Allocation of IPv4 and IPv6 addresses to Pods and Services
* **Designed for extensibility**
  Add features to your Kubernetes cluster without changing upstream source code.

## What Kubernetes is not

* != PaaS (Platform as a Service) system
    * traditional
    * all-inclusive 

TODO: 
However, Kubernetes is not monolithic, and these default solutions
    are optional and pluggable. Kubernetes provides the building blocks for building developer
    platforms, but preserves user choice and flexibility where it is important.

- != ❌all-inclusive PaaS❌
  - 👀ALTHOUGH provide SOME PaaS' feature👀
    - Reason:🧠Kubernetes 
      - operates | container level
      - NOT operate | hardware level🧠
    - _Examples:_ deployment, scaling, load balancing, integrate with logging + monitoring + alerting solutions
  - no PaaS features
    - ≠ monolithic
      - === [Extending Kubernetes](Extending%20Kubernetes%20465ae404131c433295cee5c74cf656d9.md)

Kubernetes:

* Does not limit the types of applications supported. Kubernetes aims to support an
  extremely diverse variety of workloads, including stateless, stateful, and data-processing
  workloads. If an application can run in a container, it should run great on Kubernetes.
* Does not deploy source code and does not build your application. Continuous Integration,
  Delivery, and Deployment (CI/CD) workflows are determined by organization cultures and
  preferences as well as technical requirements.
* Does not provide application-level services, such as middleware (for example, message buses),
  data-processing frameworks (for example, Spark), databases (for example, MySQL), caches, nor
  cluster storage systems (for example, Ceph) as built-in services. Such components can run on
  Kubernetes, and/or can be accessed by applications running on Kubernetes through portable
  mechanisms, such as the [Open Service Broker](https://openservicebrokerapi.org/).
* Does not dictate logging, monitoring, or alerting solutions. It provides some integrations
  as proof of concept, and mechanisms to collect and export metrics.
* Does not provide nor mandate a configuration language/system (for example, Jsonnet). It provides
  a declarative API that may be targeted by arbitrary forms of declarative specifications.
* Does not provide nor adopt any comprehensive machine configuration, maintenance, management,
  or self-healing systems.
* Additionally, Kubernetes is not a mere orchestration system. In fact, it eliminates the need
  for orchestration. The technical definition of orchestration is execution of a defined workflow:
  first do A, then B, then C. In contrast, Kubernetes comprises a set of independent, composable
  control processes that continuously drive the current state towards the provided desired state.
  It shouldn't matter how you get from A to C. Centralized control is also not required. This
  results in a system that is easier to use and more powerful, robust, resilient, and extensible.

- Kubernetes
  - allows
    - running distributed systems resiliently
      - scaling
      - failover
  - [features](../../../Kubernetes%208bb068da389f426783032ff97a9bb3af.md)



- ⚠️ Kubernetes does **NOT** ⚠️
  - limit the types of supported applications
    *Example:* variety of [Workloads](Workloads%207f133c71ced8439a97de8ca6b4c88026.md)
  - deploy source code
  - build your application
  - provide application-level services
    *Example:* middleware, databases, …
  - dictate logging / monitoring / alerting
    **Note:** 👁️ provide some integrations as proof of concepts 👁️
  - mandate a configuration language-system
    **Note:** 👁️ Kubernetes API can be targeted by many declarative specifications 👁️
  - provide / adopt systems of
    - maintenance
    - management
    - self-healing
  - merely an orchestration system
    - orchestrator
      - := executor of a defined workflow
    - control processes to drive from current state → desired state

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

Containers have become popular because they provide extra benefits, such as:

* Agile application creation and deployment: increased ease and efficiency of
  container image creation compared to VM image use.
* Continuous development, integration, and deployment: provides reliable
  and frequent container image build and deployment with quick and efficient
  rollbacks (due to image immutability).
* Dev and Ops separation of concerns: create application container images at
  build/release time rather than deployment time, thereby decoupling
  applications from infrastructure.
* Observability: not only surfaces OS-level information and metrics, but also
  application health and other signals.
* Environmental consistency across development, testing, and production: runs
  the same on a laptop as it does in the cloud.
* Cloud and OS distribution portability: runs on Ubuntu, RHEL, CoreOS, on-premises,
  on major public clouds, and anywhere else.
* Application-centric management: raises the level of abstraction from running an
  OS on virtual hardware to running an application on an OS using logical resources.
* Loosely coupled, distributed, elastic, liberated micro-services: applications are
  broken into smaller, independent pieces and can be deployed and managed dynamically –
  not a monolithic stack running on one big single-purpose machine.
* Resource isolation: predictable application performance.
* Resource utilization: high efficiency and density.

- Container deployment era
    - agile application creation and deployment
      **Note:** vs VM
    - CD / CI / Continuous Deployment
    - Dev vs Ops separation
    - Observability
      - OS-level information
      - application-level information
    - [portable](../../Documentation%202e1fbf09afca4abca8f42170786a36be.md)
    - application-centric management
      **Reason:** 🧠 VM’s focus is to run an OS  — vs — Container’s focus is to run an application on an OS 🧠
  - in a production environment, you need to
    **Note:** 👁️ an alternative is to use Kubernetes 👁️
    - manage the containers which run
    - ensure no downtime

## {{% heading "whatsnext" %}}

* Take a look at the [Kubernetes Components](/docs/concepts/overview/components/)
* Take a look at the [The Kubernetes API](/docs/concepts/overview/kubernetes-api/)
* Take a look at the [Cluster Architecture](/docs/concepts/architecture/)
* Ready to [Get Started](/docs/setup/)?
