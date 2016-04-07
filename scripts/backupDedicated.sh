FILE="/var/vcap/store/redis"
redis_command=$(awk '/rename-command BGSAVE/{print $3}' $FILE/redis.conf)
redis_port=$(awk '/port/{print $2}' $FILE/redis.conf)
redis_pwd=$(awk '/requirepass/{print $2}' $FILE/redis.conf)
last_save=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd LASTSAVE)
bg_save=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd ${redis_command//\"/})
new_last_save=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd LASTSAVE)
until [ $new_last_save -gt $last_save ]
do
  new_last_save=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd LASTSAVE)
done
cd /var/vcap/store/ && tar cz redis
