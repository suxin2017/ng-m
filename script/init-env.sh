#/bin/bash
echo "改变常用文件路径"
sudo chown -R wx /var/log/nginx 
sudo chown -R wx /var/cache/nginx 

echo "nginx可以直接监听80端口"
sudo setcap  CAP_NET_BIND_SERVICE=+eip /usr/sbin/nginx

echo "install zip package"
sudo apt isntall unzip