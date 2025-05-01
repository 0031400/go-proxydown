### title
go-proxy-down
### summary
A proxy download tool written in golang.
### usage
GET /?url=<your download url>
### feature
when the response is a html.It will transform it into plain text.
### config
|config|example|
|-|-|
|host|0.0.0.0|
|port|8080|
### log
It has a simple log for the successful proxydown.
### run
./proxydown --host [::] --port 8080
### nginx
the example nginx config is on `proxydown.conf`
### service
the example service config is on `proxydown.service`