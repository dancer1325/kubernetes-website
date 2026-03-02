---
title: API Group
id: api-group
full_link: /docs/concepts/overview/kubernetes-api/#api-groups-and-versioning
short_description: >
  A set of related paths in the Kubernetes API.

aka:
tags:
- fundamental
- architecture
---

* API group
  * == set of related paths | Kubernetes API /
    * specified | 
      * REST path
      * serialized object's `.apiVersion` field
  * allows
    * | configure your API server, enable OR disable
      * EACH API group 
      * paths -- to -- specific [resources](api-resource.md)
    * extending the Kubernetes API
  * [MORE](../../concepts/overview/kubernetes-api.md#api-groups-and-versioning)
