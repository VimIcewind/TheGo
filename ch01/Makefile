GOROOT ?= /opt/go

BIN = helloworld echo1 echo2 echo3 args 1.1 1.2 1.3 dup1 dup2 dup3 1.4 lissajous \
      1.5 1.6

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
1.2 : 1.2.go
	go build 1.2.go
1.3 : 1.3.go
	go build 1.3.go
dup1 : dup1.go
	go build dup1.go
dup2 : dup2.go
	go build dup2.go
dup3 : dup3.go
	go build dup3.go
1.4 : 1.4.go
	go build 1.4.go
lissajous : lissajous.go
	go build lissajous.go
1.5 : 1.5.go
	go build 1.5.go
1.6 : 1.6.go
	go build 1.6.go

clean :
	rm -rf $(BIN) *.gif

