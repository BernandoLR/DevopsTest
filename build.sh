#!/bin/bash

# This command for build image and tag image to bernando/backend.
# You can change tag to you docker registry
cd back-end && docker build -t bernando/backend .

# This command for build image and tag image to bernando/backend.
# You can change tag to you docker registry
cd ../front-end && docker build -t bernando/frontend .

# This command for push image to you docker registry
docker push bernando/backend

# This command for push image to docker registry
docker push bernando/frontend


# This command for install helm chart with name pintu
cd .. && helm install pintu ./helm-pintu

# This command for set minikube ip to domain pintu.test.org
dns=`minikube ip`
echo $dns  pintu.test.org > /etc/hosts
