FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN go mod init tincan-tube && go mod tidy && go build -o tincan-tube main.go

CMD ["./tincan-tube"]
