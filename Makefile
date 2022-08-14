BINARY_NAME=p2pchat-bootstrapper

.PHONY: build build_and_run clean

build:
	go build -o ./bin/${BINARY_NAME} main.go

run:
	./bin/${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ./bin/${BINARY_NAME}

docker_build:
	docker build -t p2pchat-bootstrapper .
