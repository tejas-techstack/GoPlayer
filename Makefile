GO=go
SRC_DIR=./src
BUILD_DIR=./bin
OUTPUT=$(BUILD_DIR)/Player

build:
	@$(GO) build -o $(OUTPUT) $(SRC_DIR)

run: build
	@$(OUTPUT)

