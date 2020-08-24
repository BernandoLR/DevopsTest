## General info
This project is simple kubernetes test with the requirements using Helm in operating system ubuntu.
	
## Technologies
Project is created with:
* Minikube version: 1.12.1
* Helm version: 3.3.0
* Go : 1.14.1
	
## Setup
To run this project, install it locally using Minikube and Helm:

* Minikube
```
$ curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 \
  && chmod +x minikube
$ mv minikube /usr/local/bin/
$ minikube start
$ minikube addons enable metrics-server
$ minikube addons enable ingress
```

* Helm
```
$ wget https://get.helm.sh/helm-v3.3.0-linux-amd64.tar.gz
$ tar -zxvf helm-v3.0.0-linux-amd64.tar.gz
$ mv linux-amd64/helm /usr/local/bin/
```

* Bash script

Just try to run simple bash script

```
$ sudo bash build.sh
```

## How to access the project
This project can access with hit domain pintu.test.org in your browser 
