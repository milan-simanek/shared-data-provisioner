# the image name
IMAGE?=docker.io/milansimanek/shared-data-provisioner
BINARY=shared-data-provisioner

# the current tag is stored in "./tag" file
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

$(BINARY): $(wildcard *.go)
	CGO_ENABLED=0 GO111MODULE=on go build -a -ldflags '-extldflags "-static"' -o $@ .

clean:
	rm -f shared-data-provisioner 