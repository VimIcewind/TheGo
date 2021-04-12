GOROOT ?= /opt/go

BIN = boiling ftoc echo4

all : $(BIN)

boiling : boiling.go
	go build boiling.go
ftoc : ftoc.go
	go build ftoc.go
echo4 : echo4.go
	go build echo4.go

clean :
	rm -rf $(BIN)

