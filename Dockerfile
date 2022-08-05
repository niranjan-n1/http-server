FROM golang:1.18.5-alpine3.16

COPY .  /home/app/
WORKDIR /home/app/
RUN go mod tidy
RUN go build -o http-server
EXPOSE 8081
CMD [ "./http-server" ]