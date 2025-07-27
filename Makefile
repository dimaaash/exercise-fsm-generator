# Makefile for exercise-fsm-generator



BINARY_NAME=exercise-fsm-generator
CMD_DIR=cmd/mod3
BIN_DIR=$(CMD_DIR)/..


.PHONY: all build clean run test help


all: build
help:
	@echo "exercise-fsm-generator Makefile Help"
	@echo "--------------------------------------"
	@echo "make build   - Build the mod3 binary (cmd/exercise-fsm-generator)"
	@echo "make run     - Run the mod3 application (default input: none)"
	@echo "make test    - Run unit tests for all packages"
	@echo "make clean   - Remove the built binary"
	@echo ""
	@echo "To run the mod3 application manually:"
	@echo "  ./cmd/exercise-fsm-generator <16-bit binary string>"
	@echo "Example: ./cmd/exercise-fsm-generator 1011001100110011"
	@echo ""
	@echo "Or use the Makefile to run with input (quotes required):"
	@echo "  make run ARGS=\"1011001100110011\""
	@echo ""
	@echo "The input must be a 16-character string of 1s and 0s."

build:
	go build -o $(BIN_DIR)/$(BINARY_NAME) ./$(CMD_DIR)

run: build
	$(BIN_DIR)/$(BINARY_NAME) $(ARGS)

test:
	go test ./...

clean:
	rm -f $(BIN_DIR)/$(BINARY_NAME)
