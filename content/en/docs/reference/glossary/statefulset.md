---
title: StatefulSet
id: statefulset
full_link: /docs/concepts/workloads/controllers/statefulset/
short_description: >
  A StatefulSet manages deployment and scaling of a set of Pods, with durable storage and persistent identifiers for each Pod.

aka: 
tags:
- fundamental
- core-object
- workload
- storage
---

* StatefulSet
  * allows
    * managing the deployment & scaling of a set of [pods](pod.md) /
      * guarantee ordering & uniqueness / EACH pod
      * created -- from the -- SAME spec
      * NOT interchangeable
        * Reason: 🧠tracks a state🧠
      * has a persistent identifier / EACH pod
        * maintained across any rescheduling
        * -> easy to replicate
  * use cases
    * workloads / provide persistence -- via -- storage volumes
  * ALTHOUGH individual pods fail -> you can match existing volumes -- , via persistent pod identifiers, -- to the new pods

* vs deployment
  * BOTH manage replicated pods
  * StatefulSet
    * maintains a sticky identity / EACH pod
      * _Example:_ 1 PV / EACH pod
  * deployment
    * 1 volume / ALL pods
  