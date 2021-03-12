# ingress

Ingress is created to serve paths from services on port 80 and host file-server.info:
- / from index-service 
- /files from file-service 
Note: Ingress must be created in the same namespace as application

# ingress class

ingress class is created with controller k8s.io/ingress-nginx

# ingress controller

Nginx Ingress controller is created from https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.44.0/deploy/static/provider/cloud/deploy.yaml

Note: Specify deployment arguments as --ingress-class=ingress-class in order for controller to know which ingress to serve