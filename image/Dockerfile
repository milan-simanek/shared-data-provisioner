FROM golang:1.16-alpine AS builder

ARG srcpath="/build/shared-data-provisioner/"

ADD . "$srcpath"

RUN cd "$srcpath" && \
    GO111MODULE=on \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -ldflags '-extldflags "-static"' -o /shared-data-provisioner

FROM scratch

COPY --from=builder /shared-data-provisioner /shared-data-provisioner

CMD ["/shared-data-provisioner"]
