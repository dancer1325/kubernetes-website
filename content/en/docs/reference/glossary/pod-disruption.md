---
id: pod-disruption
title: Pod Disruption
full_link: /docs/concepts/workloads/pods/disruptions/
short_description: >
  The process by which Pods on Nodes are terminated either voluntarily or involuntarily.

aka:
related:
 - pod
 - container
tags:
 - operation
---

- := process of terminating PodS | nodes
  - types of terminations
    - voluntarily
      - triggered -- by --
        - application owners or
        - cluster administrators
    - involuntarily
      - triggered -- by --
        - unavoidable issues
          - _Example1:_ Nodes run out of [resources](infrastructure-resource.md)
          - _Example2:_ accidental deletions
  - [MORE](disruption.md)
