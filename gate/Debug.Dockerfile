FROM golang:1.21-alpine
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN apk add --no-cache g++ git bash
RUN go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 7500
