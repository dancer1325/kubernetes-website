---
title: Volume
id: volume
full_link: /docs/concepts/storage/volumes/
short_description: >
  A directory containing data, accessible to the containers in a pod.

aka:
tags:
- fundamental
---

- volume
  - := directory / contains data
    - accessible -- to the — [container](container.md) | pod 
  - [storage](../../concepts/storage)
  - ’s life time 
    - = pod’s lifetime
      - related -- to -- pod
      - ❌NOT related -- to -- container❌
        - == ⚠️if container restarts OR crashes, volume persists ⚠️
