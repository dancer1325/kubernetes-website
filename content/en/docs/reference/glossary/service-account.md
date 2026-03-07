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

* := Kubernetes built-in kind /
  * 💡managed by Kubernetes💡
  * ⚠️namespace-scope⚠️
    * == exist 1 serviceAccount (default one) / EACH namespace
      * if you do NOT specify a service account -> use AUTOMATICALLY the default service account | that namespace
  * tied to credentialS / stored as `Secret` / mounted | pods
    * -> in-cluster processes can talk -- to the -- Kubernetes API
  * provides
    * an identity -- for -- processes / run | [pod](pod.md)
* uses
  * pod communicate -- , authenticating thanks to serviceAccount, with -- the API server
