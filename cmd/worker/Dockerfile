FROM golang:1.18-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ./cmd/worker/main.go

# Step 2
FROM alpine:latest

RUN apk --no-cache add tzdata

RUN apk --no-cache add ca-certificates

WORKDIR /app/

RUN ls -al

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/main .

EXPOSE 8808

CMD ["./main"]