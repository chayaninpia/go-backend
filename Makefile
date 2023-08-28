current_dir=$(shell pwd)
project_name=$(shell basename "${current_dir}" )
version=$(shell echo "1-dev")

build:
	docker build -t $(project_name) .

run:
	docker-compose -f deploy/docker-compose.yaml up -d

stop:
	docker-compose -f deploy/docker-compose.yaml down

runpg:
	docker-compose -f deploy/postgres.yaml up -d

stoppg:
	docker-compose -f deploy/postgres.yaml down

runkaf:
	docker-compose -f deploy/kafka.yaml up -d

stopkaf:
	docker-compose -f deploy/kafka.yaml down

push:
#docker login
	docker tag $(project_name) ncnewvirus2/$(project_name):$(version)
	docker push ncnewvirus2/$(project_name):$(version)

ingress:
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.6.4/deploy/static/provider/cloud/deploy.yaml
	
deploy all:
	kubectl apply -f deploy/k8s/.

swag-init:
	swag init --parseDependency --parseInternal --parseDepth 1