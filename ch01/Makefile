GOROOT ?= /opt/go

BIN = helloworld echo1

all : $(BIN)

helloworld : helloworld.go
	go build helloworld.go
echo1 : echo1.go
	go build echo1.go

clean :
	rm -rf $(BIN)


