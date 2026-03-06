A simple example of an object created using Server-Side Apply could look like this:

{{< note >}}
`kubectl get` omits managed fields by default.
Add `--show-managed-fields` to show `managedFields` when the output format is either `json` or `yaml`.
{{< /note >}}

```yaml
---

```

That example ConfigMap object contains a single field management record in
`.metadata.managedFields`. The field management record consists of basic information
about the managing entity itself, plus details about the fields being managed and
the relevant operation (`Apply` or `Update`). If the request that last changed that
field was a Server-Side Apply **patch** then the value of `operation` is `Apply`;
otherwise, it is `Update`.

There is another possible outcome. A client could submit an invalid request
body. If the fully specified intent does not produce a valid object, the
request fails.

It is however possible to change `.metadata.managedFields` through an
**update**, or through a **patch** operation that does not use Server-Side Apply.
Doing so is highly discouraged, but might be a reasonable option to try if,
for example, the `.metadata.managedFields` get into an inconsistent state
(which should not happen in normal operations).

The format of `managedFields` is [described](/docs/reference/kubernetes-api/common-definitions/object-meta/#System)
in the Kubernetes API reference.

{{< caution >}}
The `.metadata.managedFields` field is managed by the API server.
You should avoid updating it manually.
{{< /caution >}}