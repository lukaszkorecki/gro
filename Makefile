gopath=.
gobin=$(gopath)/bin
main = src/main.go
build:
	env GOPATH=$(gopath) GOBIN=$(gobin) go build $(main)

edit:
	env GOPATH=$(gopath) vim $*
