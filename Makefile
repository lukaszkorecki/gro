GOPATH=`pwd`
main = src/gro.go
# find all package names in src and add them to list
test_packages=`find ./src/ -type d | sed 's/^.*src.*\///'`

test:
	echo $(test_packages)
	env GOPATH=$(GOPATH) go test -v $(test_packages)

build:
	@echo Building in $(GOPATH)
	@env GOPATH=$(GOPATH) go build -v $(main)
