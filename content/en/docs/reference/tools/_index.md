---
title: Other Tools
reviewers:
- janetkuo
content_type: concept
weight: 150
no_list: true
---

* tools /
  * help you work -- with the -- Kubernetes system

## [crictl](https://github.com/kubernetes-sigs/cri-tools)

* CL interface /
  * allows
    * inspecting & debugging [container runtimes](../glossary/container-runtime.md) [CRI](../glossary/cri.md)-compatible

## [Dashboard](https://github.com/kubernetes/dashboard)

* == Kubernetes UI
  * web-based
    * by default, ONLY accessible | Kubernetes virtual network
  * ⚠️deprecated & unmaintained⚠️
    * use [Headlamp](#headlamp)
  * allows
    * manage (create, modify, view, ...) 
      * cluster resources
      * containerized applications | Kubernetes cluster
    * troubleshoot your containerized application
* [tasks](../../tasks/access-application-cluster/web-ui-dashboard.md)

![](/static/images/docs/ui-dashboard.png)

## [Headlamp](https://headlamp.dev/)

* == Kubernetes cluster component
  * OPTIONAL
  * provide
    * Kubernetes graphical UI
      * extensible -- via -- plugin system 
      * allows
        * cluster management & troubleshooting
    * RBAC-based controls
  * uses 
    * | cluster
    * as desktop application

## Helm
{{% thirdparty-content single="true" %}}

[Helm](https://helm.sh/) is a tool for managing packages of pre-configured
Kubernetes resources. These packages are known as _Helm charts_.

Use Helm to:

* Find and use popular software packaged as Kubernetes charts
* Share your own applications as Kubernetes charts
* Create reproducible builds of your Kubernetes applications
* Intelligently manage your Kubernetes manifest files
* Manage releases of Helm packages

## [Kompose](https://github.com/kubernetes/kompose)

* == tool /
  * help Docker Compose users move -- to -- Kubernetes
    * translate a Docker Compose file (v1 OR v2 OR [distributed application bundles](https://docs.docker.com/compose/bundles/)) -- into -- Kubernetes objects
    * manage -- , via Kubernetes, -- your application

## [Kui](https://github.com/kubernetes-sigs/kui)

* == GUI tool /
  * 's input
    * your normal `kubectl` CL requests
  * 's output
    * graphics / contain tables
  * allows
    * executing `kubectl` commands
    * querying a [job](../glossary/job.md) 

## Minikube

[`minikube`](../glossary/minikube.md)
