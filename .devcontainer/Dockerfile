FROM golang:1.12

RUN apt update && \
    apt install -y git neovim sqlite3 && \
    rm -rf /var/lib/apt/files

# 1. install goose

RUN go get bitbucket.org/liamstask/goose/cmd/goose