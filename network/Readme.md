## Golang Http file server k8s accessed through nginx ingress

Serving files on file-server.info host

In order to resolve file-server.info host to access file server:

- Get Public ip of ingress controller load balancer service
- Update /etc/hosts file in your linux machine and enter this entry:
  <public ip> file-server.info
 
