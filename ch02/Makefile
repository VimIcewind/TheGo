GOROOT ?= /opt/go

BIN = boiling ftoc

all : $(BIN)

boiling : boiling.go
	go build boiling.go
ftoc : ftoc.go
	go build ftoc.go

clean :
	rm -rf $(BIN)

