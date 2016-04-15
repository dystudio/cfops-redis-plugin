#!/bin/bash
echo "Processing /var/vcap/store/cf-redis-broker/redis-data/"
for FILE in /var/vcap/store/cf-redis-broker/redis-data/*; do
    if [[ -d $FILE ]]; then
      echo "processing $FILE"
       if [ -w "$FILE/redis.conf" ]; then
         redis_command=$(awk '/rename-command BGSAVE/{print $3}' $FILE/redis.conf)
         redis_port=$(awk '/port/{print $2}' $FILE/redis.conf)
         redis_pwd=$(awk '/requirepass/{print $2}' $FILE/redis.conf)
         echo "issuing last_save"
         last_save=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd LASTSAVE)
         echo "result of last_save $last_save"
         sleep 1
         echo "issuing bg_save as $redis_command"
         bg_save=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd ${redis_command//\"/})
         echo "result of bg_save $bg_save"
         sleep 1
         echo "issuing last_save"
         new_last_save=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd LASTSAVE)
         echo "result of last_save $new_last_save"
         until [ $new_last_save -gt $last_save ]
         do
          sleep 1
          new_last_save=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd LASTSAVE)
          echo "checking last_save $last_save against $new_last_save"
         done
      else
        echo "$FILE/redis.conf is not writable"
        exit 1
      fi
    fi
done
