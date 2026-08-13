FROM golang:1.22-alpine AS builder

RUN apk add make 

ADD . /build

RUN make -C /build shared-data-provisioner

FROM scratch

COPY --from=builder /build/shared-data-provisioner /shared-data-provisioner

CMD ["/shared-data-provisioner"]
