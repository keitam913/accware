FROM golang:1.13 AS builder
WORKDIR /build
COPY . /build
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/accware-api /
COPY schema.sql /
EXPOSE 80
CMD ["./accware-api"]
