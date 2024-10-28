FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN go mod init tincan-tube-client && go mod tidy && go build -o vpn-client main.go

CMD ["./vpn-client"]