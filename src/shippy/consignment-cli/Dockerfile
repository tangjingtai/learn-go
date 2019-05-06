FROM alpine:latest

RUN mkdir /app

WORKDIR /app

ADD consignment-client /app/consignment-client
ADD consignment.json /app/consignment.json

CMD ["./consignment-client"]