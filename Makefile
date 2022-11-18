build:
	go build app.go

run-app: build
	./app


clean:
	rm -f ./app
	docker stop app & echo 1
	docker rmi app:big & echo 1

build-image-big:
	docker build . -f 1-docker/Dockerfile -t app:big

build-image-small:
	docker build . -f 2-docker/Dockerfile -t app:small


build-image-k8s:
	docker build . -f 3-k8s/Dockerfile -t app:k8s

run-in-docker-big: clean build-image-big
	docker run -d --rm --name app -p 8080:8080/tcp app:big

run-in-docker-small: clean build-image-small
	docker run -d --rm --name app -p 8080:8080/tcp app:small

push-image: build-image-k8s
	docker tag app:k8s orlix/app:k8s
	docker push orlix/app:k8s