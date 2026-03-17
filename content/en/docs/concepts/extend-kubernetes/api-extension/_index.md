---
title: Extending the Kubernetes API
weight: 30
---

* ways to extend the Kubernetes API (== add CUSTOM resources | your cluster)
  * [CustomResourceDefinition](custom-resources) 
  * [aggregation layer](apiserver-aggregation)

| Feature                          | CRDs                                         | API Aggregation                                    |
|----------------------------------|----------------------------------------------|----------------------------------------------------|
| **kube-apiserver recognition**   | kube-apiserver DIRECTLY recognizes NEW kinds | kube-apiserver acts -- as -- proxy ONLY            |
| **Who handles objects**          | kube-apiserver                               | SEPARATE API server                                |
| **ADDITIONAL API server needed** | NO                                           | YES                                                |
| **Programming required**         | NO                                           | YES                                                |
| **Maintenance & bug fixing**     | Kubernetes project handles it                | YOU MUST handle it (fix bugs, rebuild, and update) |

# vs stand-alone API

* ANOTHER approach / you do ❌NOT extend the Kubernetes API (CRD or AA) ❌

| use cases \| extend the Kubernetes API extension ([CRD](custom-resources.md) or [AA](apiserver-aggregation.md))     | use cases \| proceed with stand-alone API                                                                                |
|---------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------|
| Your API is [declarative](../../../reference/glossary/declarative-api.md)                                           | ❌Your API does NOT fit the [Declarative](../../../reference/glossary/declarative-api.md) model❌                          |
| your NEW types MUST be readable & writable -- via -- `kubectl`                                                      | ❌`kubectl` support is NOT required❌                                                                                      |
| your NEW types MUST be displayed \| Kubernetes UI                                                                   | ❌Kubernetes UI support is NOT required❌                                                                                  |
| You are developing a NEW API                                                                                        | exist ALREADY a program / serves your API & works well                                                                   |
| You want to follow the format restriction / [Kubernetes set \| REST resource paths](../../overview/kubernetes-api.md) | You need specific REST paths / compatible -- with the -- ALREADY DEFINED REST API                                        |
| Your resources are scoped -- to a -- cluster OR namespaces                                                          | ❌You do NOT want cluster OR namespace scopes❌ <br/> &nbsp;&nbsp; Reason: 🧠you need control \| specific resource paths🧠 |
| You want to reuse [Kubernetes API features](../../overview/kubernetes-api.md)                                               | ❌You do NOT need Kubernetes API features❌                                                                                |

# vs ConfigMap OR Secret?

* ANOTHER approach / you do ❌NOT need to extend the Kubernetes API❌
  * use cases
    * just store configuration data

| use cases \| extend the Kubernetes API extension ([CRD](custom-resources.md) or [AA](apiserver-aggregation.md))                                   | use cases / add entries \| ConfigMap OR Secret                                                        |
|---------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| You want to create & update -- , through [Kubernetes client libraries & CLIs](../../../reference/glossary/kubernetes-api.md) -- the new resource  | exist well-documented configuration file format <br/>  &nbsp;&nbsp; _Example:_ "mysql.cnf", "pom.xml" |
| You want top-level support -- , WITHOUT using ADDITIONAL plugins OR external tools, from -- `kubectl`                                             | you want that \| update the file, perform rolling updates                                             |
| You want to build NEW automation -- based on -- NEW objects                                                                                       |                                                                                                       |
| You want the object is: abstraction/summarization -- over a -- collection of resources                                                            |                                                                                                       |
| You want to use [Kubernetes API conventions](../../../reference/glossary/kubernetes-api.md) <br/> _Examples:_ `.spec`, `.status`, and `.metadata` |                                                                                                       |

* well-documented configuration file
  * steps
    * place | 1 ConfigMap's key
  * uses
    * by a program / run | pod | your cluster
