FROM golang:1.18.3-alpine3.16
WORKDIR /opt
COPY ./ /opt/api-load-test
WORKDIR /opt/api-load-test
RUN go mod download
RUN go build -v -o build
ENTRYPOINT ["sh", "-c","./build"]