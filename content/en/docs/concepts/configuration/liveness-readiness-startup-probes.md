---
title: Liveness, Readiness, and Startup Probes
content_type: concept
weight: 40
---

<!-- overview -->

Kubernetes has various types of probes:

- [Liveness probe](#liveness-probe)
- [Readiness probe](#readiness-probe)
- [Startup probe](#startup-probe)

<!-- body -->

## Liveness probe

Liveness probes determine when to restart a container
* For example, liveness probes could catch a deadlock when an application is running but unable to make progress.

If a container fails its liveness probe repeatedly, the kubelet restarts the container.

Liveness probes do not wait for readiness probes to succeed
* If you want to wait before executing a liveness probe, you can either define `initialDelaySeconds` or use a
[startup probe](#startup-probe).


## Readiness probe

Readiness probes determine when a container is ready to accept traffic
* This is useful when waiting for an application to perform time-consuming initial tasks that depend on its backing services; for example: establishing network connections, loading files, and warming caches
* Readiness probes can also be useful later in the container’s lifecycle, for example, when recovering from temporary faults or overloads.

If the readiness probe returns a failed state, Kubernetes removes the pod from all matching service endpoints.

Readiness probes run on the container during its whole lifecycle.


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
