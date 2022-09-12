#/bin/bash
echo "跳到nginx配置目录"
cd ..
echo "停止nginx"

/usr/sbin/nginx -c nginx.conf -p $PWD/ng-config -s stop

if test $? -eq 0
then 
    echo "成功"
else
    echo "失败"
fi