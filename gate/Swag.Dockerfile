FROM golang:latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

EXPOSE 7500
