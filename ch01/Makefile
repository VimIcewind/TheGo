GOROOT ?= /opt/go

BIN = helloworld echo1 echo2

all : $(BIN)

helloworld : helloworld.go
	go build helloworld.go
echo1 : echo1.go
	go build echo1.go
echo2 : echo2.go
	go build echo2.go

clean :
	rm -rf $(BIN)

