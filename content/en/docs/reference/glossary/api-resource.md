---
title: API resource
id: api-resource
full_link: /docs/reference/using-api/api-concepts/#standard-api-terminology
short_description: >
  A Kubernetes entity, representing an endpoint on the Kubernetes API server.

aka:
 - Resource
tags:
- architecture
---

* API resource
  * := entity | Kubernetes / 👀correspond -- to a -- Kubernetes API's endpoint👀
    * can represent
      * objects
        * _ExampleS:_ 
          * /api/v1/<resource-type>
          * /api/v1/namespaces/<namespace>/<resource-type>
      * operation | other objects
        * _ExampleS:_
          * /api/v1/namespaces/default/pods/nginx/eviction
          * /apis/authorization.k8s.io/v1/subjectaccessreviews

* ALL | Kubernetes 
  * is treated -- as an -- API object
