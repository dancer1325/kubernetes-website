---
title: Storage Class
id: storageclass
full_link: /docs/concepts/storage/storage-classes
short_description: >
  A StorageClass provides a way for administrators to describe different available storage types.

aka: 
tags:
- core-object
- storage
---

* StorageClass
  * == available storage types /
    * can map -- to --
      * quality-of-service levels
      * backup policies
      * arbitrary policies 
  * 's fields
    * are
      * `provisioner`
      * `parameters`
      * `reclaimPolicy`
    * uses
      * class' [PV](persistent-volume.md) needs to be DYNAMICALLY provisioned
  * uses
    * by administrators
  * use cases
    * | PVC, specified -- by -- StorageClass name
