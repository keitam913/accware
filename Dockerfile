FROM golang:1.13 AS builder

WORKDIR /build

COPY . /build

RUN go build -o accware --ldflags "-linkmode 'external' -extldflags '-static'" .

FROM scratch

WORKDIR /bin

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /build/accware /bin/accware
COPY ./schema.sql /usr/share/accware/schema.sql
COPY ./assets /usr/share/accware/assets

EXPOSE 80

ENTRYPOINT [ "/bin/accware" ]
