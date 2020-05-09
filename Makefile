GEN_OUT = elements.go attributes/attributes.go
GEN_SRC = $(shell find htmlgogen -type f -name '*.go')
LIB_SRC = $(shell find . -type f -name '*.go' ! -wholename './htmlgogen*')

.PHONY: all
all: $(GEN_OUT) $(LIB_SRC)
	go build .

$(GEN_OUT) &: $(GEN_SRC)
	go generate .
