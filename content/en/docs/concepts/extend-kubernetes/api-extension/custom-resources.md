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
    * [CRD](#customresourcedefinitions)
    * [AA](apiserver-aggregation.md)
      * Reason: 🧠ALSO extend the Kubernetes API🧠
  * 's storage
    * served & handled -- by the -- Kubernetes control plane
  * | running cluster,
    * can be added/removed/updated -- due to --
      * dynamic registration
      * MANUAL cluster admin
  * 's lifecycle != built-in Kubernetes objects

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

## Adding custom resources

Kubernetes provides two ways to add custom resources to your cluster:

- CRDs are simple and can be created without any programming.
- [API Aggregation](/docs/concepts/extend-kubernetes/api-extension/apiserver-aggregation/)
  requires programming, but allows more control over API behaviors like how data is stored and
  conversion between API versions.

Kubernetes provides these two options to meet the needs of different users, so that neither ease
of use nor flexibility is compromised.

Aggregated APIs are subordinate API servers that sit behind the primary API server, which acts as
a proxy. This arrangement is called [API Aggregation](/docs/concepts/extend-kubernetes/api-extension/apiserver-aggregation/)(AA).
To users, the Kubernetes API appears extended.

CRDs allow users to create new types of resources without adding another API server. You do not
need to understand API Aggregation to use CRDs.

Regardless of how they are installed, the new resources are referred to as Custom Resources to
distinguish them from built-in Kubernetes resources (like pods).

{{< note >}}
Avoid using a Custom Resource as data storage for application, end user, or monitoring data:
architecture designs that store application data within the Kubernetes API typically represent
a design that is too closely coupled.

Architecturally, [cloud native](https://www.cncf.io/about/faq/#what-is-cloud-native) application architectures
favor loose coupling between components. If part of your workload requires a backing service for
its routine operation, run that backing service as a component or consume it as an external service.
This way, your workload does not rely on the Kubernetes API for its normal operation.
{{< /note >}}

## CustomResourceDefinitions

* [here](../../../reference/glossary/customresourcedefinition.md)
The [CustomResourceDefinition](/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/)
API resource allows you to define custom resources.
Defining a CRD object creates a new custom resource with a name and schema that you specify.
The Kubernetes API serves and handles the storage of your custom resource.
The name of the CRD object itself must be a valid
[DNS subdomain name](/docs/concepts/overview/working-with-objects/names#dns-subdomain-names) derived from the defined resource name and its API group; see [how to create a CRD](/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions#create-a-customresourcedefinition) for more details.
Further, the name of an object whose kind/resource is defined by a CRD must also be a valid DNS subdomain name.

This frees you from writing your own API server to handle the custom resource,
but the generic nature of the implementation means you have less flexibility than with
[API server aggregation](#api-server-aggregation).

Refer to the [custom controller example](https://github.com/kubernetes/sample-controller)
for an example of how to register a new custom resource, work with instances of your new resource type,
and use a controller to handle events.

## API server aggregation

Usually, each resource in the Kubernetes API requires code that handles REST requests and manages
persistent storage of objects. The main Kubernetes API server handles built-in resources like
*pods* and *services*, and can also generically handle custom resources through
[CRDs](#customresourcedefinitions).

The [aggregation layer](/docs/concepts/extend-kubernetes/api-extension/apiserver-aggregation/)
allows you to provide specialized implementations for your custom resources by writing and
deploying your own API server.
The main API server delegates requests to your API server for the custom resources that you handle,
making them available to all of its clients.

## Choosing a method for adding custom resources

CRDs are easier to use. Aggregated APIs are more flexible. Choose the method that best meets your needs.

Typically, CRDs are a good fit if:

* You have a handful of fields
* You are using the resource within your company, or as part of a small open-source project (as
  opposed to a commercial product)

### Comparing ease of use

CRDs are easier to create than Aggregated APIs.

| CRDs                                                                                                                                                  | Aggregated API                                                                                                       |
|-------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------|
| Do not require programming. Users can choose any language for a CRD controller.                                                                       | Requires programming and building binary and image.                                                                  |
| No additional service to run; CRDs are handled by API server.                                                                                         | An additional service to create and that could fail.                                                                 |
| No ongoing support once the CRD is created. Any bug fixes are picked up as part of normal Kubernetes Master upgrades.                                 | May need to periodically pickup bug fixes from upstream and rebuild and update the Aggregated API server.            |
| No need to handle multiple versions of your API; for example, when you control the client for this resource, you can upgrade it in sync with the API. | You need to handle multiple versions of your API; for example, when developing an extension to share with the world. |

### Advanced features and flexibility

Aggregated APIs offer more advanced API features and customization of other features; for example, the storage layer.

| Feature               | Description                                                                                                                                                                                                                                                                                                                          | CRDs                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Aggregated API |
|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------| -------------- |
| Validation            | Help users prevent errors and allow you to evolve your API independently of your clients. These features are most useful when there are many clients who can't all update at the same time.                                                                                                                                          | Yes.  Most validation can be specified in the CRD using [OpenAPI v3.0 validation](/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#validation). [CRDValidationRatcheting](/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#validation-ratcheting) feature gate allows failing validations specified using OpenAPI also can be ignored if the failing part of the resource was unchanged.  Any other validations supported by addition of a [Validating Webhook](/docs/reference/access-authn-authz/admission-controllers/#validatingadmissionwebhook-alpha-in-1-8-beta-in-1-9). | Yes, arbitrary validation checks |
| Defaulting            | See above                                                                                                                                                                                                                                                                                                                            | Yes, either via [OpenAPI v3.0 validation](/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#defaulting) `default` keyword (GA in 1.17), or via a [Mutating Webhook](/docs/reference/access-authn-authz/admission-controllers/#mutatingadmissionwebhook) (though this will not be run when reading from etcd for old objects).                                                                                                                                                                                                                                                                               | Yes |
| Multi-versioning      | Allows serving the same object through two API versions. Can help ease API changes like renaming fields. Less important if you control your client versions.                                                                                                                                                                         | [Yes](/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definition-versioning)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Yes |
| Custom Storage        | If you need storage with a different performance mode (for example, a time-series database instead of key-value store) or isolation for security (for example, encryption of sensitive information, etc.)                                                                                                                            | No                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Yes |
| Custom Business Logic | Perform arbitrary checks or actions when creating, reading, updating or deleting an object                                                                                                                                                                                                                                           | Yes, using [Webhooks](/docs/reference/access-authn-authz/extensible-admission-controllers/#admission-webhooks).                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Yes |
| Scale Subresource     | Allows systems like HorizontalPodAutoscaler and PodDisruptionBudget interact with your new resource                                                                                                                                                                                                                                  | [Yes](/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#scale-subresource)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Yes |
| Status Subresource    | Allows fine-grained access control where user writes the spec section and the controller writes the status section. Allows incrementing object Generation on custom resource data mutation (requires separate spec and status sections in the resource)                                                                              | [Yes](/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#status-subresource)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | Yes |
| Other Subresources    | Add operations other than CRUD, such as "logs" or "exec".                                                                                                                                                                                                                                                                            | No                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Yes |
| strategic-merge-patch | The new endpoints support PATCH with `Content-Type: application/strategic-merge-patch+json`. Useful for updating objects that may be modified both locally, and by the server. For more information, see ["Update API Objects in Place Using kubectl patch"](/docs/tasks/manage-kubernetes-objects/update-api-object-kubectl-patch/) | No                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Yes |
| Protocol Buffers      | The new resource supports clients that want to use Protocol Buffers                                                                                                                                                                                                                                                                  | No                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Yes |
| OpenAPI Schema        | Is there an OpenAPI (swagger) schema for the types that can be dynamically fetched from the server? Is the user protected from misspelling field names by ensuring only allowed fields are set? Are types enforced (in other words, don't put an `int` in a `string` field?)                                                         | Yes, based on the [OpenAPI v3.0 validation](/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#validation) schema (GA in 1.16).                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Yes |
| Instance Name         | Does this extension mechanism impose any constraints on the names of objects whose kind/resource is defined this way?                                                                                                                                                                                                                | Yes, such an object's name must be a valid DNS subdomain name.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | No |

## Preparing to install a custom resource

There are several points to be aware of before adding a custom resource to your cluster.

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

## {{% heading "whatsnext" %}}

* Learn how to [Extend the Kubernetes API with the aggregation layer](/docs/concepts/extend-kubernetes/api-extension/apiserver-aggregation/).
* Learn how to [Extend the Kubernetes API with CustomResourceDefinition](/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/).

