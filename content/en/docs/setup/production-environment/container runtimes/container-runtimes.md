---
reviewers:
- vincepri
- bart0sh
title: Container Runtimes
content_type: concept
weight: 20
---

* goal
  * how to use [container runtimes](../../../reference/glossary/container-runtime.md) -- for -- setting up nodes
    * [CRI version support](#cri-version-support-cri-versions)

* requirements
  * install a container runtime | EACH cluster's node /
    * Reason:🧠run pods | EACH cluster's node🧠
    * conforms -- with the -- [CRI](../../../reference/glossary/cri.md)

## Install and configure prerequisites

### Network configuration

* Kubernetes cluster networking implementations
  * require IPv4 packets can be routed BETWEEN interfaces
    * Linux kernel
      * by default, ❌NOT allowed❌
    * MOST change this setting
    * SOME might require MANUALLY
  * SOME, expect other
    * sysctl parameters
    * kernel modules
    * etc

### Enable IPv4 packet forwarding {#prerequisite-ipv4-forwarding-optional}

* MANUALLY enable IPv4 packet forwarding

  ```bash
  # sysctl params required by setup, params persist across reboots
  cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
  net.ipv4.ip_forward = 1
  EOF
  
  # Apply sysctl params without reboot
  sudo sysctl --system
  ```

* verify / `net.ipv4.ip_forward` == 1

  ```bash
  sysctl net.ipv4.ip_forward
  ```

## cgroup drivers

- [here](../../../reference/glossary/cgroup.md)

## CRI version support {#cri-versions}

* requirements
  * CRI v1alpha2+
  * | [Kubernetes v1.26+](/content/en/blog/_posts/2022/kubernetes-1.26-deprecations-and-removals.md)
    * CRI v1

* if you use CRI BEFORE v1alpha2 -> use v1
  * if a container runtime does NOT support v1 -> kubelet falls back to v1alpha2

## Container runtimes

* [dockershim removal](../../../reference/glossary/dockershim.md)
  * [how to migrate](../../../tasks/administer-cluster/migrating-from-dockershim/)

### containerd

* goal
  * how to use [containerd](../../../reference/glossary/container.md) -- as -- CRI runtime

* steps
  * [install](https://github.com/containerd/containerd/blob/main/docs/getting-started.md)
    * valid "config.toml" has been created
      * | Linux: "/etc/containerd/config.toml"
      * | Windows: "C:\Program Files\containerd\config.toml"
    * default CRI
      * socket | Linux: "/run/containerd/containerd.sock"
      * endpoint | Windows: "npipe://./pipe/containerd-containerd"

#### Configuring the `systemd` cgroup driver {#containerd-systemd}

To use the `systemd` cgroup driver in `/etc/containerd/config.toml` with `runc`,
set the following config based on your Containerd version

Containerd versions 1.x:

```
[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
  ...
  [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
    SystemdCgroup = true
```

Containerd versions 2.x:

```
[plugins.'io.containerd.cri.v1.runtime'.containerd.runtimes.runc]
  ...
  [plugins.'io.containerd.cri.v1.runtime'.containerd.runtimes.runc.options]
    SystemdCgroup = true
```

The `systemd` cgroup driver is recommended if you use [cgroup v2](/docs/concepts/architecture/cgroups).

{{< note >}}
If you installed containerd from a package (for example, RPM or `.deb`), you may find
that the CRI integration plugin is disabled by default.

You need CRI support enabled to use containerd with Kubernetes
* Make sure that `cri`
is not included in the`disabled_plugins` list within `/etc/containerd/config.toml`;
if you made changes to that file, also restart `containerd`.

If you experience container crash loops after the initial cluster installation or after
installing a CNI, the containerd configuration provided with the package might contain
incompatible configuration parameters
* Consider resetting the containerd configuration
with `containerd config default > /etc/containerd/config.toml` as specified in
[getting-started.md](https://github.com/containerd/containerd/blob/main/docs/getting-started.md#advanced-topics)
and then set the configuration parameters specified above accordingly.
{{< /note >}}

If you apply this change, make sure to restart containerd:

```shell
sudo systemctl restart containerd
```

* if you use kubeadm -> MANUALLY configure the [cgroup driver for kubelet](../../../tasks/administer-cluster/kubeadm/configure-cgroup-driver.md#configuring-the-kubelet-cgroup-driver)

In Kubernetes v1.28, you can enable automatic detection of the
cgroup driver as an alpha feature
* See [systemd cgroup driver](#systemd-cgroup-driver)
for more details.

#### Overriding the sandbox (pause) image {#override-pause-image-containerd}

In your [containerd config](https://github.com/containerd/containerd/blob/main/docs/cri/config.md) you can overwrite the
sandbox image by setting the following config:

```toml
[plugins."io.containerd.grpc.v1.cri"]
  sandbox_image = "registry.k8s.io/pause:3.10"
```

You might need to restart `containerd` as well once you've updated the config file: `systemctl restart containerd`.

### CRI-O

* goal
  * how to [install CRI-O](https://github.com/cri-o/packaging/blob/main/README.md#usage) -- as a -- container runtime

* requirements
  * Linux

#### cgroup driver

* by default,
  * use systemd cgroup driver

* To switch to the `cgroupfs` cgroup driver, either edit
`/etc/crio/crio.conf` or place a drop-in configuration in
`/etc/crio/crio.conf.d/02-cgroup-manager.conf`, for example:

```toml
[crio.runtime]
conmon_cgroup = "pod"
cgroup_manager = "cgroupfs"
```

You should also note the changed `conmon_cgroup`, which has to be set to the value
`pod` when using CRI-O with `cgroupfs`
* It is generally necessary to keep the
cgroup driver configuration of the kubelet (usually done via kubeadm) and CRI-O
in sync.

In Kubernetes v1.28, you can enable automatic detection of the
cgroup driver as an alpha feature
* See [systemd cgroup driver](#systemd-cgroup-driver)
for more details.

For CRI-O, the CRI socket is `/var/run/crio/crio.sock` by default.

#### Overriding the sandbox (pause) image {#override-pause-image-cri-o}

In your [CRI-O config](https://github.com/cri-o/cri-o/blob/main/docs/crio.conf.5.md) you can set the following
config value:

```toml
[crio.image]
pause_image="registry.k8s.io/pause:3.10"
```

This config option supports live configuration reload to apply this change: `systemctl reload crio` or by sending
`SIGHUP` to the `crio` process.

### Docker Engine {#docker}

- requirements
  - [cri-dockerd](https://github.com/Mirantis/cri-dockerd)

- cri-dockerd
  - := adapter /
    - allows, 
      - Docker Engine — is integrated with — Kubernetes
  - CRI socket
    - by default, "/run/cri-dockerd.sock"

* steps
  * | EACH cluster's node
  1. [install Docker Engine](https://docs.docker.com/engine/install/#server)
  2. [install `cri-dockerd`](https://mirantis.github.io/cri-dockerd/usage/install)

### Mirantis Container Runtime (MCR) {#mcr}

* [overview](https://docs.mirantis.com/mcr/25.0/overview.html) 
* commercial
* origin
  * PREVIOUS name: Docker Enterprise Edition

- requirements
  - [cri-dockerd](https://github.com/Mirantis/cri-dockerd)
    - included -- with -- MCR
    - if you want to specify the CI to use -- as -- pod infrastructure container -> use the CL argument `--pod-infra-container-image`

- [Deployment Guide](https://docs.mirantis.com/mcr/25.0/install.html)

* if you want to find out the path to the CRI socket -> check `cri-docker.socket`
  * `cri-docker.socket` == systemd unit

## {{% heading "whatsnext" %}}

As well as a container runtime, your cluster will need a working
[network plugin](/docs/concepts/cluster-administration/networking/#how-to-implement-the-kubernetes-network-model).
