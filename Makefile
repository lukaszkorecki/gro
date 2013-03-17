main = gro.go
build:
	go build -v $(main)

edit:
	env GOPATH=$(gopath) vim $*
