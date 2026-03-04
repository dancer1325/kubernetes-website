---
title: Service
id: service
full_link: /docs/concepts/services-networking/service/
short_description: >
  A way to expose an application running on a set of Pods as a network service.
tags:
- fundamental
- core-object
---

* := method /
  * allows
    * exposing a network application / 
      * is running | >=1 cluster's podS
      * targeted set of pods is determined -- by a -- [selector](selector.md)
        * if number of pods is changed -> set of pods / match de selector change
  * uses either
    * IP networking (IPv4, IPv6, or both), OR
    * external name reference | Domain Name System (DNS)
  * makes sure that
    * network traffic can be directed -- to the -- workload's current set of pods
  * enables
    * [Ingress](ingress.md)
    * [Gateway](gateway.md)
