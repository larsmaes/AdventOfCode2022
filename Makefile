BINARIES=$(shell ls ./cmd)
all: $(BINARIES)

$(BINARIES): 
	go build -o ./bin ./cmd/$@

test:
	go test ./...