clean:
	rm -rf build

test:
	go fmt ./...
	go vet ./...
	go test -race ./...

build:
	mkdir build
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/server .

docker:
	docker build . -t git.itzana.me/strafesnet/data-service:latest
	docker push git.itzana.me/strafesnet/data-service:latest

all: clean build docker