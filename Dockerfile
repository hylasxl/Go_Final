
FROM golang:1.23.2-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o go_final

EXPOSE 50051

CMD ["./go_final"]