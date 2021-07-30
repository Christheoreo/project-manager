FROM golang:latest
EXPOSE 1323

WORKDIR /usr/src/app/go-api

COPY . /usr/src/app/go-api

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon@master
RUN GO111MODULE=on
# RUN CompileDaemon
ENTRYPOINT CompileDaemon -build="go build main.go" -directory="." -polling -command=./main 
