#/bin/bash
echo "跳到nginx配置目录"
cd ..
echo "优雅加载配置nginx"

/usr/sbin/nginx -c nginx.conf -p $PWD/ng-config -s reload

