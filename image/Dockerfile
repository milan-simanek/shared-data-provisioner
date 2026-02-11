FROM golang:1.16-alpine AS builder

ARG srcpath="/build/shared-data-provisioner/"

ADD . "$srcpath"

RUN make -C "$srcpath"

FROM scratch

COPY --from=builder /shared-data-provisioner /shared-data-provisioner

CMD ["/shared-data-provisioner"]
