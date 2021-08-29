.PHONY: execute

build: cmd/main.go
	cd cmd && go build -o ../bin/pisang 

execute:
	cd bin && ./pisang sample.pi

buildexec: build execute