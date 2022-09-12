#/bin/bash
echo "改变常用文件路径"
sudo chown -R wx /var/log/nginx 
sudo chown -R wx /var/cache/nginx 

echo "设置ip转换"
sudo iptables -A PREROUTING -t nat -p tcp --dport 80 -j REDIRECT --to-port 8080