# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

RUN apk update && apk add --no-cache build-base git

ADD . /app
WORKDIR /app

COPY *.go ./
RUN go mod vendor && CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /go-backend 

EXPOSE 9997

FROM alpine

COPY conf /app/conf
COPY asset /app/asset
WORKDIR /app
COPY --from=build /go-backend /app/go-backend 
RUN apk update && apk add nano 

CMD [ "/app/go-backend" ]