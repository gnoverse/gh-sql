FROM golang:1.23-alpine AS build

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download -x

COPY . /build
RUN go build -v

FROM busybox:latest

WORKDIR /gh-sql

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /build/gh-sql /bin/gh-sql

CMD gh-sql --help
