---
title: Group Version Resource
id: gvr
short_description: >
  The API group, API version and name of a Kubernetes API.

aka: ["GVR"]
tags:
- architecture
---

* Group Version Resources (GVRs)
  * == specify API group + API version + API resource
  * allows
    * 👀representing UNIQUELY specific Kubernetes APIs👀
      * EXCEPT to APIs / namespaced
  * _Example:_ "apps/v1/deployments"
    * apps  
      * Group
    * v1
      * version
    * deployments
      * resource
