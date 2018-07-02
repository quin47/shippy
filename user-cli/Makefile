build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
	docker build  -t user-cli .
run:
	docker run -e MICRO_REGISTRY=mdns user-cli
test:
	MICRO_REGISTRY=mdns \
	go run *.go
