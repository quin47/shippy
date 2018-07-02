build:
	protoc -I. --go_out=. --micro_out=.  \
		proto/user/user.proto
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
	docker build -t user-service  .

run:
	docker run -p 50053:50051 \
	 -e MICRO_SERVER_ADDRESS=:50051 \
      -e MICRO_REGISTRY=mdns \
       -e DB_NAME="postgres"\
       -e DB_HOST="192.168.50.213"\
       -e DB_PORT="5432"\
       -e DB_USER="postgres"\
       -e DB_PASSWORD="postgres"\
       user-service

test:
	 MICRO_SERVER_ADDRESS=:50051 \
      MICRO_REGISTRY=mdns \
       DB_NAME="postgres" \
       DB_HOST="192.168.50.213" \
        DB_PORT="5432" \
        DB_USER="postgres" \
        DB_PASSWORD="postgres"\
        go run *.go
