FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go build -o /app/tmp ./main.go

EXPOSE 9000:9000

CMD ["./tmp"]

