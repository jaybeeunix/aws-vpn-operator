TESTTARGETS := $(shell go list -e ./... | egrep -v "/(vendor)/")
# ex, -v
TESTOPTS :=

default: gobuild

Gopkg.lock:
	rm -f Gopkg.toml
	dep init

vendor: Gopkg.lock
	dep ensure -v

.PHONY: gocheck
gocheck:
	gofmt -s -w $(shell go list -f '{{ .Dir }}' ./... )
	go vet $(shell go list -f '{{ .Dir }}' ./... )

.PHONY: gotest
gotest:
	go test $(TESTOPTS) $(TESTTARGETS)

.PHONY: gobuild
gobuild: vendor gocheck gotest
	go build $(shell go list -f '{{ .Dir }}' ./... )

