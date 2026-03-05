---
title: Container
id: container
full_link: /docs/concepts/containers/
short_description: >
  A lightweight and portable executable image that contains software and all of its dependencies.

aka: 
tags:
- fundamental
- workload
---

* container
  * := executable image /
    * lightweight
    * portable
      * == consistency ACROSS 
        * environments
        * OS
    * == 👀software + ALL its dependencies👀
    * 's goal
      * stateless
      * [immutable](https://glossary.cncf.io/immutable-infrastructure/)
        * == if you have a containerized application / ALREADY running & want to make changes
          * ❌NO change the code DIRECTLY❌
          * steps
            * build a new image / contains the changes
            * recreate the container -- based on -- updated image
  * 👀decouple applications -- from -- underlying host infrastructure (cloud, OS, ...) 👀
    * == repeatable
      * == you get the SAME behavior | ANYWHERE you run it
    * -> 
      * less separation BETWEEN dev & devOps
      * recommendations
        * create application containers 
          * | build/release time
          * NOT | deployment time
  * allows
    * bundle & run your applications
    * observability |
      * OS-level
      * application-level
  * == OS-level virtualization
  * [vs virtual machines (VMs)](#containers-vs-virtual-machines)
  * ways to expose it
    * DNS name
    * OWN IP address

* containerized applications
  * == applications / run | containers 

* containerization
  * := process of bundling containerized applications + their dependencies | container image

## Containers vs Virtual Machines

| Aspect                      | Containers                                               | Virtual Machines (VMs)                           |
|-----------------------------|----------------------------------------------------------|--------------------------------------------------|
| **Operating System**        | Share the host OS kernel                                 | OWN complete OS / EACH VM                        |
| **Size**                    | Lightweight (10-100 MB)                                  | Heavy (500 MB - several GB)                      |
| **Startup Time**            | Seconds (1-2s)                                           | Minutes (30-60s)                                 |
| **Isolation Level**         | Process-level isolation (relaxed)                        | FULL isolation <br/> == SEPARATE kernel          |
| **Isolated Resources**      | OWN filesystem <br/> share of CPU, memory, process space | OWN filesystem, CPU, memory, process space       |
| **Security**                | Less isolated <br/> Reason: share kernel                 | More isolated <br/> Reason: SEPARATE OS/kernel   |
| **Performance**             | Near-native performance                                  | overhead <br/> Reason: due to hypervisor         |
| **Resource Efficiency**     | High (100 / host)                                        | Lower (10-20 / host)                             |
| **Portability**             | Highly portable <br/> Reason: SAME image everywhere      | Less portable <br/> Reason: hypervisor-dependent |
| **Infrastructure Coupling** | Decoupled (works across clouds/platforms)                | Coupled to specific hypervisor/cloud             |
| **Image Format**            | Standardized (OCI standard)                              | Proprietary (.vmdk, .vhd, .ami)                  |
| **Boot Process**            | Starts a process                                         | Boots entire OS                                  |
| **Memory Footprint**        | Minimal (shares OS)                                      | Large (full OS per VM)                           |
| **Compatibility**           | Requires Linux kernel (for Linux containers)             | Can run different OS types                       |
| **Use Case**                | Microservices, cloud-native apps                         | Legacy apps, different OS requirements           |
| **Management**              | Docker, containerd, Kubernetes                           | vSphere, Hyper-V, KVM                            |
| **Scalability**             | Very fast (seconds to scale)                             | Slower (minutes to provision)                    |
| **Migration**               | Simple (push/pull images)                                | Complex (format conversion needed)               |
| **Dependencies**            | Container runtime + Linux kernel                         | Hypervisor + hardware virtualization             |
| **Update/Patch**            | Replace container (immutable)  -> faster                 | Patch OS inside VM -> slower                     |
| **Networking**              | Shared network with isolation                            | Virtual network adapters                         |
| **Storage**                 | Layered filesystem (copy-on-write)                       | Full virtual disks                               |


![Deployment evolution](/static/images/docs/Container_Evolution.svg)
