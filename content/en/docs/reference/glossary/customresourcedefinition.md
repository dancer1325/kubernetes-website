---
title: CustomResourceDefinition
id: CustomResourceDefinition
full_link: /docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/
short_description: >
  Custom code that defines a resource to add to your Kubernetes API server without building a complete custom server.

aka: 
tags:
- fundamental
- operation
- extension
---

* CustomResourceDefinitions (CRD)
  * := kind of [Kubernetes object](object.md) /
    * defines DECLARATIVELY a NEW CUSTOM API (== [API group](api-group.md) + [API version](group-version-resource.md) + schema + ...)
    * added -- , ⚠️WITHOUT building a COMPLETE CUSTOM API server⚠️, -- | your Kubernetes API server
  * allow you
    * extend the Kubernetes API | your environment
  * use cases
    * ❌there is NO API resource / meet your needs❌
