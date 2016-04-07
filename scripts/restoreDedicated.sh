#!/bin/bash
target_dir=/var/vcap/store/redis
cd /var/vcap/store/tmp_backup/ && tar xvf /var/vcap/store/tmp_backup/redis-tile.tar > /dev/null 2>&1
echo "Restoring redis"
redis_port=$(awk '/port/{print $2}' $target_dir/redis.conf)
redis_pwd=$(awk '/requirepass/{print $2}' $target_dir/redis.conf)
PID=$(/usr/bin/pgrep redis-server)
kill -9 $PID
## set to appendonly no
sed -i 's/appendonly yes/appendonly no/g' $target_dir/redis.conf
sed -i 's/rename-command BGREWRITEAOF ""/rename-command BGREWRITEAOF "BGREWRITEAOF"/g' $target_dir/redis.conf
rm -rf $target_dir/*.aof
rm -rf $target_dir/*.rdb
cp /var/vcap/store/tmp_backup/redis/dump.rdb $target_dir/.
chmod 660 $target_dir/dump.rdb

while ! nc -q 1 localhost $redis_port </dev/null; do sleep 10; done
targetProgressStatus=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd info | grep 'aof_rewrite_in_progress')
/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd BGREWRITEAOF
progressStatus=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd info | grep 'aof_rewrite_in_progress')
until [ "$progressStatus" == "$targetProgressStatus" ]
do
    progressStatus=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd info | grep 'aof_rewrite_in_progress')
    sleep 1
done

PID=$(/usr/bin/pgrep redis-server)
kill -9 $PID
sed -i 's/appendonly no/appendonly yes/g' $target_dir/redis.conf
sed -i 's/rename-command BGREWRITEAOF "BGREWRITEAOF"/rename-command BGREWRITEAOF ""/g' $target_dir/redis.conf

rm -rf /var/vcap/store/tmp_backup/redis
rm -rf /var/vcap/store/tmp_backup/redis-tile.tar
