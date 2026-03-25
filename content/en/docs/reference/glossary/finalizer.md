---
title: Finalizer
id: finalizer
full_link: /docs/concepts/overview/working-with-objects/finalizers/
short_description: >
  A namespaced key that tells Kubernetes to wait until specific conditions are met
  before it fully deletes an object marked for deletion.
aka: 
tags:
- fundamental
- operation
---

* finalizers
  * := 💡namespaced keys /
    * related -- to -- controller's actions💡
    * allows
      * alert [controllers](controller.md) -- to -- perform cleanup work | BEFORE resource deletion
        * _Example:_ create backups, delete external resources, ...
      * block deletion UNTIL cleanup complete
    * uses
      * control [garbage collection](garbage-collection.md) of resources
  * workflow
    1. user wants to delete a resource / contains a finalizer specified → Kubernetes API
       * mark the object for deletion
          * == populate `.metadata.deletionTimestamp`
       * returns a `202` status code (HTTP "Accepted")
    2. TILL control plane OR other components take the finalizers actions, object (to be deleted)'s state == `Terminating`
    3. ONCE 
       * finalizer actions complete → controller removes that finalizer | object (to be deleted)'s `metadata.finalizers`
       * ALL finalizers are removed → resource can be deleted
