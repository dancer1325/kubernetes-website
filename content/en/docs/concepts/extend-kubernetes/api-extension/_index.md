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
