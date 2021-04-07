GOROOT ?= /opt/go

BIN = helloworld echo1 echo2 echo3 args 1.1

all : $(BIN)

helloworld : helloworld.go
	go build helloworld.go
echo1 : echo1.go
	go build echo1.go
echo2 : echo2.go
	go build echo2.go
echo3 : echo3.go
	go build echo3.go
args : args.go
	go build args.go
1.1 : 1.1.go
	go build 1.1.go

clean :
	rm -rf $(BIN)

