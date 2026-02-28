---
title: Secret
id: secret
full_link: /docs/concepts/configuration/secret/
short_description: >
  Stores sensitive information, such as passwords, OAuth tokens, and ssh keys.

aka:
tags:
- core-object
- security
---

* secret
  * := API object /
    * enable you,
      * more control about how to use sensitive information /
        * reduces the risk of accidental exposure | create OR view OR edit pods
          * Reason: 🧠secrets can be created INDEPENDENTLY of the pods / use them🧠
        * ALTERNATIVE to store | 
          * pod specification, OR
          * container image
    * 's values
      * ⚠️by default,⚠️
        * encoded -- as -- base64 strings
        * stored unencrypted
      * you can [encrypt at rest](../../tasks/administer-cluster/encrypt-data.md#ensure-all-relevant-data-are-encrypted-ensure-all-secrets-are-encrypted)
    * can be reused ACROSS MULTIPLE containers 
  * ways / pod can reference a secret
    * volume mount
    * environment variable
  * use cases
    * store sensitive information
      * _Example:_ passwords, OAuth tokens, SSH keys
  * vs [ConfigMaps](configmap.md)
    * configMaps use cases
      * NON-confidential data
