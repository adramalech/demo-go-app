TAG_VERSION=0.1.2

docker-login:
	docker login

docker-build:
	docker build -t demo-app-go .

docker-tag:
	docker tag demo-app-go jthrone/demo-app-go:$(TAG_VERSION)

docker-push:
	docker push jthrone/demo-app-go:$(TAG_VERSION)

docker-pull:
	docker pull jthrone/demo-app-go:$(TAG_VERSION)

docker-run:
	docker run --publish 3000:3000 --detach jthrone/demo-app-go:$(TAG_VERSION)

docker-publish: docker-login docker-build docker-tag docker-push

docker-deploy: docker-login docker-pull docker-run
	
	
