---
title: Kubernetes z-pages
content_type: reference
weight: 60
reviewers:
- dashpole
description: >-
  Provides runtime diagnostics for Kubernetes components, offering insights into component runtime status and configuration flags.
---

* requirements
  * Kubernetes v1.32 -> state="alpha"

* z-endpoints
  * == endpoints / 
    * exposed -- by -- Kubernetes core components
    * provide
      * internal information -- about -- running components 
  * use cases
    * users (MANUALLY) can 👀debug their cluster + its components👀
      * MANUALLY == avoid scraping them

## z-pages

For Kubernetes {{< skew currentVersion >}}, components
serve the following endpoints (when enabled):

- [z-pages](#z-pages)
	- [statusz](#statusz)
		- [statusz (structured)](#statusz-structured)
	- [flagz](#flagz)
		- [flagz (structured)](#flagz-structured)

### "/statusz"

* requirements
  * use the `ComponentStatusz` [feature gate](../command-line-tools-reference/feature-gates#ComponentStatusz)
* displays
  * high level information -- about the -- component 
    * _Example:_ Kubernetes version, emulation version, start time, ...

#### "/statusz" -- with --  structured response

* requirements
  * Kubernetes v1.35 / alpha feature
  * request contains `Accept` header: `Accept: application/json;v=v1alpha1;g=config.k8s.io;as=Statusz`
    * if you use ONLY `Accept: application/json` -> respond: `406 Not Acceptable`

* [response's schema](https://github.com/dancer1325/kubernetes/tree/master/staging/src/k8s.io/apiserver/pkg/server/statusz)

### "/flagz"

* requirements
  * use `ComponentFlagz` [feature gate](../command-line-tools-reference/feature-gates#ComponentFlagz)
* displays
  * CL arguments / were used | start a component

#### "/flagz" -- with --  structured response

* requirements
  * Kubernetes v1.35 / alpha feature
  * request contains `Accept` header: `Accept: application/json;v=v1alpha1;g=config.k8s.io;as=Flagz`
    * if you use ONLY `Accept: application/json` -> respond: `406 Not Acceptable`

* [response's schema](https://github.com/dancer1325/kubernetes/tree/master/staging/src/k8s.io/apiserver/pkg/server/flagz)
