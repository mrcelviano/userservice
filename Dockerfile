FROM golang:alpine as builder
RUN mkdir -p /go/src/userservice
ENV SOCIAL_TECH_ENV=local
RUN apk update && apk upgrade && apk add git && export GOOS=linux && export GOARCH=amd64
RUN apk update && apk upgrade && apk --no-cache add ca-certificates && update-ca-certificates
WORKDIR /go/src/userservice
COPY . .
RUN go build -o /compiled/userservice && mkdir -p /compiled/configs
COPY ./configs /compiled/configs

FROM alpine:latest
RUN mkdir -p /go/src/userservice && chmod -R 0777 /go/* && apk add --no-cache bash && apk add --no-cache tzdata
RUN apk update && apk upgrade && apk --no-cache add ca-certificates && update-ca-certificates
COPY --from=builder /compiled/ /go/src/userservice
WORKDIR /go/src/userservice/
RUN chmod +x userservice
ENV PATH="/go/src/userservice"
ENV SOCIAL_TECH_ENV=local
EXPOSE 8080
CMD ["userservice"]
