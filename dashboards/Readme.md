Graphana install
https://grafana.com/grafana/download?platform=mac
curl -O https://dl.grafana.com/enterprise/release/grafana-enterprise-10.4.1.darwin-amd64.tar.gz
tar -zxvf grafana-enterprise-10.4.1.darwin-amd64.tar.gz

Point Domain name
grafana.office.net A 142.92.23.24

Reverse Proxy grafana with Nginz
grafana.office.net:3000 grafana.office.net
```
cd /etc/nginx/sites-enabled
sudo nano grafana.office.net.conf
server {
    listen 80;
    listen [::]:80;
    server_name  grafana.office.net;

    location / {
        proxy_set_header Host $http_host;
        proxy_pass           http://localhost:3000/;
    }
}
```

Restart nginz
```
sudo service nginx restart
sudo service nginx status
```
Add SSL
http://grafana.office.net -> https://grafana.office.net
- To get the certificate,  share cert request and get the signed certificate from CA(cloudflare etc.)
- associate the certificate against app(using tools like IIS/HTTP.sys etc).
RunTls //https://pkg.go.dev/github.com/gin-gonic/gin
https://learn.microsoft.com/en-us/iis/manage/configuring-security/how-to-set-up-ssl-on-iis


Graphna start
./grafana-server
or
systemctl start grafana-server

#Graphana version

http://localhost:3000/api/health
