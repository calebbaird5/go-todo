# Simple Makefile for Go CLI Tool
BINARY=todo
COMPLETIONS=todo_completions
INSTALL_PATH=~/bin
CONFIG_PATH=~/.config/todo

all: install
all: build 

build:
	go build -o $(BINARY)


install: build
	go run main.go completion zsh > _$(BINARY)
	mv $(BINARY) $(INSTALL_PATH)/$(BINARY)
	mkdir -p $(CONFIG_PATH)
	mv _$(BINARY) $(CONFIG_PATH)/$(COMPLETIONS)
	@echo "Generated completions script for zsh: $(CONFIG_PATH)/$(COMPLETIONS)."
	@echo "To enable zsh completions for your CLI, add the following line to your .zshrc:"
	@echo ""
	@echo "    source ~/.config/todo/todo_completions"
	@echo ""

clean:
	rm -f $(BINARY)

.PHONY: all build install clean
