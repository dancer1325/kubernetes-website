# TODO:

# "/statusz"

* TODO:
* _Example of response:_
  ```
  kube-apiserver statusz
  Warning: This endpoint is not meant to be machine parseable, has no formatting compatibility guarantees and is for debugging purposes only.
  
  Started: Wed Oct 16 21:03:43 UTC 2024
  Up: 0 hr 00 min 16 sec
  Go version: go1.23.2
  Binary version: 1.32.0-alpha.0.1484&#43;5eeac4f21a491b-dirty
  Emulation version: 1.32.0-alpha.0.1484
  Paths: /healthz /livez /metrics /readyz /statusz /version
  ```

## -- with --  structured response

* _Example of response:_ 
  ```json
  {
    "kind": "Statusz",
    "apiVersion": "config.k8s.io/v1alpha1",
    "metadata": {
      "name": "kube-apiserver"
    },
    "startTime": "2025-10-29T00:30:01Z",
    "uptimeSeconds": 856,
    "goVersion": "go1.23.2",
    "binaryVersion": "1.35.0",
    "emulationVersion": "1.35",
    "paths": [
      "/healthz",
      "/livez",
      "/metrics",
      "/readyz",
      "/statusz",
      "/version"
    ]
  }
  ```

# "/flagz"

* TODO:
* _Example of response:_

  ```
  kube-apiserver flags
  Warning: This endpoint is not meant to be machine parseable, has no formatting compatibility guarantees and is for debugging purposes only.
  
  advertise-address=192.168.8.2
  contention-profiling=false
  enable-priority-and-fairness=true
  profiling=true
  authorization-mode=[Node,RBAC]
  authorization-webhook-cache-authorized-ttl=5m0s
  authorization-webhook-cache-unauthorized-ttl=30s
  authorization-webhook-version=v1beta1
  default-watch-cache-size=100
  ```

## -- with --  structured response

* _Example of response:_
  ```json
  {
    "kind": "Flagz",
    "apiVersion": "config.k8s.io/v1alpha1",
    "metadata": {
      "name": "kube-apiserver"
    },
    "flags": {
      "advertise-address": "192.168.8.4",
      "allow-privileged": "true",
      "anonymous-auth": "true",
      "authorization-mode": "[Node,RBAC]",
      "enable-priority-and-fairness": "true",
      "profiling": "true",
      "default-watch-cache-size": "100"
    }
  }
  ```