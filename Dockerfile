FROM alpine:latest

ENV app=tfc-transfer-validator

RUN mkdir /app
WORKDIR /app
ADD $app /app/$app

CMD ./$app