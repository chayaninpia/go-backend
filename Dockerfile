# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

RUN apk update && apk add --no-cache git

ADD . /app
WORKDIR /app

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /go-backend 

EXPOSE 9997

FROM alpine

COPY conf /app/conf
WORKDIR /app
COPY --from=build /go-backend /app/go-backend 
RUN apk update && apk add nano 

CMD [ "/app/go-backend" ]