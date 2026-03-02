---
title: EndpointSlice
id: endpoint-slice
full_link: /docs/concepts/services-networking/endpoint-slices/
short_description: >
  EndpointSlices track the IP addresses of Pods for Services.

aka:
tags:
- networking
---

- := [Kubernetes Object](object.md) /
  - TODO: == slice / subset of backing network endpoints for a [Service](service.md)
    - 👁️Kubernetes track number of [Endpoints](endpoint.md) /  EndpointSlices 👁️
      - By default 100 endpoints / EndpointSlices
        **Note:** 👁️ If more endpoints are added → Kubernetes makes a new EndpointSlice 👁️
  - naming rules
    - valid anything
    - 1! name / namespace
  - setting the labels
    - `kubernetes.io/serviceName` — links to the → service
    - `endpointslice.kubernetes.io/managed-by`
      - for EndpointSlices created by your own
      - values
        - `“controller”` — managed by Kubernetes —
        - `"my-domain.example/name-of-controller”` — managed by your own controller —
        - `“nameOfTheToolInLowerCaseReplacingSpacesByDashes”` — managed by a TP tool —
        - `“randomNameToDescribeTheTool”` — managed by `kubectl` —
          *Example:* `“staff”` , `“cluster-admins”`
  - If you use Kubernetes APIs for service discovery → You can query the [API server](../Overview/Kubernetes%20Components%20cc7b751f553341049dd7c054085c18f7.md) for matching EndpointSlices
  - If the set of Pods in a Service changes → Kubernetes updates the EndpointSlices
  - [Service](Service%207251b1e8fc6c4f869671495a6dd6eaec.md) — can be linked to —   ≥ 1  [EndpointSlices](EndpointSlices%20d7743e6ed5654ed8a4c5aca430296c6c.md)
- Uses
  - services
    - ← are linked to → Pods
    - scale — to handle — large numbers of backends
  - cluster updates its list of healthy backends


EndpointSlices track the IP addresses of backend endpoints.
EndpointSlices are normally associated with a
{{< glossary_tooltip text="Service" term_id="service" >}} and the backend endpoints typically represent
{{< glossary_tooltip text="Pods" term_id="pod" >}}.

<!--more-->
One Service can be backed by multiple Pods. Kubernetes represents the backing endpoints of a Service
with a set of EndpointSlices that are associated with that Service.
The backing endpoints are usually, but not always, pods running in the cluster.

The control plane usually manages EndpointSlices for you automatically. However,
EndpointSlices can be defined manually for {{< glossary_tooltip text="Services" term_id="service" >}} without
{{< glossary_tooltip text="selectors" term_id="selector" >}} specified.
