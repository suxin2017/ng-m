#/bin/bash
echo "改变常用文件路径"
chown -R wx /var/log/nginx 
chown -R wx /var/cache/nginx 

echo "设置ip转换"
iptables -A PREROUTING -t nat -p tcp --dport 80 -j REDIRECT --to-port 8080