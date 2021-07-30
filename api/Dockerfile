FROM golang:latest

EXPOSE $PORT

WORKDIR /usr/src/app/api

COPY . /usr/src/app/api

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon@master
RUN GO111MODULE=on
# RUN CompileDaemon
ENTRYPOINT CompileDaemon -build="go build main.go" -directory="." -polling -command=./main 
