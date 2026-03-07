---
title: "Security"
weight: 85
description: >
  Concepts for keeping your cloud-native workload secure.
simple_list: true
---

* goal
  * how to 
    * run workloads MORE securely
    * keep a Kubernetes cluster secure

* Kubernetes
  * follows good practice for cloud native information security

## Kubernetes security mechanisms {#security-mechanisms}

### Control plane protection

* [control access -- to the -- Kubernetes API](controlling-access)

* recommended requirements by Kubernetes
  * [configure TLS + data encrypt data in transit -- via -- TLS](../../tasks/tls/managing-tls-in-a-cluster) |   
    * control plane
    * BETWEEN control plane -- & -- its clients
  * [encrypt data / stored | Kubernetes control plane](../../tasks/administer-cluster/encrypt-data)

### [Secrets](../configuration/secret)

### Workload protection

* [enforce pod security standards](pod-security-standards) 
* [RuntimeClasses](../containers/runtime-class)
* [Network policies](../services-networking/network-policies)

TODO: You can deploy security controls from the wider ecosystem to implement preventative
or detective controls around Pods, their containers, and the images that run in them.

### Admission control {#admission-control}

* [Admission controllers](../../reference/access-authn-authz/admission-controllers)
* [Admission Webhook Good Practices](../cluster-administration/admission-webhooks-good-practices)

### Auditing

* [audit logging](../../tasks/debug/debug-cluster/audit)

## Cloud provider security

| IaaS Provider               | Security Documentation                                                                                     |
|-----------------------------|------------------------------------------------------------------------------------------------------------|
| Alibaba Cloud               | [Trust Center](https://www.alibabacloud.com/trust-center)                                                  |
| Amazon Web Services         | [AWS Security](https://aws.amazon.com/security)                                                            |
| Google Cloud Platform       | [GCP Security](https://cloud.google.com/security)                                                          |
| Huawei Cloud                | [Security Center](https://www.huaweicloud.com/intl/en-us/securecenter/overallsafety)                       |
| IBM Cloud                   | [IBM Cloud Security](https://www.ibm.com/cloud/security)                                                   |
| Microsoft Azure             | [Azure Security](https://docs.microsoft.com/en-us/azure/security/azure-security)                           |
| Oracle Cloud Infrastructure | [Oracle Security](https://www.oracle.com/security)                                                         |
| Tencent Cloud               | [Data Security Solutions](https://www.tencentcloud.com/solutions/data-security-and-information-protection) |
| VMware vSphere              | [Security Hardening Guides](https://www.vmware.com/solutions/security/hardening-guides)                    |

## Policies

* ways to define security policies -- via -- Kubernetes-native mechanisms
  * [NetworkPolicy](../services-networking/network-policies)
  * [ValidatingAdmissionPolicy](../../reference/access-authn-authz/validating-admission-policy)

* [MORE](../policy)

## {{% heading "whatsnext" %}}

Learn about related Kubernetes security topics:

* [Securing your cluster](/docs/tasks/administer-cluster/securing-a-cluster/)
* [Known vulnerabilities](/docs/reference/issues-security/official-cve-feed/)
* [Data encryption in transit](/docs/tasks/tls/managing-tls-in-a-cluster/) for the control plane
* [Data encryption at rest](/docs/tasks/administer-cluster/encrypt-data/)
* [Controlling Access to the Kubernetes API](/docs/concepts/security/controlling-access)
* [Network policies](/docs/concepts/services-networking/network-policies/) for Pods
* [Secrets in Kubernetes](/docs/concepts/configuration/secret/)
* [Pod security standards](/docs/concepts/security/pod-security-standards/)
* [RuntimeClasses](/docs/concepts/containers/runtime-class)
* [Certified Kubernetes Security Specialist](https://training.linuxfoundation.org/certification/certified-kubernetes-security-specialist/)
