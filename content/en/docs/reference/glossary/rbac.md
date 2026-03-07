---
title: RBAC (Role-Based Access Control)
id: rbac
full_link: /docs/reference/access-authn-authz/rbac/
short_description: >
  Manages authorization decisions, allowing admins to dynamically configure access policies through the Kubernetes API.

aka: 
tags:
- security
- fundamental
---

* RBAC
  * == method /
    * regulate access -- , based on roles of individual users | your organization, -- to computer OR network resources
  * API group: `rbac.authorization.k8s.io`
  * uses
    * cluster admins can DYNAMICALLY configure -- , through [Kubernetes API](kubernetes-api.md), -- access policies
  * 👀Kubernetes objects / declare👀
    * `Role`
    * `ClusterRole`
    * `RoleBinding`
    * `ClusterRoleBinding`
  * [MORE](../access-authn-authz/rbac)

## RBAC's objects related
### `Role`

* == permission ruleS | specific namespace
  * | create a role, you need to specify the namespace
* 's name
  * [path segment name valid](../../concepts/overview/working-with-objects/names#path-segment-names)

### `ClusterRole`

* == permission rules | cluster-wide
* 's name
  * [path segment name valid](../../concepts/overview/working-with-objects/names#path-segment-names)
* uses
  * define permissions | namespaced resources & be granted access | individual namespace(s)
  * define permissions | namespaced resources & be granted access ACROSS ALL namespaces
    * _Example:_ [pods](pod.md)
  * define permissions | cluster-scoped resources
    * _Example:_ [nodes](node.md)
  * define permissions | non-resource endpoints 
    * _Example:_ "/healthz"

### `RoleBinding`

* allows
  * role's permissions can be granted -- to a -- set of subjects (users, groups, [service accounts](service-account.md)) | specific namespace
    * role can be specified -- via --
      * DIRECTLY a role
      * ClusterRole
* 's name
  * [path segment name valid](../../concepts/overview/working-with-objects/names#path-segment-names)
* AFTER creating it, if you want to change the referred `Role`
  * ❌you can NOT change the referred one❌
    * Reason: 🧠
      * privilege escalation vulnerability
      * DIFFERENT binding would be🧠
  * steps
    * remove the `RoleBinding` object
    * create a NEW one / changed

### `ClusterRoleBinding`

* allows
  * role's permissions can be granted -- to a -- set of subjects (users, groups, [service accounts](service-account.md)) | cluster-wide
* 's name
  * [path segment name valid](../../concepts/overview/working-with-objects/names#path-segment-names)
* AFTER creating it, if you want to change the referred `ClusterRole`
  * ❌you can NOT change the referred one❌
    * Reason: 🧠
      * privilege escalation vulnerability
      * DIFFERENT binding would be🧠
  * steps
    * remove the `ClusterRoleBinding` object
    * create a NEW one / changed

## notes

* rule
  * == set of permissionS

* permission
  * ⚠️ONLY additive⚠️
    * ❌there are NO "deny" rules❌
