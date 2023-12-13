FROM golang:latest as builder

WORKDIR /app

COPY ./gateway/go.mod ./gateway/go.sum ./

RUN go mod download

COPY ./gateway/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/config/config.yml /root/

EXPOSE 8000

CMD ["./main"]