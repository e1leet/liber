FROM golang:1.19-alpine

ENV CONFIG_PATH=configs/config.local.env

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/liber ./cmd/app/main.go

CMD ["./bin/liber"]