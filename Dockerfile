FROM golang:latest as builder

WORKDIR /app

COPY ./lunar-gateway/go.mod ./lunar-gateway/go.sum ./

COPY ./oort ../oort
COPY ./magnetar ../magnetar
COPY ./iam-service ../iam-service

RUN go mod download

COPY ./oort ../oort
COPY ./magnetar ../magnetar
COPY ./iam-service ../iam-service

COPY ./lunar-gateway .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/config/config.yml /root/
COPY --from=builder /app/config/no_auth_config.yml /root/

EXPOSE 5555

CMD ["./main"]
