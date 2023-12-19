FROM golang:1.19.3-alpine3.16 as builder

WORKDIR /go/src/api

COPY . .

RUN apk add --update make
RUN make build

FROM alpine:3.16

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/api/api /app
COPY --from=builder /go/src/api/.env /app

EXPOSE 8080

CMD [ "/app/api" ]
