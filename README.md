# cfops-redis-plugin
cfops-redis-plugin

### Version Compatibility ###
Requires 2.2.4+ of cfops

### Shared VM Backup/Restore ###
For the shared VM plan it will only restore to the same guid directories found in /var/vcap/store/cf-redis-broker/redis-data/{guid}.  If a shared instance has been add/removed since the previous backup it will be ignored.

### Dedicated VM Backup/Restore ###
For the dedicated VM plan it will only restore to the same IP that the backup was taken from.  

### Restore Notes ###
* To skip a restore of either a dedicated node or sharedVM just remove the .tar file from the directory used as -d argument as this will drive which VMs to Backup

* To only backup a subset of redis instances on a shared VM you must modify the redis-tile-sharedVMPlan.tar and remove the directories that represent the instances you don't wish to restore.

* In the case of needing to restore to a shared instance that guid has changed you must untar the redis-tile-sharedVMPlan.tar and change the folder name to match the new target and re-tar to be named redis-tile-sharedVMPlan.tar

* To restore to a different IP must rename the backup file and replace the IP part of the name with the target IP.
