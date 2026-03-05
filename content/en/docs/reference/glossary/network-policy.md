---
title: Network Policy
id: network-policy
full_link: /docs/concepts/services-networking/network-policies/
short_description: >
  A specification of how groups of Pods are allowed to communicate with each other and with other network endpoints.

aka: 
tags:
- networking
- architecture
- extension
- core-object
---

* NetworkPolicies
  * == built-in Kubernetes API (-> declarative specification) 
    * allows, 
      * controlling the traffic BETWEEN
        * pods
        * pods -- & -- outside
  * implemented 
    * -- by a -- supported network plugin / 
      * provided -- by a -- network provider 
      * requirements
        * controller / implement it
    * NORMALLY, -- by the -- pod network implementation

* NetworkPolicy objects
  * select pods & define traffic rules -- through -- [labels](label.md)
