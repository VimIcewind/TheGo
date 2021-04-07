GOROOT ?= /opt/go

BIN = helloworld echo1 echo2 echo3 args

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

clean :
	rm -rf $(BIN)

