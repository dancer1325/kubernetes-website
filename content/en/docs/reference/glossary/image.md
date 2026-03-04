---
title: Image
id: image
full_link: 
short_description: >
  Stored instance of a container that holds a set of software needed to run an application.

aka: 
tags:
- fundamental
---

* == way of packaging software (== binary data) /
  * hold a software / can run an application
    * -> contains: application code + runtime + system libraries + application's dependencies + default values about settings
  * can be
    * stored | container registry
      * ⚠️required if you want to use it | pod⚠️
    * pulled -- to a -- local system
    * run -- as an -- application
  * 's metadata
    * _Example:_
      * what executable to run
      * who built it

* == executable software bundles / 
  * can run standalone
  * make very well-defined assumptions about their runtime environment

* images / run | Kubernetes
  * accept ALL OS (ALSO Windows)
