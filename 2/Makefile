build:
	go build ../app

clean:
	rm -f ./app
	docker stop app & echo 1
		docker rmi app:small

build-image:
	docker build . -t app:small

run-in-docker: clean build-image
	docker run -d --rm --name app -p 8080:8080/tcp app:small