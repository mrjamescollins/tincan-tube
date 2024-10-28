FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN go mod init tincan-tube/server && go mod tidy && go build -o vpn-server main.go

CMD ["./vpn-server"]