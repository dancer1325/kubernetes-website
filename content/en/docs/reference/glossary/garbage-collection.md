---
title: Garbage Collection
id: garbage-collection
full_link: /docs/concepts/architecture/garbage-collection/
short_description: >
  A collective term for the various mechanisms Kubernetes uses to clean up cluster
  resources.

aka: 
tags:
- fundamental
- operation
---

* Garbage collection
  * == collective term -- about -- MULTIPLE Kubernetes mechanisms / 
    * 👀clean up cluster resources👀
      * [unused containers & images](/content/en/docs/concepts/architecture/garbage-collection.md#garbage-collection-of-unused-containers-and-images-containers-images)
      * [terminated Pods](/content/en/docs/concepts/workloads/pods/pod-lifecycle.md#garbage-collection-of-pods-pod-garbage-collection)
      * [objects / owned -- by the -- targeted resource](/content/en/docs/concepts/overview/working-with-objects/owners-dependents.md)
      * [completed Jobs](/content/en/docs/concepts/workloads/workload%20resources/ttlafterfinished.md)
      * [Dynamically provisioned PV / StorageClass reclaim policy of Delete](/content/en/docs/concepts/storage/persistent-volumes.md#delete)
      * [Stale OR expired CertificateSigningRequests (CSRs)](/content/en/docs/reference/access-authn-authz/certificate-signing-requests.md#request-signing-process)
      * [Node Lease objects](/content/en/docs/concepts/architecture/nodes.md#node-heartbeats)
      * nodes / deleted |
        * cloud / cluster uses a [cloud controller manager](cloud-controller-manager.md)
        * On-premises / cluster uses an addon similar == [cloud controller manager](cloud-controller-manager.md)

