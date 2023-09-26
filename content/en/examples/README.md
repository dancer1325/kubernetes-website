## How to run locally?
* `go test k8s.io/website/content/<lang>/examples`
  * Run the tests for a localization
  * `go test k8s.io/website/content/en/examples`
    * Run the examples placed under en folder

## Notes:
* Problems:
  * Problem1: "k8s.io/kubernetes/pkg/api/legacyscheme" does not exist
    * Attempt1: `go get`
