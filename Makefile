# Simple Makefile for Go CLI Tool
BINARY=todo
INSTALL_PATH=~/bin

all: install
all: build 

build:
	go build -o $(BINARY)

install: build
	mv $(BINARY) $(INSTALL_PATH)/$(BINARY)

clean:
	rm -f $(BINARY)

.PHONY: all build install clean
