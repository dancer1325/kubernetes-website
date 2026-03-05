---
title: Job
id: job
full_link: /docs/concepts/workloads/controllers/job/
short_description: >
  A finite or batch task that runs to completion.

aka: 
tags:
- fundamental
- core-object
- workload
---

* == task / runs TILL be completed
  * finite
    * == if task was succeed -> end
  * batch
    * == AUTOMATICALLY 
      * != MANUAL
  * how does it work?
    * creates >=1 pod
    * pods execute the task
    * pods end successfully
      * if some pod fail meantime -> create a NEW one
    * if specific number of successful completions is reached -> job marked as "completed"
    * job ends
  * if you 
    * delete a job -> clean up the pods created
    * suspend a job -> TILL resume the job,
      * delete its active pods
