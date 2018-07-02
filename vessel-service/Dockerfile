
FROM golang:1.10.3 as builder
# set our workdir to current service in path
WORKDIR /go/src/github.com/quin47/shippy/vessel-service
#copy current code to workdir
COPY . .
# pull deps
RUN go get -u -v github.com/golang/dep/cmd/dep
RUN dep init -v
RUN  dep ensure -v
#build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/quin47/shippy/vessel-service/vessel-service .

CMD ["./vessel-service"]