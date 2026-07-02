---
title: Liveness, Readiness, and Startup Probes
content_type: concept
weight: 40
---

* types of probes
  * [Liveness probe](#liveness-probe)
  * [Readiness probe](#readiness-probe)
  * [Startup probe](#startup-probe)

## Liveness probe

* Liveness probes
  * responsible for
    * determine when to restart a container
      * _Example:_ deadlock
        * == application / is running BUT UNABLE to make progress
  * if a container fails its liveness probe REPEATEDLY -> the kubelet restarts the container
  * ⚠️succeed INDEPENDENTLY of readiness probes⚠️
  * ways to delay liveness probe execution
    * define `initialDelaySeconds`
    * use a [startup probe](#startup-probe)

## Readiness probe

* Readiness probes
  * responsible for
    * determine when a container is ready -- to -- accept traffic
  * use cases
    * application / perform time-consuming initial tasks
      * _Example:_ establish network connections, load files, and warming caches
    * | container’s lifecycle
      * _Example:_ | recover -- from -- temporary faults OR overloads
  * if it returns a failed state -> Kubernetes removes the pod | ALL matching service endpoints
  * run | container | its WHOLE lifecycle

## Startup probe

* startup probe 
  * == probe /
    * ⚠️ONLY executed | startup⚠️
  * responsible for
    * 👀verifies whether the application | container, is started👀
      * if the startup probe NEVER succeeds -> the container is 
        * killed
        * subject -- to the -- pod's `restartPolicy`
  * uses
    * adopt liveness checks | slow starting containers / avoid getting -- , by the kubelet, -- killed | BEFORE they are up & running
  * use cases
    * applications / SLOW FIRST initialization
      * Reason:🧠if you try to set up liveness probe parameters -> they can compromise the fast response to deadlocks🧠
  * ⚠️if you configure it -> NOT set | `livenessProbe` OR `readinessProbe`, `initialDelaySeconds`⚠️
  * vs liveness & readiness probes
    * those are run periodically
      * == AFTER succeed startup probe -> liveness probe provides a fast response -- to -- container deadlocks

* [MORE](/content/en/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes.md)
