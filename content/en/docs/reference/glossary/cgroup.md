---
title: cgroup (control group)
id: cgroup
full_link:
short_description: >
  A group of Linux processes with optional resource isolation, accounting and limits.

aka:
tags:
- fundamental
---

# cgroup
- cgroup
  - := Linux kernel feature /
    - | group of Linux processes, can about [resources](infrastructure-resource.md) (CPU/memory/disk I/O) be 
      - restricted
      - accounted
      - isolated
  - EXISTING versions
    - v1
    - v2
      - NEW generation

# cgroup drivers
- cgroup drivers
  - allows
    - interface with cgroup
  - uses
    - kubelet & kubelet's CR
      - requirements
        - BOTH use the SAME cgroup driver / SAME configuration
      - can -- , thanks to cgroup drivers, --
        - enforce [resource management | Pods & Containers**](../../concepts/configuration/manage-resources-containers.md)
        - set [resources (CPU/memory)](resource-quota.md)
          - _Example:_ requests & limits
  - types
    - cgroupfs driver
      - kubelet’s default cgroup driver
      - if it’s used → kubelet & Kubelet's CR — interface directly via cgroup filesystem to — configure cgroups
      - if [systemd](https://www.freedesktop.org/wiki/Software/systemd/) is the init system
        - ❌ NOT recommended to use cgroupfs❌
          - **Reason:** 🧠systemd expects 1! cgroup type | system🧠
        - \+ cgroupfs driver → 2 DIFFERENT cgroup managers
          - **Reason:** 🧠systemd
            1. — tightly integrated with — cgroups
            2. allocates 1 cgroup / systemd unit 🧠
          - == 2 views of the resources available & in-use | system
      - if nodes are configured / kubelet & CR use cgroupfs driver + rest of the processes use systemd → | resource pressure, unstable 
        - **Solution:** ⚠️kubelet's systemd cgroup driver & kubelet's CR systemd cgroup driver🧠
      - 👁️If you use [Kubeadm](../../Reference/Setup%20tools/Kubeadm%200af94946cb364c4d96df7338a0a22751.md) with cgroupfs driver and managed cluster + you want to migrate to [systemd cgroup driver](Container%20Runtimes%2016213c35786e48d7aecf9ddf4ef0b686.md) → [**Configuring a cgroup driver**](../../Tasks/Administer%20a%20Cluster/Administration%20with%20kubeadm/Configuring%20a%20cgroup%20driver%20532a3232770d4c4b8de5c4c421e91da9.md)  👁️
    - systemd cgroup driver
      - ⚠️if cgroup v2 is used → use systemd cgroup ⚠️
      - if systemd is used -- as -- init system -> systemd
        - generates & consumes a root cgroup &
        - acts as cgroup manager
      - steps to configure

        ```yaml
        apiVersion: kubelet.config.k8s.io/v1beta1
        kind: KubeletConfiguration
          ...
          cgroupDriver: systemd
          ...
        ```

      - | Kubernetes v1.22+,
        - if you install the cluster via [Kubeadm](../../Reference/Setup%20tools/Kubeadm%200af94946cb364c4d96df7338a0a22751.md) → by default use systemd cgroup driver 👁️
      - if you configure kubelet's cgroup driver = systemd — you must configure kubelet's CR's cgroup driver = systemd 
  - recommendations
    - | existing node (== certain cgroup driver configured), if you want to change the cgroup drive →
      - ❌NOT recommended to change the cgroup driver❌
      - 👁️replace the node👁️
  - if you enable the `KubeletCgroupDriverFromCRI` feature gate + using a CR / 
    - supports `RuntimeConfig` CRI RPC -> kubelet
      - AUTOMATICALLY detects the appropriate cgroup driver | runtime
      - ignores the kubelet configuration's `cgroupDriver` setting
    - NOT support `RuntimeConfig` CRI RPC 
      - _Example:_ containerd v1-
      - | Kubernetes v1.36-
        - -> kubelet use the `--cgroup-driver` flag value
      - | Kubernetes v1.36+
        - -> kubelet fails
