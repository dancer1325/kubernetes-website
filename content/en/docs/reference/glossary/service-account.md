---
title: ServiceAccount
id: service-account
full_link: /docs/tasks/configure-pod-container/configure-service-account/
short_description: >
  Provides an identity for processes that run in a Pod.

aka: 
tags:
- fundamental
- core-object
---

* 👀exist 1 serviceAccount / EACH namespace👀
  * | create a Pod, if you do NOT specify a service account -> use AUTOMATICALLY the default service account | that [namespace](namespace.md)
* provides
  * an identity -- for -- processes / run | [pod](pod.md)
* uses
  * pod communicate -- , authenticating thanks to serviceAccount, with -- the API server
