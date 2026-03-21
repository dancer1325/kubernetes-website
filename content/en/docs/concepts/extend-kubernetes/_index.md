---
title: Extending Kubernetes
weight: 999 # this section should come last
description: Different ways to change the behavior of your Kubernetes cluster.
reviewers:
- erictune
- lavalamp
- cheftako
- chenopis
feature:
  title: Designed for extensibility
  description: >
    Add features to your Kubernetes cluster without changing upstream source code.
content_type: concept
no_list: true
---

* goal
  * how to customize a Kubernetes cluster

* audience
  * [cluster operator](../../reference/glossary/cluster-operator.md)
  * [platform developer](../../reference/glossary/platform-developer.md)
  * [Kubernetes contributors](../../reference/glossary/contributor.md)

* ways to customize Kubernetes
  * [configuration](#configuration)
  * [policy APIs](#policy-apis)
  * [extensions](#extensions)
    * == run ADDITIONAL programs &/OR ADDITIONAL network services

## Configuration

* == change Kubernetes component configuration
  * ways (== changing)
    * CL arguments
    * local configuration files
    * [API resources](../../reference/glossary/api-resource.md)
  * ALLOWED | the Kubernetes components
    * [`kube-apiserver`](../../reference/command-line-tools-reference/kube-apiserver)
    * [`kube-controller-manager`](../../reference/command-line-tools-reference/kube-controller-manager)
    * [`kube-scheduler`](../../reference/command-line-tools-reference/kube-scheduler)
    * [`kubelet`](../../reference/command-line-tools-reference/kubelet)
    * [`kube-proxy`](../../reference/command-line-tools-reference/kube-proxy)
  * cons
    * NORMALLY, ONLY ALLOWED -- by -- [cluster operator](../../reference/glossary/cluster-operator.md)
    * Kubernetes-version dependant
      * == if there are modifications BETWEEN Kubernetes versions -> you need to modify & restart processes

## Policy APIs

* built-in Policy APIs
  * == built-in Kubernetes APIs /
    * provide
      * declaratively configured policy settings
  * _Examples:_ 
    * [ResourceQuota](../policy/resource-quotas)
    * [NetworkPolicy](../services-networking/network-policies)
    * [RBAC](../../reference/access-authn-authz/rbac)
  * use cases | they can be managed
    * hosted Kubernetes services
    * managed Kubernetes installations
  * 's conventions / follow == OTHER Kubernetes resources' conventions / follow

* recommendation
  * use policy APIs -- RATHER THAN -- [configuration](#configuration)
    * Reason:🧠if you use a policy APIs / [stable](../../reference/using-api/_index.md#api-versioning) -> you benefit from a [defined support policy](../../reference/using-api/deprecation-policy.md)
      * == -- as -- other Kubernetes APIs🧠

## Extensions

* Extensions
  * == software components / 
    * extend Kubernetes -- to -- support
      * NEW types
      * NEW kinds of hardware
    * deeply integrate -- with -- Kubernetes
  * audience
    * Kubernetes cluster administrators /
      * ❌NO use❌
        * hosted Kubernetes
          * _Examples:_ GKE, EKS, AKS
        * instance of Kubernetes
          * _Examples:_ OpenShift, Rancher

* Kubernetes' design -- about -- automation:
  * 💡-- by -- writing client programs💡
    * recommendation
      * use controller pattern
    * -- through -- Kubernetes API
    * / can run
      * | cluster
      * outside cluster

### Extension patterns

* controller pattern
  * == pattern -- for -- writing client programs
  * [controller](../../reference/glossary/controller.md)

### Integration Mechanism

#### webhook

* webhook
  * :=
    * (| Kubernetes ecosystem),
      * mechanism -- for -- (⚠️async & sync⚠️) notifications /
        * "kube-apiserver" calls -- to a -- remote service (webhook backend)
    * (| outside Kubernetes),
      * mechanism -- for -- asynchronous notifications /
        * webhook call == 1-way notification -- to -- ANOTHER system or component
  * webhook backend
    * == remote service

  ┌──────────────────┐             ┌─────────────────┐
  │  kube-apiserver  │  ────────>  │ Webhook Backend │
  │  (client)        │  HTTP req   │  (server)       │
  └──────────────────┘             └─────────────────┘

* webhook vs controller
  * BOTH add a point of failure 
    * ⚠️ALTHOUGH webhook failures are MORE critic⚠️
      * Reason:🧠they can block cluster operations🧠

#### binary Plugin model

* binary Plugin model
  * == Kubernetes executes a binary (program)
  * uses
    * by [kubelet](../../tasks/extend-kubectl/kubectl-plugins.md)
      * _Examples:_
        * [CSI storage plugins](https://kubernetes-csi.github.io/docs/)
        * [CNI network plugins](compute-storage-net/network-plugins.md)

### Extension points

![extension points | Kubernetes cluster + clients / access it](/content/en/docs/concepts/extend-kubernetes/extension-points.png)

TODO: 
#### Key to the figure

1. [Client extensions](#client-extensions)
2. [API Access Extensions](#api-access-extensions)
3. [API extensions](#api-extensions)
4. [Scheduling extensions](#scheduling-extensions)
5. Create custom [Controllers](../../reference/glossary/controller.md) 
   * [combining new APIs with automation](#new-apis--automation)
   * [Changing built-in resources](#changing-built-in-resources)
6. The kubelet runs on servers (nodes), and helps pods appear like virtual servers with their own IPs on
   the cluster network. [Network Plugins](#network-plugins) allow for different implementations of
   pod networking.
7. [Device Plugins](#device-plugins)

   The kubelet also mounts and unmounts
   {{< glossary_tooltip text="volume" term_id="volume" >}} for pods and their containers.
   You can use [Storage Plugins](#storage-plugins) to add support for new kinds
   of storage and other volume types.


#### Extension point choice flowchart {#extension-flowchart}

![how to choose the extension point?](/content/en/docs/concepts/extend-kubernetes/flowchart.svg)

## Client extensions

* plugins
  * == extensions
  * allow
    * customise the behaviour of clients
  * types
    * generic 
      * == apply | DIFFERENT clients
      * _Example:_ [credential plugins](../../reference/access-authn-authz/authentication.md#client-go-credential-plugins)
    * [specific ways to extend `kubectl`](../../tasks/extend-kubectl/kubectl-plugins)

## API extensions

### Custom resource definitions (CRD)

* [here](api-extension/custom-resources.md)

### API aggregation layer (AA)

* [here](api-extension/apiserver-aggregation.md)

### NEW APIs + Automation

TODO: how is it related with API extension? or NOT equivalent to CRD?

* operator-pattern
  * == pattern / controller deploys infrastructure -- based on a -- desired state
  * uses
    * manage specific applications
      * _Example:_ applications / 
        * maintain state
        * | manage, require care

* define policies
  * _Example:_ access control restriction

### Changing built-in resources

* API access extensions
  * allow
    * changing the behavior of EXISTING APIs 
      * _Example:_ pods

* Kubernetes API extensions -- by -- adding custom resources
  * add NEW API Groups' resources
    * ❌!= replace OR change EXISTING API groups❌

## API access extensions

API server handles all requests
* Several types of extension points in the API server allow
  authenticating requests, or blocking them based on their content, editing content, and handling
  deletion

* When a request reaches the Kubernetes API Server, it is first _authenticated_, then _authorized_,
and is then subject to various types of _admission control_ (some requests are in fact not
authenticated, and get special treatment)
* [Controlling Access to the Kubernetes API](../../concepts/security/controlling-access.md)

Each of the steps in the Kubernetes authentication / authorization flow offers extension points.

### Authentication

[Authentication](/docs/reference/access-authn-authz/authentication/) maps headers or certificates
in all requests to a username for the client making the request.

Kubernetes has several built-in authentication methods that it supports. It can also sit behind an
authenticating proxy, and it can send a token from an `Authorization:` header to a remote service for
verification (an [authentication webhook](/docs/reference/access-authn-authz/authentication/#webhook-token-authentication))
if those don't meet your needs.

### Authorization

[Authorization](/docs/reference/access-authn-authz/authorization/) determines whether specific
users can read, write, and do other operations on API resources. It works at the level of whole
resources -- it doesn't discriminate based on arbitrary object fields.

If the built-in authorization options don't meet your needs, an
[authorization webhook](/docs/reference/access-authn-authz/webhook/)
allows calling out to custom code that makes an authorization decision.

### Dynamic admission control

After a request is authorized, if it is a write operation, it also goes through
[Admission Control](/docs/reference/access-authn-authz/admission-controllers/) steps.
In addition to the built-in steps, there are several extensions:

* The [Image Policy webhook](/docs/reference/access-authn-authz/admission-controllers/#imagepolicywebhook)
  restricts what images can be run in containers.
* To make arbitrary admission control decisions, a general
  [Admission webhook](/docs/reference/access-authn-authz/extensible-admission-controllers/#admission-webhooks)
  can be used. Admission webhooks can reject creations or updates.
  Some admission webhooks modify the incoming request data before it is handled further by Kubernetes.

## Infrastructure extensions

### Device plugins

* Device plugins
  * allow
    * node can discover NEW node resources
      * _Example of EXISTING resources:_ cpu and memory
  * uses
    * integrate custom hardware OR other special node-local facilities /
      * AVAILABLE | pods / run | your custer
  * [MORE](compute-storage-net/device-plugins.md)

### Storage plugins

{{< glossary_tooltip text="Container Storage Interface" term_id="csi" >}} (CSI) plugins provide
a way to extend Kubernetes with supports for new kinds of volumes. The volumes can be backed by
durable external storage, or provide ephemeral storage, or they might offer a read-only interface
to information using a filesystem paradigm.

Kubernetes also includes support for [FlexVolume](/docs/concepts/storage/volumes/#flexvolume) plugins,
which are deprecated since Kubernetes v1.23 (in favour of CSI).

FlexVolume plugins allow users to mount volume types that aren't natively supported by Kubernetes. When
you run a Pod that relies on FlexVolume storage, the kubelet calls a binary plugin to mount the volume.
The archived [FlexVolume](https://git.k8s.io/design-proposals-archive/storage/flexvolume-deployment.md)
design proposal has more detail on this approach.

The [Kubernetes Volume Plugin FAQ for Storage Vendors](https://github.com/kubernetes/community/blob/master/sig-storage/volume-plugin-faq.md#kubernetes-volume-plugin-faq-for-storage-vendors)
includes general information on storage plugins.

### Network plugins

Your Kubernetes cluster needs a _network plugin_ in order to have a working Pod network
and to support other aspects of the Kubernetes network model.

[Network Plugins](/docs/concepts/extend-kubernetes/compute-storage-net/network-plugins/)
allow Kubernetes to work with different networking topologies and technologies.

### Kubelet image credential provider plugins

{{< feature-state for_k8s_version="v1.26" state="stable" >}}
Kubelet image credential providers are plugins for the kubelet to dynamically retrieve image registry
credentials. The credentials are then used when pulling images from container image registries that
match the configuration.

The plugins can communicate with external services or use local files to obtain credentials. This way,
the kubelet does not need to have static credentials for each registry, and can support various
authentication methods and protocols.

For plugin configuration details, see
[Configure a kubelet image credential provider](/docs/tasks/administer-cluster/kubelet-credential-provider/).

## Scheduling extensions

The Kubernetes scheduler [decides](/docs/concepts/scheduling-eviction/assign-pod-node/)
which nodes to place pods on

- [Scheduling](Scheduling,%20Preemption%20and%20Eviction%2011a13fe1aaf54bcfa09ecf21179846c1.md) extensions
  - [Schedulers](Scheduling,%20Preemption%20and%20Eviction/Kubernetes%20Scheduler%2028b495fc09a54694a371e5c85524eb1e.md)
  - == webhook which permits a remote HTTP backend
    **Note:** 👁️ Check [textProposal](https://github.com/kubernetes/design-proposals-archive/blob/main/scheduling/scheduler_extender.md) 👁️
  - available configurations
    - [Scheduling plugins](../Reference/Scheduling/Scheduler%20Configuration%207400516594184fb4b4ff406c9642a662.md)
    - [Scheduling Profile](../Reference/Scheduling/Scheduler%20Configuration%207400516594184fb4b4ff406c9642a662.md)

The scheduler is a special type of controller that watches pods, and assigns
pods to nodes. The default scheduler can be replaced entirely, while
continuing to use other Kubernetes components, or
[multiple schedulers](/docs/tasks/extend-kubernetes/configure-multiple-schedulers/)
can run at the same time.

This is a significant undertaking, and almost all Kubernetes users find they
do not need to modify the scheduler.

You can control which [scheduling plugins](/docs/reference/scheduling/config/#scheduling-plugins)
are active, or associate sets of plugins with different named [scheduler profiles](/docs/reference/scheduling/config/#multiple-profiles).
You can also write your own plugin that integrates with one or more of the kube-scheduler's
[extension points](/docs/concepts/scheduling-eviction/scheduling-framework/#extension-points).

Finally, the built in `kube-scheduler` component supports a
[webhook](https://git.k8s.io/design-proposals-archive/scheduling/scheduler_extender.md)
that permits a remote HTTP backend (scheduler extension) to filter and / or prioritize
the nodes that the kube-scheduler chooses for a pod.

{{< note >}}
You can only affect node filtering
and node prioritization with a scheduler extender webhook; other extension points are
not available through the webhook integration.
{{< /note >}}

## {{% heading "whatsnext" %}}

* Learn more about infrastructure extensions
  * [Device Plugins](/docs/concepts/extend-kubernetes/compute-storage-net/device-plugins/)
  * [Network Plugins](/docs/concepts/extend-kubernetes/compute-storage-net/network-plugins/)
  * CSI [storage plugins](https://kubernetes-csi.github.io/docs/)
* Learn about [kubectl plugins](/docs/tasks/extend-kubectl/kubectl-plugins/)
* Learn more about [Custom Resources](/docs/concepts/extend-kubernetes/api-extension/custom-resources/)
* Learn more about [Extension API Servers](/docs/concepts/extend-kubernetes/api-extension/apiserver-aggregation/)
* Learn about [Dynamic admission control](/docs/reference/access-authn-authz/extensible-admission-controllers/)
* Learn about the [Operator pattern](/docs/concepts/extend-kubernetes/operator/)
