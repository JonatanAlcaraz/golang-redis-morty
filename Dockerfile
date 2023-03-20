FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go mod download github.com/golang/protobuf
RUN go build -o main .

EXPOSE 3000

CMD ["./main"]
