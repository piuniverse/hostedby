#Build Stage - Create a go standalone binary using fat container
FROM golang:1.19-alpine AS builder

RUN apk --update add \
    go \
    musl-dev
RUN apk --update add \
    util-linux-dev
RUN apk add --no-cache tzdata
RUN apk --update --no-cache add curl
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache gcc g++ git openssh-client

COPY go.mod /project/go.mod
COPY go.sum /project/go.sum
COPY pkg /project/pkg
COPY /cmd/ /project/cmd

WORKDIR /project/cmd/api
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o server .

#Create the actual container no with just the binary
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /project/cmd/api
COPY --from=builder /project/cmd/api/ /project/cmd/api/


ENTRYPOINT ["/project/cmd/api/server"] 