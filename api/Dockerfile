FROM golang:1.14.2-alpine3.11

RUN apk --no-cache update && \
    apk --no-cache add python py-pip py-setuptools && \
    pip --no-cache-dir install awscli && \
    rm -rf /var/cache/apk/*

#create a file
RUN mkdir /app

#copies file from . to /app
ADD . /app

#change current directory to /app
WORKDIR /app

#get deps
RUN go mod download

#build
RUN go build 

#expose this port bc we the server runs on
#port 8080
EXPOSE 8080

#docker will start the server when the 
#container starts 
ENTRYPOINT [ "go",  "run", "server/server.go" ] 

