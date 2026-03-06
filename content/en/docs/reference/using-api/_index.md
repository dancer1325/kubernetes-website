---
title: API Overview
reviewers:
- erictune
- lavalamp
- jbeda
content_type: concept
weight: 20
no_list: true
card:
  name: reference
  weight: 50
  title: Overview of API
---

* goal
  * Kubernetes API

## API versioning

- 👀>1 are supported👀
  - DIFFERENT API path / EACH version
    - _Examples:_ 
    - _Examples:_ 
      - /api/v1
      - /apis/rbac.authorization.k8s.io/v1alpha1
  - if your objects are using a version 
    - ⚠️in DEPRECATION period -> transition to the stable one⚠️
    - ⚠️ALREADY NO served -> replace to the NEW API version⚠️
- 👀| API-level👀
  - != resource-level OR field-level⚠️
- follow semVer
  - ❌NO CURRENT plans -- for a -- MAJOR version❌
    - == Kubernetes v1.Y.Z


TODO: 
The JSON and Protobuf serialization schemas follow the same guidelines for
schema changes

* [API versioning vs software versioning](https://git.k8s.io/sig-release/release-engineering/versioning.md)

* DIFFERENT levels of stability & support / EACH API version
  * [MORE](https://git.k8s.io/community/contributors/devel/sig-architecture/api_changes.md#alpha-beta-and-stable-versions).
  * levelS
    - Alpha
      - names
        - contain `alpha`
          - _Example:_ `v1alpha1`
      - disabled, by default
        - if you want to use -> EXPLICITLY enable it | `kube-apiserver`
      - software 
        - may contain bugs
      - may, | ANY time WITHOUT notice 
        - be dropped 
        - be changed in incompatible ways
      - recommendations
        - use it ONLY | short-lived testing clusters
    - Beta:
      - names
        - contain `beta` 
          - _Example:_ `v2beta3`
      - enabling
        - | Kubernetes v1.22-
          - enabled, by default
        - | Kubernetes v1.22+
          - disabled, by default
            - if you want to use -> EXPLICITLY enable it | `kube-apiserver`
      - lifetime
        - introduction -- to -- deprecation
          - <= (9 months, 3 minor releases)
        - deprecation -- to -- removal
          - <= (9 months, 3 minor releases) 
      - software
        - well tested
      - you can enable it, safely
      - support for a feature
        - NOT dropped
        - may change
      - schema & semantics of objects
        - | subsequent beta OR stable API version, may INCOMPATIBLY change 
          - ⚠️if you want to migrate -> may require editing OR re-creating API objects⚠️
      - recommendations
        - ❌NOT use | production❌
          - Reason:🧠NEXT releases may introduce incompatible changes🧠
        - use them (BEFORE transition to stable) & provide feedback
          - Reason:🧠once they are promoted to stable -> may NOT be practical to make MORE changes🧠
    - Stable:
      - `vX` 
        - `X` == integer
        - version name 
      - AVAILABLE -- for -- ALL Kubernetes future releases / SAME major version

## [API groups](https://git.k8s.io/design-proposals-archive/api-machinery/api-group.md)

* allows
  * extend easily the Kubernetes API
* can be enabled OR disabled

The API group is specified in a REST path and in the `apiVersion` field of a
serialized object.

There are several API groups in Kubernetes:

*  The *core* (also called *legacy*) group is found at REST path `/api/v1`.
   The core group is not specified as part of the `apiVersion` field, for
   example, `apiVersion: v1`.
*  The named groups are at REST path `/apis/$GROUP_NAME/$VERSION` and use
   `apiVersion: $GROUP_NAME/$VERSION` (for example, `apiVersion: batch/v1`).
   You can find the full list of supported API groups in
   [Kubernetes API reference](/docs/reference/generated/kubernetes-api/{{< param "version" >}}/#-strong-api-groups-strong-).

## Enabling or disabling API groups   {#enabling-or-disabling}

Certain resources and API groups are enabled by default. You can enable or
disable them by setting `--runtime-config` on the API server.  The
`--runtime-config` flag accepts comma separated `<key>[=<value>]` pairs
describing the runtime configuration of the API server. If the `=<value>`
part is omitted, it is treated as if `=true` is specified. For example:

 - to disable `batch/v1`, set `--runtime-config=batch/v1=false`
 - to enable `batch/v2alpha1`, set `--runtime-config=batch/v2alpha1`
 - to enable a specific version of an API, such as `storage.k8s.io/v1beta1/csistoragecapacities`, set `--runtime-config=storage.k8s.io/v1beta1/csistoragecapacities`

{{< note >}}
When you enable or disable groups or resources, you need to restart the API
server and controller manager to pick up the `--runtime-config` changes.
{{< /note >}}

## Persistence

Kubernetes stores its serialized state in terms of the API resources by writing them into
{{< glossary_tooltip term_id="etcd" >}}.

## {{% heading "whatsnext" %}}

- Learn more about [API conventions](https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#api-conventions)
- Read the design documentation for
  [aggregator](https://git.k8s.io/design-proposals-archive/api-machinery/aggregated-api-servers.md)
- Learn about [Declarative API Validation](/docs/reference/using-api/declarative-validation/).