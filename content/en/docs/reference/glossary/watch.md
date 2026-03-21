---
title: Watch
id: watch
full_link: /docs/reference/using-api/api-concepts/#api-verbs
short_description: >
  A verb that is used to track changes to an object in Kubernetes as a stream.

aka:
tags:
- API verb
- fundamental
---

* watch
  * == verb /   
    * ❌!= poll❌
  * uses
    * track Kubernetes object change -- as a -- stream
      * _Example:_ controller watches ConfigMap changes
  * allows
    * [optimize detection of changes](../using-api/api-concepts.md#efficient-detection-of-changes)
