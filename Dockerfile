
FROM golang:latest

ENV PORT=8081


WORKDIR /app
# Копируем файлы go.mod и go.sum в директорию /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD go run ./server/server.go --port=${PORT}