# audit.json
* "defaultAction"
  * Seccomp profile's field for the system calls which don't match any of the specified filter rules
* "SCMP_ACT_LOG"
  * If the containerized process make a system call matching it ->
    * no effect on the system call
    * is logged
  * [Reference link](https://man7.org/linux/man-pages/man3/seccomp_rule_add.3.html)

# violation.json
* "SCMP_ACT_ERRNO"
  * If the containerized process make a system call matching the filter rule
  * [Reference link](https://man7.org/linux/man-pages/man3/seccomp_rule_add.3.html)
  * [errno](errno)

# fine-grained.json
* "architectures"
  * Seccomp profile's field which refers to the target CPU architectures
* "syscalls"
  * Linux system calls
  * := fundamental interface between application < -- > Linux kernel
  * [Reference link](https://man7.org/linux/man-pages/man2/syscalls.2.html)