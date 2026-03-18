---
title: Custom Resources
reviewers:
- enisoc
- deads2k
api_metadata:
- apiVersion: "apiextensions.k8s.io/v1"
  kind: "CustomResourceDefinition"
content_type: concept
weight: 10
---

* goal
  * custom resources (CR)
  * how to choose BETWEEN CR vs API Aggregation vs standalone service?

## Custom resources

* [Resource OR Kubernetes API Resource](../../../reference/glossary/api-resource.md)

* Custom Resource (CR)
  * == extension -- of the -- Kubernetes API /
    * ❌NOT necessary AVAILABLE | default Kubernetes installation❌
    * you can reuse [Kubernetes API features](../../overview/kubernetes-api.md)
  * == customization -- of a -- particular Kubernetes installation
  * use cases
    * ⚠️EVEN core Kubernetes functions⚠️
      * == built -- via -- CR
      * -> Kubernetes MORE modular
  * ❌NOT use cases❌
    * data storage 
      * -- about --
        * application data
        * end user data
        * monitoring data
      * Reason: 🧠tight coupled design
        * != [cloud native application architecture principle: loose coupling BETWEEN components](https://www.cncf.io/about/faq/#what-is-cloud-native)🧠
  * 's storage
    * served & handled -- by the -- Kubernetes control plane
  * | running cluster,
    * can be added/removed/updated -- due to --
      * dynamic registration
      * MANUAL cluster admin
  * 's lifecycle != built-in Kubernetes objects
  * != Kubernetes built-in resources
    * _Examples:_ pods

### ways to access -- to -- a custom resource

* ways to access -- to --
  * [Kubernetes API](../../../reference/glossary/kubernetes-api.md)
    * ❌!= access -- to -- custom resource❌
  * custom resource
    * [`kubectl`](../../../reference/kubectl)
    * REST calls
    * [Kubernetes dynamic client](https://github.com/kubernetes/client-go/tree/master/dynamic)
    * Kubernetes client libraries
      * ONLY supported by 
        * [Go](https://github.com/kubernetes/client-go/)
        * [Python](https://github.com/kubernetes-client/python/)
    * client /
      * you create -- based on -- [Kubernetes CodeGenerator](https://github.com/kubernetes/code-generator)

### ways to add custom resources

* CRD vs AA
  * ABOUT complexity,
    * CRDs are easier -- than -- Aggregated APIs

| Characteristic                       | CRD                                                                                                          | Aggregated API                                                            |
|--------------------------------------|--------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------|
| Requires programming                 | ❌NO❌ <br/>  &nbsp;&nbsp; freely to choose the language -- for your -- CRD controller                         | Yes <br/>  &nbsp;&nbsp; programming + building binary + image             |
| Requires ADDITIONAL service          | ❌NO❌ <br/>  &nbsp;&nbsp; Reason: 🧠CRDs are handled -- by -- API server🧠                                    | Yes <br/> &nbsp;&nbsp; create ADDITIONAL EXTERNAL API server              |
| Requires maintenance & bug fixes     | ❌NO❌ <br/> &nbsp;&nbsp; Reason: 🧠part -- of -- normal Kubernetes upgrades🧠                                 | Yes <br/> &nbsp;&nbsp; you need to address them                           |
| Need to handle MULTIPLE API versions | ❌NO❌ <br/> _Example:_ when you control the client for this resource, you can upgrade it in sync with the API | Yes <br/> _Example:_ when developing an extension to share with the world |

* CRD vs AA
  * ABOUT features,
    * Aggregated APIs
      * offer
        * MORE advanced API features
        * MORE flexibility to customize the features

| Feature                      | Description                                                                                                                                                                                                                                                                                                        | CRDs                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Aggregated API                                                                                                              |
|------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------|
| Validation                   | allows <br/> &nbsp;&nbsp; prevent errors <br/> &nbsp;&nbsp; evolve your API INDEPENDENTLY -- of -- your clients <br/> useful <br/> &nbsp;&nbsp; there are MANY clients / can NOT ALL update \| SAME time                                                                                                           | YES <br/> &nbsp;&nbsp; MOST can be specified \| CRD -- through -- [OpenAPI v3.0 validation](../../../tasks/extend-kubernetes/custom-resources/custom-resource-definitions.md#validation) <br/> &nbsp;&nbsp; [CRDValidationRatcheting](../../../tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#validation-ratcheting) == feature gate / if the resource's failing part was unchanged -> ignored  <br/> &nbsp;&nbsp; you can ALSO supported -- by -- adding a [Validating Webhook](../../../reference/access-authn-authz/admission-controllers.md#validatingadmissionwebhook-alpha-in-1-8-beta-in-1-9). | YES <br/> &nbsp;&nbsp; arbitrary validation checks                                                                          |
| Defaulting                   | default fields                                                                                                                                                                                                                                                                                                     | YES <br/> ways <br/> &nbsp;&nbsp; -- via -- `default` keyword \| [OpenAPI v3.0 validation](../../../tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#defaulting) <br/> &nbsp;&nbsp; -- via -- [Mutating Webhook](../../../reference/access-authn-authz/admission-controllers/#mutatingadmissionwebhook) <br/> &nbsp;&nbsp; &nbsp;&nbsp; if you read FROM etcd -- for -- old objects -> it will NOT be run                                                                                                                                                                                               | Yes                                                                                                                         |
| Multi-versioning             | == serve the SAME object -- through -- 2 API versions <br/> allows <br/> &nbsp;&nbsp; making easier API changes (_Example:_ renaming fields) <br/> if you control your client versions -> LESS important                                                                                                           | [Yes](../../../tasks/extend-kubernetes/custom-resources/custom-resource-definition-versioning)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | Yes                                                                                                                         |
| Custom Storage               | use cases <br/> &nbsp;&nbsp; choose DIFFERENT performance mode (_Example:_ time-series database vs key-value store) <br/> &nbsp;&nbsp; isolate (_Example:_ encrypt sensitive information, etc.)                                                                                                                    | No                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Yes                                                                                                                         |
| Custom Business Logic        | \| create/read/update/delete an object, perform arbitrary checks OR actions                                                                                                                                                                                                                                        | Yes -- via -- [Webhooks](../../../reference/access-authn-authz/extensible-admission-controllers.md#admission-webhooks)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Yes                                                                                                                         |
| Scale Subresource            | enable <br/> &nbsp;&nbsp; some systems (_Examples:_ `HorizontalPodAutoscaler` & `PodDisruptionBudget`) can interact -- with -- your NEW resource                                                                                                                                                                   | [Yes](../../../tasks/extend-kubernetes/custom-resources/custom-resource-definitions.md#scale-subresource)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Yes                                                                                                                         |
| Status Subresource           | allows <br/> &nbsp;&nbsp; fine-grained access control / <br/> &nbsp;&nbsp; &nbsp;&nbsp; user writes the spec section <br/> &nbsp;&nbsp; &nbsp;&nbsp; controller writes the status section <br/> &nbsp;&nbsp; if custom resource's `.spec` changes -> incremente object generation (`metadata.generation`)            | [Yes](../../../tasks/extend-kubernetes/custom-resources/custom-resource-definitions.md#status-subresource)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Yes                                                                                                                         |
| Other Subresources           | operations / != CRUD (_Examples:_ "logs" or "exec")                                                                                                                                                                                                                                                                | No                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Yes                                                                                                                         |
| strategic-merge-patch        | == NEW endpoints support PATCH -- via -- `Content-Type: application/strategic-merge-patch+json` <br/> useful for: update objects / may be modified locally & by the server <br/> [update EXISTING API Objects -- via -- `kubectl patch`](../../../tasks/manage-kubernetes-objects/update-api-object-kubectl-patch) | No                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Yes                                                                                                                         |
| Protocol Buffers             | == NEW resource supports clients / want to use Protocol Buffers                                                                                                                                                                                                                                                    | No                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Yes                                                                                                                         |
| OpenAPI Schema               | Is there an OpenAPI (swagger) schema -- for the -- types / can be DYNAMICALLY fetched -- from the -- server? <br/> Is the user protected -- from -- misspelling field names / ensure ONLY ALLOWED fields are set? <br/> Are types enforced ?                                                                       | Yes -- based on -- [OpenAPI v3.0 validation schema](../../../tasks/extend-kubernetes/custom-resources/custom-resource-definitions.md#validation)                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Yes                                                                                                                         |
| Instance Name                | impose any constraints \| kind/resource' objects' name?                                                                                                                                                                                                                                                            | Yes <br/> &nbsp;&nbsp; object's name MUST be a [valid DNS subdomain name](../../overview/working-with-objects/names.md#dns-subdomain-names)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | No                                                                                                                          |
| "kube-apiserver" recognition | kube-apiserver recognizes the NEW custom resources                                                                                                                                                                                                                                                                 | Yes                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | NO <br/> &nbsp;&nbsp; kube-apiserver acts -- , ONLY,  as -- proxy <br/> &nbsp;&nbsp; external "kube-apiserver" recognize it |

#### CustomResourceDefinitions

* [here](../../../reference/glossary/customresourcedefinition.md)
* [MORE](../../../tasks/extend-kubernetes/custom-resources/custom-resource-definitions)

* use cases
  * small number of small objects
    * Reason:🧠they are stored | etcd
      * != DDBB🧠
  * use the resource | 
    * your company OR
    * part of a small open-source project

* _Example:_ [define a CRD + controller / handle events](https://github.com/dancer1325/sample-controller)

#### API server aggregation

* Reason why it adds custom resources: 🧠ALSO extend the Kubernetes API🧠
* [here](apiserver-aggregation.md)

## Custom controllers

* Custom controllers
  * == [controllers](../../../reference/glossary/controller.md) / 
    * wrap the domain knowledge
    * 's lifecycle
      * INDEPENDENT -- of -- cluster's lifecycle
  * use cases
    * 👀| ANY kind of resource👀
      * == ❌NOT ONLY custom resource❌
      * \+ custom resource

TODO:

## Prerequirements -- to -- install a custom resource

### Third party code and new points of failure

While creating a CRD does not automatically add any new points of failure (for example, by causing
third party code to run on your API server), packages (for example, Charts) or other installation
bundles often include CRDs as well as a Deployment of third-party code that implements the
business logic for a new custom resource.

Installing an Aggregated API server always involves running a new Deployment.

### Storage

Custom resources consume storage space in the same way that ConfigMaps do. Creating too many
custom resources may overload your API server's storage space.

Custom resources are placed into storage based upon the the current storage
version of the resource, defined in the CRD spec. Any update to a custom
resource will use the currently defined storage version to store the resource.
All other versions either need to have all the fields of that version or define
conversions to work properly.

Aggregated API servers may use the same storage as the main API server, in which case the same
warning applies.

### Authentication, authorization, and auditing

CRDs always use the same authentication, authorization, and audit logging as the built-in
resources of your API server.

If you use RBAC for authorization, most RBAC roles will not grant access to the new resources
(except the cluster-admin role or any role created with wildcard rules). You'll need to explicitly
grant access to the new resources. CRDs and Aggregated APIs often come bundled with new role
definitions for the types they add.

Aggregated API servers may or may not use the same authentication, authorization, and auditing as
the primary API server.


## Custom resource field selectors

[Field Selectors](/docs/concepts/overview/working-with-objects/field-selectors/)
let clients select custom resources based on the value of one or more resource
fields.

All custom resources support the `metadata.name` and `metadata.namespace` field
selectors.

Fields declared in a {{< glossary_tooltip term_id="CustomResourceDefinition" text="CustomResourceDefinition" >}}
may also be used with field selectors when included in the `spec.versions[*].selectableFields` field of the
{{< glossary_tooltip term_id="CustomResourceDefinition" text="CustomResourceDefinition" >}}.

### Selectable fields for custom resources {#crd-selectable-fields}

{{< feature-state feature_gate_name="CustomResourceFieldSelectors" >}}

The `spec.versions[*].selectableFields` field of a {{< glossary_tooltip term_id="CustomResourceDefinition" text="CustomResourceDefinition" >}} may be used to
declare which other fields in a custom resource may be used in field selectors.

The following example adds the `.spec.color` and `.spec.size` fields as
selectable fields.

{{% code_sample file="customresourcedefinition/shirt-resource-definition.yaml" %}}

Field selectors can then be used to get only resources with a `color` of `blue`:

```shell
kubectl get shirts.stable.example.com --field-selector spec.color=blue
```

The output should be:

```
NAME       COLOR  SIZE
example1   blue   S
example2   blue   M
```
