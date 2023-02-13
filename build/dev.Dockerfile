FROM golang:1.19-alpine

ENV CONFIG_PATH=configs/config.local.env

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

COPY . .

ENTRYPOINT CompileDaemon -build="go build -o ./bin/liber ./cmd/app/main.go" -command="./bin/liber"