FROM golang:1.20

WORKDIR /assignment3

COPY . .

RUN go mod tidy

RUN go build -o book-api ./cmd/ 

CMD ["./book-api"]
