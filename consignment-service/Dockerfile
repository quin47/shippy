#build

FROM golang:1.10.3 as builder
# set our workdir to current service in path
WORKDIR /go/src/github.com/quin47/shippy/consignment-service
#copy current code to workdir
COPY . .
# pull deps
RUN go get -u -v github.com/golang/dep/cmd/dep
RUN dep init -v
RUN  dep ensure -v
#build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

#packaging
FROM alpine:latest
RUN mkdir /app
WORKDIR /app
#pull binary from builder
COPY --from=builder /go/src/github.com/quin47/shippy/consignment-service/consignment-service .


CMD ["./consignment-service"]