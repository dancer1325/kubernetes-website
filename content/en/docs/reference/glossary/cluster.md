---
title: Cluster
id: cluster
full_link: 
short_description: >
   A set of worker machines, called nodes, that run containerized applications. Every cluster has at least one worker node.

aka: 
tags:
- fundamental
- operation
---

* cluster
  * := set of [nodes](node.md) / run containerized applications
    * | production environments,
      * NORMALLY runs MULTIPLE (>=3) nodes
        * Reason: 🧠provide fault-tolerance & high availability🧠
  * exist >= 1 [worker node](node.md) / EACH cluster
    * Reason: 🧠they run pods🧠
  * _Example:_ Kubernetes cluster

```
┌─────────────────────────────────────────────────────────────────┐
│                    KUBERNETES CLUSTER                           │
│                                                                 │
│  ┌──────────────────────────────────────────────────────────┐  │
│  │              CONTROL PLANE (Brain)                       │  │
│  │                                                          │  │
│  │  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  │  │
│  │  │ API Server   │  │  Scheduler   │  │  Controller  │  │  │
│  │  │              │  │              │  │   Manager    │  │  │
│  │  └──────────────┘  └──────────────┘  └──────────────┘  │  │
│  │                                                          │  │
│  │  ┌──────────────┐                                       │  │
│  │  │    etcd      │  ← Cluster database                  │  │
│  │  └──────────────┘                                       │  │
│  └────────────────┬─────────────────────────────────────────┘  │
│                   │                                            │
│                   │ Orders and coordinates                     │
│                   ▼                                            │
│  ┌────────────────────────────────────────────────────────┐   │
│  │               WORKER NODES                             │   │
│  │                                                        │   │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐   │   │
│  │  │   Node 1    │  │   Node 2    │  │   Node 3    │   │   │
│  │  │             │  │             │  │             │   │   │
│  │  │ ┌─────────┐ │  │ ┌─────────┐ │  │ ┌─────────┐ │   │   │
│  │  │ │  Pod    │ │  │ │  Pod    │ │  │ │  Pod    │ │   │   │
│  │  │ │ ┌─────┐ │ │  │ │ ┌─────┐ │ │  │ │ ┌─────┐ │ │   │   │
│  │  │ │ │ App │ │ │  │ │ │ App │ │ │  │ │ │ App │ │ │   │   │
│  │  │ │ └─────┘ │ │  │ │ └─────┘ │ │  │ │ └─────┘ │ │   │   │
│  │  │ └─────────┘ │  │ └─────────┘ │  │ └─────────┘ │   │   │
│  │  │             │  │             │  │             │   │   │
│  │  │ ┌─────────┐ │  │ ┌─────────┐ │  │ ┌─────────┐ │   │   │
│  │  │ │  Pod    │ │  │ │  Pod    │ │  │ │  Pod    │ │   │   │
│  │  │ └─────────┘ │  │ └─────────┘ │  │ └─────────┘ │   │   │
│  │  └─────────────┘  └─────────────┘  └─────────────┘   │   │
│  └────────────────────────────────────────────────────────┘   │
│                                                                │
└────────────────────────────────────────────────────────────────┘
```
