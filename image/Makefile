#
IMAGE?=docker.io/milansimanek/shared-data-provisioner

TAG:=$(strip $(shell head -n 1 tag))
TAG_GIT=$(IMAGE):$(TAG)
TAG_LATEST=$(IMAGE):latest

PHONY: all
all: image

PHONY: image
image:
	docker build -t $(TAG_GIT) -f Dockerfile .
	docker tag $(TAG_GIT) $(TAG_LATEST)

PHONY: push
push: image
	docker push $(TAG_GIT)
	docker push $(TAG_LATEST)

PHONY: shared-data-provisioner
shared-data-provisioner: export CGO_ENABLED=0
shared-data-provisioner: export GO111MODULE=on
shared-data-provisioner: $(shell find . -name "*.go")
	go build -a -ldflags '-extldflags "-static"' -o shared-data-provisioner .
