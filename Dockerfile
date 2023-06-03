FROM golang:1.19.0

WORKDIR /usr/src/app


## เพิ่มที่หลัง

RUN go install github.com/cosmtrek/air@latest

#Updating Dockerfile
#In our Dockerfile, we will add 2 new lines. First, we will use the COPY instruction to copy all the files into the container’s working directory. Then we will run the command go mod tidy to install and clean up our dependencies

COPY . .
RUN go mod tidy