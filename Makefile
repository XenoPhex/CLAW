NODES ?= 4
GOSRC = $(shell find . -name "*.go" ! -name "*test.go" ! -name "*fake*")

all : build test

build : claw

claw : $(GOSRC) format
	go build .

format :
	go fmt ./...

test : vet
	ginkgo -nodes $(NODES) -randomizeAllSpecs -randomizeSuites ./...

vet :
	@echo  "Vetting packages for potential issues..."
	go tool vet -all -shadow=true ./request ./exec
	@echo

.PHONY : all build format test vet
