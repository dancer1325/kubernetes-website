---
reviewers:
- mikedanese
content_type: concept
title: Tools for Monitoring Resources
weight: 15
---

* goal
  * understand the application's behavior | being deployed
    * Reason:🧠lets you
      * scale an application
      * provide a reliable service🧠

* if you want to examine application performance | Kubernetes cluster -> examine the application's resource usage |
  * containers level
  * [pods](../../../concepts/workloads/pods/) level
  * [services](../../../concepts/services-networking/service.md) level
  * overall cluster's characteristics level

* | Kubernetes,
  * you can implement ANY application monitoring 
    * Reason:🧠[Kubernetes is agnostic to it](../../../concepts/overview/_index.md)🧠
  * | NEW clusters, if you want to collect monitoring statistics -> you can use
    * [resource metrics](#resource-metrics-pipeline)
    * [full metrics](#full-metrics-pipeline) pipelines 

## Resource metrics pipeline

* resource metrics pipeline
  * provides
    * a limited set of metrics -- related to -- cluster components 
      * 👀collected -- by -- [metrics-server](https://github.com/kubernetes-sigs/metrics-server)👀
      * exposed -- via -- `metrics.k8s.io` API
      * _Examples:_ 
        * [Horizontal Pod Autoscaler controller](../../run-application/horizontal-pod-autoscale-walkthrough.md)
        * `kubectl top` utility

* metrics-server
  * lightweight
  * short-term
  * in-memory
  * discovers ALL cluster's nodes
  * queries EACH node's [kubelet](../../../reference/command-line-tools-reference/kubelet) -- for --
    * CPU usage
    * memory usage

* TODO: The kubelet acts as a bridge between the Kubernetes master and 
the nodes, managing the pods and containers running on a machine
* The kubelet 
translates each pod into its constituent containers and fetches individual 
container usage statistics from the container runtime through the container 
runtime interface
* If you use a container runtime that uses Linux cgroups and
namespaces to implement containers, and the container runtime does not publish
usage statistics, then the kubelet can look up those statistics directly
(using code from [cAdvisor](https://github.com/google/cadvisor)).
No matter how those statistics arrive, the kubelet then exposes the aggregated pod
resource usage statistics through the metrics-server Resource Metrics API.
This API is served at `/metrics/resource/v1beta1` on the kubelet's authenticated and 
read-only ports. 

## Full metrics pipeline

* full metrics pipeline 
  * fetches metrics -- from -- kubelet
  * exposes the metrics -- to -- Kubernetes
    * via -- , by implementing `custom.metrics.k8s.io` OR `external.metrics.k8s.io` API, an -- adapter
  * provides
    * richer metrics
  * use cases
    * Kubernetes respond -- to -- these metrics
      * _Examples:_ AUTOMATICALLY -- , based on cluster's current state, --
        * scaling the cluster
        * adapting the cluster

* Kubernetes' design
  * work with [OpenMetrics](https://openmetrics.io/)
    * Reason:🧠part of [CNCF Observability and Analysis - Monitoring Projects](https://landscape.cncf.io/?group=projects-and-products&view-mode=card#observability-and-analysis--monitoring)🧠

* MANY [CNCF monitoring projects can work -- , by scraping data, with -- Kubernetes](https://landscape.cncf.io/?group=projects-and-products&view-mode=card#observability-and-analysis--monitoring),
  * select the tool OR toolS / suit your needs

TODO: 
When you design and implement a full metrics pipeline you can make that monitoring data
available back to Kubernetes
* For example, a HorizontalPodAutoscaler can use the processed
metrics to work out how many Pods to run for a component of your workload.

Integration of a full metrics pipeline into your Kubernetes implementation is outside
the scope of Kubernetes documentation because of the very wide scope of possible
solutions.

The choice of monitoring platform depends heavily on your needs, budget, and technical resources.
Kubernetes does not recommend any specific metrics pipeline; [many options](https://landscape.cncf.io/?group=projects-and-products&view-mode=card#observability-and-analysis--monitoring) are available.
Your monitoring system should be capable of handling the [OpenMetrics](https://openmetrics.io/) metrics
transmission standard and needs to be chosen to best fit into your overall design and deployment of
your infrastructure platform.

records generic time-series metrics about containers in a central database, and provides a UI for browsing that data.

## {{% heading "whatsnext" %}}


Learn about additional debugging tools, including:

* [Logging](/docs/concepts/cluster-administration/logging/)
* [Getting into containers via `exec`](/docs/tasks/debug/debug-application/get-shell-running-container/)
* [Connecting to containers via proxies](/docs/tasks/extend-kubernetes/http-proxy-access-api/)
* [Connecting to containers via port forwarding](/docs/tasks/access-application-cluster/port-forward-access-application-cluster/)
* [Inspect Kubernetes node with crictl](/docs/tasks/debug/debug-cluster/crictl/)
