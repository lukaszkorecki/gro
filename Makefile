GOPATH=`pwd`
main = src/gro.go
build:
	@echo Building in $(GOPATH)
	@env GOPATH=$(GOPATH) go build -v $(main)
