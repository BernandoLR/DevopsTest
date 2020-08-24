#!/bin/bash

cd back-end && docker build -t bernando/backend .
cd ../front-end && docker build -t bernando/frontend .

docker push bernando/backend
docker push bernando/frontend


cd .. && helm install pintu ./helm-pintu
