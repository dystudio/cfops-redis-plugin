#!/bin/bash
target_dir=/var/vcap/store/cf-redis-broker/redis-data/
cd /var/vcap/store/tmp_backup/ && tar xvf /var/vcap/store/tmp_backup/redis-tile.tar > /dev/null 2>&1
for FILE in /var/vcap/store/tmp_backup/redis-data/*; do
    if [[ -d $FILE ]]; then
       directoryName=$(basename $FILE)
       installDir=$target_dir$directoryName
       if [ -d "$installDir" ]; then
          echo "Restoring $directoryName"
          redis_port=$(awk '/port/{print $2}' $installDir/redis.conf)
          redis_pwd=$(awk '/requirepass/{print $2}' $installDir/redis.conf)
          PID=$(cat $installDir/redis-server.pid)
          kill -9 $PID
          rm -rf $installDir/redis-server.pid
          ## set to appendonly no
          sed -i 's/appendonly yes/appendonly no/g' $installDir/redis.conf
          sed -i 's/rename-command BGREWRITEAOF ""/rename-command BGREWRITEAOF "BGREWRITEAOF"/g' $installDir/redis.conf
          rm -rf $installDir/db/*
          cp /var/vcap/store/tmp_backup/redis-data/$directoryName/db/dump.rdb $installDir/db/.
          chmod 660 $installDir/db/dump.rdb

          while ! nc -q 1 localhost $redis_port </dev/null; do sleep 10; done
          targetProgressStatus=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd info | grep 'aof_rewrite_in_progress')
          /var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd BGREWRITEAOF
          progressStatus=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd info | grep 'aof_rewrite_in_progress')
          until [ "$progressStatus" == "$targetProgressStatus" ]
          do
            progressStatus=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd info | grep 'aof_rewrite_in_progress')
            sleep 1
          done

          PID=$(cat $installDir/redis-server.pid)
          kill -9 $PID
          sed -i 's/appendonly no/appendonly yes/g' $installDir/redis.conf
          sed -i 's/rename-command BGREWRITEAOF "BGREWRITEAOF"/rename-command BGREWRITEAOF ""/g' $installDir/redis.conf
       else
          echo "$installDir doesn't exist"
       fi
    fi
done
rm -rf /var/vcap/store/tmp_backup/redis-data
rm -rf /var/vcap/store/tmp_backup/redis-tile.tar
