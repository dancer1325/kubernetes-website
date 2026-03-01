---
title: Node
id: node
full_link: /docs/concepts/architecture/nodes/
short_description: >
  A node is a worker machine in Kubernetes.

aka:
tags:
- core-object
- fundamental
---

* node
  * == worker machine | Kubernetes
  * | Kubernetes early versions,
    * "Minions"
  * types
    * [control plane / master node](control-plane.md)
    * worker node

* worker node
  * types 
    * -- based on the -- cluster
      * VM
      * physical machine 
  * run a [pod](pod.md) -- thanks to -- 
    * local daemons OR
    * services 
  * managed -- by the -- control plane

* daemons
  * := OS services
  * | node
    * 👀include👀
      * [kubelet](kubelet.md)
      * [kube-proxy](kube-proxy.md)
      * container runtime / implement the [CRI](cri.md)
        * _Example:_ [docker](docker.md)

![Untitled](/static/images/docs/node.png)
