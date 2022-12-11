########################################
# prod build
########################################
FROM golang:1.19.0-alpine AS builder

WORKDIR /go/src/app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

########################################
# production image
########################################
FROM alpine:3.16 AS production

WORKDIR /app
RUN apk update

# 実行に必要なバイナリをコピー
COPY --from=builder /go/src/app/main .

EXPOSE 8080

CMD ["./main"]

########################################
# Local environment
########################################
FROM golang:1.19.0 AS develop

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]

