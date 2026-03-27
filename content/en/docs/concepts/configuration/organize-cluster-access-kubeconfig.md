---
title: Organizing Cluster Access Using kubeconfig Files
content_type: concept
weight: 60
---

* kubeconfig file
  * == file / 
    * configure access -- to -- clusters
  * allows
    * organize information -- about -- clusters + users + namespaces + authentication mechanisms
  * uses
    * `kubectl` uses it -- to --
      * choose a cluster
      * communicate -- with the -- cluster's API server
  * recommendations
    * ⚠️ONLY use those / come -- from -- trusted sources⚠️
      * Reason: 🧠if it's malicious -> it could execute code OR expose file🧠
  * support
    * [MULTIPLE authentication mechanisms](../../tasks/access-application-cluster/configure-access-multiple-clusters)
      * _ExampleS:_ certificates, tokens
  * ['s context](#context)

* `kubectl`
  * 👀looks for the kubeconfig file |👀 
    * `$HOME/.kube/config`
      * default one
    * [`KUBECONFIG` environment variable](#kubeconfig-environment-variable)
      * OPTIONAL
        * == if you specified
    * [`--kubeconfig`](../../reference/kubectl/kubectl)👀
      * OPTIONAL
        * == you specify the CL flag
  * ⚠️`--kubeconfig`'s priority > `KUBECONFIG`'s priority > `$HOME/.kube/config`'s priority⚠️ 

## Context

* context
  * == cluster + namespace + user
  * uses
    * group access parameters / convenient name

* current context
  * uses
    * by `kubectl`, to communicate -- with the -- cluster
  * if you want to choose -> use `kubectl config use-context`

## `KUBECONFIG` environment variable

* `KUBECONFIG` environment variable
  * == list of kubeconfig fileS pathS /
    * are [merged](#merging-kubeconfig-files)
    * -- based on -- OS
      * | Linux & Mac, 
        * `:`-delimited
          * == `path1:path2:...`
      * | Windows,
        * `;`-delimited
          * == `path1,path2;...`
  * OPTIONAL
    * != MANDATORY

## Merging kubeconfig files

* merge kubeconfig fileS
  * 💡happen | ANY `kubectl ANYCOMMAND`💡
  * rules
    1. if you set `--kubeconfig` CL flag -> ONLY use 1 (the specified one) kubeconfig file
       * == ❌NOT merge❌
       * Reason:🧠| `kubectl`, you can ONLY pass 1! `--kubeconfig`🧠
       * TODO: 

          Otherwise, if the `KUBECONFIG` environment variable is set, use it as a
          list of files that should be merged.
          Merge the files listed in the `KUBECONFIG` environment variable
          according to these rules:

          * Ignore empty filenames.
          * Produce errors for files with content that cannot be deserialized.
          * The first file to set a particular value or map key wins.
          * Never change the value or map key.
            Example: Preserve the context of the first file to set `current-context`.
            Example: If two files specify a `red-user`, use only values from the first file's `red-user`.
            Even if the second file has non-conflicting entries under `red-user`, discard them.

   For an example of setting the `KUBECONFIG` environment variable, see
   [Setting the KUBECONFIG environment variable](/docs/tasks/access-application-cluster/configure-access-multiple-clusters/#set-the-kubeconfig-environment-variable).

   Otherwise, use the default kubeconfig file, `$HOME/.kube/config`, with no merging.

  1. Determine the context to use based on the first hit in this chain:

     1. Use the `--context` command-line flag if it exists.
     1. Use the `current-context` from the merged kubeconfig files.

   An empty context is allowed at this point.

  1. Determine the cluster and user. At this point, there might or might not be a context.
     Determine the cluster and user based on the first hit in this chain,
     which is run twice: once for user and once for cluster:

     1. Use a command-line flag if it exists: `--user` or `--cluster`.
     1. If the context is non-empty, take the user or cluster from the context.

   The user and cluster can be empty at this point.

  1. Determine the actual cluster information to use. At this point, there might or
     might not be cluster information.
     Build each piece of the cluster information based on this chain; the first hit wins:

     1. Use command line flags if they exist: `--server`, `--certificate-authority`, `--insecure-skip-tls-verify`.
     1. If any cluster information attributes exist from the merged kubeconfig files, use them.
     1. If there is no server location, fail.

  1. Determine the actual user information to use. Build user information using the same
     rules as cluster information, except allow only one authentication
     technique per user:

     1. Use command line flags if they exist: `--client-certificate`, `--client-key`, `--username`, `--password`, `--token`.
     1. Use the `user` fields from the merged kubeconfig files.
     1. If there are two conflicting techniques, fail.

  1. For any information still missing, use default values and potentially
     prompt for authentication information.

* `kubectl config view`
  * display merged kubeconfig fileS

## File references

File and path references in a kubeconfig file are relative to the location of the kubeconfig file.
File references on the command line are relative to the current working directory.
In `$HOME/.kube/config`, relative paths are stored relatively, and absolute paths
are stored absolutely.

## Proxy

* goal
  * specify a proxy / cluster

* steps
  * | your kubeconfig file,
    * specify `clusters[SOME].proxy-url`
  You can configure `kubectl` to use a proxy per cluster using `proxy-url` in 

## {{% heading "whatsnext" %}}

* [`kubectl config`](/docs/reference/generated/kubectl/kubectl-commands#config)
