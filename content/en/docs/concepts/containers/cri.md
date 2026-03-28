---
title: Container Runtime Interface (CRI)
content_type: concept
weight: 60
---

* [CRI](../../reference/glossary/cri.md)

## The API {#api}

{{< feature-state for_k8s_version="v1.23" state="stable" >}}

* [CRI protocol definition](https://github.com/kubernetes/cri-api/blob/v0.33.1/pkg/apis/runtime/v1/api.proto)
  * services
    * runtime service
      * endpoints,
        * by default, == image service
        * if you want to configure -> `kubelet ----container-runtime-endpoint`
    * image service

## Upgrading

TODO:
If a gRPC re-dial is required because the container
runtime has been upgraded, the runtime must support the `v1` CRI API for the
connection to succeed
This might require a restart of the kubelet after the
container runtime is correctly configured.
