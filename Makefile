#
IMAGE?=docker.io/milansimanek/shared-data-provisioner

TAG_GIT=$(IMAGE):v0.1.2
TAG_LATEST=$(IMAGE):latest

PHONY: test-image
test-image:
	docker build -t shared-data-provisioner -f Dockerfile .

PHONY: all
all: image push

PHONY: image
image:
	docker build -t $(TAG_GIT) -f Dockerfile .
	docker tag $(TAG_GIT) $(TAG_LATEST)

PHONY: push
push:
	docker push $(TAG_GIT)
	docker push $(TAG_LATEST)

PHONY: shared-data-provisioner
shared-data-provisioner: export CGO_ENABLED=0
shared-data-provisioner: export GO111MODULE=on
shared-data-provisioner: $(shell find . -name "*.go")
	go build -a -ldflags '-extldflags "-static"' -o shared-data-provisioner .
