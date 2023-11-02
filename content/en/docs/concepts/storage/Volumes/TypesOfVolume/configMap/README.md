# Steps
* `kubectl apply -f configMap.yaml`
* `kubectl apply -f pod.yaml`
* `kubectl exec -it configmap-pod sh`
  * Go into the pod
  * `ls /etc/config`
    * There is just 1! file 'log_level', since it was just specified it
  * `cat /etc/config/log_level`
    * display the value of that configMap's entry
  * `rm /etc/config/log_level`
    * you can NOT remove it, since it's READ-ONLY file system
  * `ls -l /etc/config/log_level` with output 'lrwxrwxrwx'
    * 'l -- symbolic link -- to the configMap resource --
    * 'rwx' represents the owner's permissions
      * 'r' read permission
      * 'w' write permission, but on the real one, not in the symbolic link
      * 'x' execute / search permission
    * 'rwx' represents the group's permissions
    * 'rwx' represents other user's permissions