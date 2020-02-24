FROM golang:1.13 AS builder
WORKDIR /build
COPY . .
RUN go mod vendor
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' .

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/accware/api /go/bin/accware/api
COPY schema.sql /
ENTRYPOINT ["/go/bin/accware-api"]
