#!/usr/bin/env sh

echo "--------------初始化配置----------------------"

sed -i "s/execution/$server/g" /etc/supervisor/conf.d/application.conf

cp app.$env.yaml app.yaml

sed -i "s/micro_scrm.log/micro_scrm_$server.log/g" app.yaml

if [ "$server" = "job" ]; then
  rm -rf /etc/supervisor/conf.d/application.conf
  sed -i "s/console/file/g" app.yaml
fi

supervisord
