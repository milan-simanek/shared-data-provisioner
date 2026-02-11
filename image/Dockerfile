FROM golang:1.16-alpine AS builder

ARG srcpath="/build/shared-data-provisioner/"

ADD . "$srcpath"

RUN apk add make && make -C "$srcpath" shared-data-provisioner

FROM scratch

COPY --from=builder "$srcpath"/shared-data-provisioner /shared-data-provisioner

CMD ["/shared-data-provisioner"]
