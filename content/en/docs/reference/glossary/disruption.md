---
title: Disruption
id: disruption
full_link: /docs/concepts/workloads/pods/disruptions/
short_description: >
  An event that leads to Pod(s) going out of service
aka:
tags:
- fundamental
---

* := events / cause >=1 pod go out of service
  * -> consequences for [workload management resources](api-resource.md)
    * _Example:_ [deployment](deployment.md) / rely on the affected pods
* types
  * voluntary
    * _Example:_ cluster operator destroys a pod / belongs -- to an -- application 
  * involuntary
    * _Example:_ pod goes offline -- due to -- 
      * node failure
      * outage
* [MORE](../../concepts/workloads/pods/disruptions.md)
