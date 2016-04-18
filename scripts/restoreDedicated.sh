#!/bin/bash
target_dir=/var/vcap/store/redis
cd /var/vcap/store/tmp_backup/ && tar xvf /var/vcap/store/tmp_backup/redis-tile.tar > /dev/null 2>&1
echo "Restoring redis"
redis_port=$(awk '/port/{print $2}' $target_dir/redis.conf)
redis_pwd=$(awk '/requirepass/{print $2}' $target_dir/redis.conf)
PID=$(/usr/bin/pgrep redis-server)
echo "stopping server with PID $PID"
kill -9 $PID
## set to appendonly no
sed -i 's/appendonly yes/appendonly no/g' $target_dir/redis.conf
rm -rf $target_dir/*.aof
rm -rf $target_dir/*.rdb
cp /var/vcap/store/tmp_backup/redis/dump.rdb $target_dir/.
chmod 660 $target_dir/dump.rdb
echo "Restarting monit"
/var/vcap/bosh/bin/monit restart all

echo "Waiting for server to start on port $redis_port"
while ! nc -q 1 localhost $redis_port </dev/null; do sleep 10; done
  targetProgressStatus=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd info | grep 'aof_rewrite_in_progress')
  echo "Target Progress Status $targetProgressStatus"
  /var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd BGREWRITEAOF
  sleep 1
  progressStatus=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd info | grep 'aof_rewrite_in_progress')
  echo "Progress Status $progressStatus"
  until [ "$progressStatus" == "$targetProgressStatus" ]
  do
    progressStatus=$(/var/vcap/packages/redis/bin/redis-cli -h 127.0.0.1 -p $redis_port -a $redis_pwd info | grep 'aof_rewrite_in_progress')
    sleep 1
  done

PID=$(/usr/bin/pgrep redis-server)
echo "stopping server with PID $PID"
kill -9 $PID
sed -i 's/appendonly no/appendonly yes/g' $target_dir/redis.conf

echo "Restarting monit"
/var/vcap/bosh/bin/monit restart all

rm -rf /var/vcap/store/tmp_backup/redis
rm -rf /var/vcap/store/tmp_backup/redis-tile.tar
