* Use configMap to provide configurations for system components

# How to run it locally?
* [Install GO locally](https://go.dev/dl/)
* `go get k8s.io/client-go@v0.21.0`
  * Problems
    * Problem1: "no such file or directory"
      * Attempt1: Run outside this directory -> error to run `go get` outside modules
      * Attempt2: If you are running on Mac M1 -> crete and change manually "/mnt/" -> "/Volumes/mnt/"
      * Solution: `go mod init` & `go get k8s.io/client-go@v0.21.0` to add via module
* `go build -o my-controller`
  * Compile with the Go compiler
  * Problems:
    * Problem1:  undefined: metav1
      * Attempt1: `import "k8s.io/apimachinery/pkg/apis/meta/v1" `
* `./my-controller`
  * Run the app locally
* Set `KUBECONFIG` to point the Kubernetes configuration file
* `kubectl apply -f configmap.yaml`
  * Deploy the configMap in the cluster
  