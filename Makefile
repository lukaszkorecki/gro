GOPATH=`pwd`
main = src/gro.go
test_packages="fileReader"

test:
	@env GOPATH=$(GOPATH) go test -v $(test_packages)

build:
	@echo Building in $(GOPATH)
	@env GOPATH=$(GOPATH) go build -v $(main)
