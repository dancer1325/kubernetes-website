---
title: "Production-Grade Container Orchestration"
abstract: "Automated container deployment, scaling, and management"
cid: home
layout: home
sitemap:
  priority: 1.0
---

- Kubernetes OR k8s
  - == 💡container orchestrator💡
    - production-grade
    - open-source
    - allows, about containerized applications,
      - automatic deployments
        - _Example:_ canary
      - scaling & failover
      - management
  - Greek original word
    - ==
      - helmsman
      - pilot
  - k8s
    - 8
      - 8 letters BETWEEN “k” -- & -- “s” | “kubernetes”
  - 👀[origin: Google](https://queue.acm.org/detail.cfm?id=2898444)👀
  - 👀characteristics👀
    - planet scale
      - === can scale WITHOUT increasing your DevOps team
        - _Example:_ Google run billions of containers / week 👀
    - never outgrow
      - Reason: 🧠flexible growing🧠
      - _Example:_ [small startups / 3-5 nodes, big companies / run 1xx nodes]
    - run K8s | ANYWHERE (on-premises, hybrid, cloud, ..., any OS, ...)
  - [features](docs/concepts/overview/_index.md)

- containers
  - 👀are grouped -- into -- [pod](docs/reference/glossary/pod.md)👀
    - Reason: 🧠
      - pod == Kubernetes object
      - container != Kubernetes object🧠


TODO:

{{< blocks/section id="video" background-image="video_banner_homepage_XNZvGHlpJ_4" >}}
<div class="video-text-optional">
    <h2>The Absolute Beginner's Guide To Cloud Native</h2>
    <p class="presenter-byline">by Kyle Penfound, Dagger &amp; Cortney Nickerson, Kubeshop</p>
    <button id="desktopShowVideoButton" onclick="kub.showVideo()">Watch Video</button>
</div>
<div id="videoPlayer">
    <iframe data-url="https://www.youtube.com/embed/XNZvGHlpJ_4?autoplay=1" frameborder="0" allowfullscreen></iframe>
    <button id="closeButton"></button>
</div>
{{< /blocks/section >}}
{{< blocks/section id="upcoming-events" >}}
<h2>Attend upcoming KubeCon + CloudNativeCon events</h2>
  <!-- TODO: change this to a shortcode that auto-updates from a schedule -->
  <div>
    <a href="https://events.linuxfoundation.org/kubecon-cloudnativecon-europe-2026/" class="desktopKCButton"><strong>Europe</strong> (Amsterdam, Mar 23-26, 2026)</a>
  </div>
  <div>
    <a href="https://events.linuxfoundation.org/kubecon-cloudnativecon-north-america-2026/" class="desktopKCButton"><strong>North America</strong> (Salt Lake City, Nov 9-12, 2026)</a>
  </div>
{{< /blocks/section >}}

{{< blocks/kubernetes-features >}}

{{< blocks/case-studies >}}
