---
title: Kubernetes API Aggregation Layer
reviewers:
- lavalamp
- cheftako
- chenopis
content_type: concept
weight: 20
---

- API Aggregation (AA)
  - == technical approach
  - ⚠️requirements⚠️
    - programming
      - Reason: 🧠code ADDITIONAL API server🧠
  - allows
    - 👀MORE controlling | API👀
      - _Examples:_ 
        - how store data
        - how convert BETWEEN API versions
    - extend Kubernetes API -- with -- ADDITIONAL APIs
      - types of ADDITIONAL APIs
        - ready-made solutions
          - _Example:_ [metrics server](https://github.com/kubernetes-sigs/metrics-server)
        - develop yourself
  - use cases
    - you need flexibility/customization

## Aggregation layer

- Aggregation layer
  - == 💡SUBORDINATED (!= PRIMARY one) API server💡 /
    - sit behind PRIMARY API server
      - == 👀if you request for your CUSTOM APIs -> PRIMARY one delegate -- , acting as a proxy, through the `APIService`, to -- SUBORDINATED API server 👀
    - ⚠️you write + deploy + maintain ⚠️
    - 👀runs |  kube-apiserver 👀
  - ⚠️requirements⚠️
    - register an extension resource -- through -- `APIService`
      - OTHERWISE, the aggregation layer will do NOTHING
      - [MORE](../../../reference/kubernetes-api/cluster-resources/api-service-v1.md)
  - [how to configure](../../../tasks/extend-kubernetes/configure-aggregation-layer)

* 👀ways to implement the `APIService`👀
  * 💡run an extension API server | Pod(s) / run | your cluster💡
    * MOST common way 
    * [how to setup](../../../tasks/extend-kubernetes/setup-extension-api-server)
    * if you use the extension API server -- to -- manage resources | your cluster -> pair the extension API server + >= 1 [controllers](../../../reference/glossary/controller.md)
    * SHOULD have/CLOSE TO REQUIREMENTS
      * low latency networking 
        * to
        * from the kube-apiserver
      * discovery requests (kube-apiserver -- to -- extension-api-server) <= 5 seconds

```title="diagram"
┌─────────────────────────────────────────────────────────────────────┐
│                     kube-apiserver (single process)                 │
│                                                                     │
│  ┌───────────────────────────────────────────────────────────────┐ │
│  │  Core API                                                     │ │
│  │  • Handles: /api/v1/pods, /api/v1/services, etc.            │ │
│  │  • Built-in resources                                        │ │
│  └───────────────────────────────────────────────────────────────┘ │
│                                                                     │
│  ┌───────────────────────────────────────────────────────────────┐ │
│  │  Aggregation Layer                                           │ │
│  │  • Reads APIService objects                                  │ │
│  │  • Routes requests to extension API servers                  │ │
│  │  • Acts as proxy/delegator                                   │ │
│  │                                                               │ │
│  │  Routing Table:                                              │ │
│  │  /apis/metrics.k8s.io/v1beta1/* → metrics-server             │ │
│  │  /apis/custom.io/v1/*           → my-custom-api              │ │
│  └───────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────────┘
                        │                              │
                        │ (delegates via Service)      │
                        ↓                              ↓
         ┌─────────────────────────┐   ┌─────────────────────────┐
         │  Extension API Server   │   │  Extension API Server   │
         │  (metrics-server)       │   │  (my-custom-api)        │
         │  • Out-of-process       │   │  • Out-of-process       │
         │  • Runs in Pod(s)       │   │  • Runs in Pod(s)       │
         │  • Your custom code     │   │  • Your custom code     │
         └─────────────────────────┘   └─────────────────────────┘
```

```title="requestFlow"
Client (kubectl/app)
    │
    │ GET /apis/metrics.k8s.io/v1beta1/nodes
    ↓
┌───────────────────────────────────┐
│  kube-apiserver                   │
│  1. Receives request              │
│  2. Aggregation Layer checks:     │
│     • Path matches APIService?    │
│     • Yes → Proxy to target       │
└───────────────────────────────────┘
    │
    │ Forwards request (acts as proxy)
    ↓
┌───────────────────────────────────┐
│  Extension API Server             │
│  (metrics-server)                 │
│  3. Processes request             │
│  4. Returns response              │
└───────────────────────────────────┘
    │
    │ Response flows back
    ↓
Client receives data
```

* [apiserver-builder](https://github.com/kubernetes-sigs/apiserver-builder-alpha)
  * == library
    * provides
      * skeleton -- for --
        * extension API servers
        * associated controller(s)
