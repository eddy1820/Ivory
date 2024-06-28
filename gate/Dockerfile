FROM golang:latest
RUN mkdir /gate
ADD . /gate

#RUN ls -l /gate && sleep 10
WORKDIR /gate
RUN go mod download
RUN go build -o main .
EXPOSE 7500
CMD ["./main"]