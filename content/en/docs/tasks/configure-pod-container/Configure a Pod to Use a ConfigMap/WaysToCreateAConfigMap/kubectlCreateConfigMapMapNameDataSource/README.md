# Content
* `kubectl create configmap MapName DataSource`
    * MapName
      * configMap name
    * DataSource
      * types
        * directory
          * `mkdir files`, `vim files/ui.properties` and `vim files/game.properties`
            * Create a directory and files included
          * Download the files in the previous directory
            * Via `wget`
              * `wget https://kubernetes.io/examples/configmap/game.properties -O configure-pod-container/configmap/game.properties`
              * `wget https://kubernetes.io/examples/configmap/ui.properties -O configure-pod-container/configmap/ui.properties`
            * Via `curl`
              * `curl -o files/game.properties  https://kubernetes.io/examples/configmap/game.properties`
              * `curl -o files/ui.properties https://kubernetes.io/examples/configmap/ui.properties`
          * `kubectl create configmap game-config --from-file=files/`
            * Create the configMap with name "game-config" based on the directory
        * file
          * `--from-file`
            * `mkdir files`, `vim files/ui.properties` and `vim files/game.properties`
              * Create a directory and files included
            * Download the files in the previous directory
              * Via `wget`
                * `wget https://kubernetes.io/examples/configmap/game.properties -O configure-pod-container/configmap/game.properties`
                * `wget https://kubernetes.io/examples/configmap/ui.properties -O configure-pod-container/configmap/ui.properties`
              * Via `curl`
                * `curl -o files/game.properties  https://kubernetes.io/examples/configmap/game.properties`
                * `curl -o files/ui.properties https://kubernetes.io/examples/configmap/ui.properties`
            * `kubectl create configmap game-config-files --from-file=path1/file1 --from-file=path2/file2`
              * Create the configMap with name "game-config-files" based on files
            * `kubectl create configmap game-config-files --from-file=Key1=path1/file1 --from-file=key2=path2/file2`
              * Create the configMap with name "game-config-files" based on files, but passing a key for the dataSource
          * `--from-env-file`
            * `mkdir env-files`, `vim env-files/ui-env-file.properties` and `vim env-files/game-env-file.properties`
              * Create a directory and files included
            * Download the files in the previous directory
              * Via `wget`
                * `wget https://kubernetes.io/examples/configmap/game-env-file.properties -O configure-pod-container/configmap/game-env-file.properties`
                * `wget https://kubernetes.io/examples/configmap/ui-env-file.properties -O configure-pod-container/configmap/ui-env-file.properties`
              * Via `curl`
                * `curl -o env-files/game-env-file.properties https://kubernetes.io/examples/configmap/game-env-file.properties`
                * `curl -o env-files/ui-env-file.properties https://kubernetes.io/examples/configmap/ui-env-file.properties`
            * `kubectl create configmap config-multi-env-files \
              --from-env-file=env-files/game-env-file.properties \
              --from-env-file=env-files/ui-env-file.properties`
              * Kubernetes > v1.23
        * literal value
          * `kubectl create configmap withliteralvalues --from-literal=key1=value1 --from-literal=key2=value2`
            * Create the configMap with name "game-config" based on the directory