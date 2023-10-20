# Content
* 'kustomization.yaml'
  * common file name to use as generator
* from
   * files
     * using default key
       * Go to '/configMapGenerator/FromFilesDefaultKey/'
       * suffix is added to the specified name
       * Problems:
         * Problem1: "error: loading KV pairs: is not in or below "
           * Note: [It's not valid by Kustomize](https://stackoverflow.com/questions/67481924/referring-a-resource-yaml-from-another-directory-in-kustomization)
           * Solution: Create files under it
     * defining the key
       * Go to '/configMapGenerator/FromFilesCustomKey/'
       * suffix is added to the specified name
       * Problems:
         * Problem1: "error: loading KV pairs: is not in or below "
           * Note: [It's not valid by Kustomize](https://stackoverflow.com/questions/67481924/referring-a-resource-yaml-from-another-directory-in-kustomization)
           * Solution: Create files under it
   * literals
     * Go to '/configMapGenerator/FromLiterals/'
     * suffix is added to the specified name
* `kubectl apply -k .`
  * `-k` looks for 'kustomization.yaml' file in the current path

# Note
* ``cat <<Whatever > path/fileName.yaml
  line1
  ...
  Whatever``
  * `>`
    * redirect the content to introduce to
  * [Here document](https://tldp.org/LDP/abs/html/here-docs.html)
    * `<<Whatever`
      * 's beginning
    * `Whatever`
      * 's end
* How to create the 'kustomization.yaml' files?
  * Check 'configMapGeneratorCommands.txt' file