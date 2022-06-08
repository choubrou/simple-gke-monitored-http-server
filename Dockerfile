# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY main.go ./
RUN go build -o /my-http-server

FROM alpine
WORKDIR /
COPY --from=build /my-http-server /my-http-server
EXPOSE 8080
#USER nonroot:nonroot
ENTRYPOINT [ "/my-http-server" ]
