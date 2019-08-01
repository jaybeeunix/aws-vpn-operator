TESTTARGETS := $(shell go list -e ./... | egrep -v "/(vendor)/")
# ex, -v
TESTOPTS :=

GO111MODULE?=on

default: gobuild

.PHONY: gocheck
gocheck:
	gofmt -s -w $(shell go list -f '{{ .Dir }}' ./... )
	go vet $(shell go list -f '{{ .Dir }}' ./... )

.PHONY: gotest
gotest:
	go test $(TESTOPTS) $(TESTTARGETS)

.PHONY: gobuild
gobuild: gocheck gotest
	go build $(shell go list -f '{{ .Dir }}' ./... )

