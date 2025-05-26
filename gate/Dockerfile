FROM golang:latest
RUN mkdir /server
ADD . /gate

#RUN ls -l /server && sleep 10
WORKDIR /gate
RUN go mod download
RUN go build -o main .
EXPOSE 7500
CMD ["./main"]