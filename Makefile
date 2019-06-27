CONTAINER_NAME=some-nginx
CONTAINER_ID=$(shell docker ps -a -f 'name=^some-nginx$$' --format '{{.ID}}')

docker-kill:
	@if [ ! -z $(CONTAINER_ID) ]; then docker kill $(CONTAINER_ID) > /dev/null; fi

docker-rm:
	@if [ ! -z $(CONTAINER_ID) ]; then docker rm $(CONTAINER_ID); fi

docker-shutdown: docker-kill docker-rm

docker-run:
	docker run -p 8080:80 --name $(CONTAINER_NAME) \
		-v $(shell pwd):/usr/share/nginx/html:ro \
		-v $(shell pwd)/mime.types:/etc/nginx/mime.types:ro \
		-v $(shell pwd)/nginx.conf:/etc/nginx/nginx.conf:ro \
		-d \
		nginx

docker-start: docker-shutdown docker-run

build:
	GOOS=js GOARCH=wasm go build -o main.wasm

clean:
	rm ./main.wasm

install: clean build docker-start