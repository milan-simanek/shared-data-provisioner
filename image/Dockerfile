FROM golang:1.16-alpine AS builder

ADD . /build

RUN apk add make && make -C /build shared-data-provisioner

FROM scratch

COPY --from=builder /build/shared-data-provisioner /shared-data-provisioner

CMD ["/shared-data-provisioner"]
