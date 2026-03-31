---
title: Compute, Storage, and Networking Extensions
weight: 30
no_list: true
---

* goal
  * cluster extensions / ❌NOT part -- as -- Kubernetes itself❌
    * uses
      * enhance the nodes | your cluster
      * provide the network fabric / links Pods together

* storage plugins
  * [CSI](../../../concepts/storage/volumes.md#csi)
    * [CSI plugins](../../../reference/glossary/csi.md)
    * extend Kubernetes -- through -- supporting NEW kinds of volumes 
  * [FlexVolume](../../../concepts/storage/volumes.md#flexvolume-deprecated-flexvolume)
    * TODO: The volumes can
      be backed by durable external storage, or provide ephemeral storage, or they might offer a
      read-only interface to information using a filesystem paradigm.

      Kubernetes also includes support for [FlexVolume](/docs/concepts/storage/volumes/#flexvolume)
      plugins, which are deprecated since Kubernetes v1.23 (in favour of CSI).

      FlexVolume plugins allow users to mount volume types that aren't natively
      supported by Kubernetes. When you run a Pod that relies on FlexVolume
      storage, the kubelet calls a binary plugin to mount the volume. The archived
      [FlexVolume](https://git.k8s.io/design-proposals-archive/storage/flexvolume-deployment.md)
      design proposal has more detail on this approach.

      * [Kubernetes Volume Plugin FAQ -- for -- Storage Vendors](https://github.com/kubernetes/community/blob/master/sig-storage/volume-plugin-faq.md#kubernetes-volume-plugin-faq-for-storage-vendors)

* [Device plugins](device-plugins)

* [Network plugins](network-plugins)
