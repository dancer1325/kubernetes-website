---
title: "Services, Load Balancing, and Networking"
weight: 60
description: >
  Concepts and resources behind networking in Kubernetes.
---

## The Kubernetes network model

- Kubernetes network model
  - characteristics
    - 1! IP / EACH cluster's pod
      - ⚠️cluster-wide⚠️
        - != node-scope
      - 👁️ you don’t need manually to 👁️
        **Reason:** 🧠 it’s abstracted and managed by Kubernetes  🧠
        - pod1 ← link → pod2
        - container’s ports ← mapping to → [host’s ports](Services,%20Load%20Balancing%20and%20Networking%2093f52609b9cd4c7a924b63a2c5f80cab.md)
      - pods == VMs / Physical Hosts from perspective of
        - port allocation
        - naming
        - service discovery
        - load balancing
        - application configuration
        - migration
    - pod network OR cluster network
      - namespace setup
        - system-level software implement the [CRI](../containers/cri.md)
      - itself
        - managed -- by -- a [pod network implementation](../cluster-administration/addons.md#networking-and-network-policy)
          - | Linux, MOST container runtimes use  _CNI plugins_
      - handle communication -- , WITHOUT using proxies OR NATs, -- BETWEEN pods /
        - node1's pod1 can communicate -- with -- node1's pod2
          - | Windows,
            - NOT valid | host-network pods
        - node1's pod1 can communicate -- with -- node2's pod2
          - | Windows,
            - NOT valid | host-network pods
      - node’s agent can communicate -- with — ALL nod's pod 
    - service proxying
      - [kube-proxy](../../reference/glossary/kube-proxy.md)
        - default Kubernetes implementation
      - SOME pod network implementations use their own service
        - Reason:🧠MORE integrated with the rest of the implementation🧠
    - [network policies](../../reference/glossary/network-policy.md)
    - [Gateway API](gateway) 
      - if you use a [supported cloud provider](../../reference/glossary/cloud-provider.md) & want simpler & less-configurable -> use Service API's [`type: LoadBalancer`](service.md#loadbalancer)
    - [Service API](service.md)
    - OWN private network namespace / EACH pod
      - shared by ALL containers | pod
      - pod1's container1 can communicate -- , through "localhost", with -- pod1's container2 
  - implementation
    - SOME are implemented by Kubernetes
    - ⚠️SOME is provided -- , through Kubernetes API, by -- other parts⚠️
  - networking implementation
    - available | DIFFERENT CR (Container Runtime)

* history
  * | older container systems,
    * ❌there was NO AUTOMATIC connectivity BETWEEN containers | DIFFERENT hosts❌
      * SOLUTION:
        * solution1: explicitly create links BETWEEN containers
        * solution2: map container ports -- to -- host ports
          * Reason:🧠make them reachable by containers | other hosts🧠
  * | Kubernetes
    * pods == VMs OR physical hosts, about
      * port allocation
      * naming
      * service discovery
      * load balancing
      * application configuration
      * migration

## {{% heading "whatsnext" %}}

* [how to connect applications -- via -- Services](../../tutorials/services/connect-applications-service.md)

* [Cluster Networking](../cluster-administration/networking.md) 
  * how to set up networking | your cluster
