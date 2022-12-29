#/bin/bash

user=`whoami`
echo $user

echo "改变常用文件路径"
sudo chown -R $user /var/log/nginx 
sudo chown -R $user /var/cache/nginx 
sudo touch /var/log/nginx/error.log
sudo touch /var/log/nginx/access.log

echo "nginx可以直接监听80端口"
sudo setcap  CAP_NET_BIND_SERVICE=+eip /usr/sbin/nginx

echo "install zip package"
sudo apt install unzip