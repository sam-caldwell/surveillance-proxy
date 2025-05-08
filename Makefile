.PHONY: build clean docker test

BUILD_DIR := $(CURDIR)/build
APP_NAME := surveillance-proxy

build:
	@mkdir -p $(BUILD_DIR)
	@echo "Building $(APP_NAME)..."
	@cd app && go build -o $(BUILD_DIR)/$(APP_NAME) main.go
	@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)"

clean:
	@echo "Cleaning build directory..."
	@rm -rf $(BUILD_DIR)
	@mkdir -p $(BUILD_DIR)

docker:
	@docker build -t $(APP_NAME):latest .

test:
	@echo 'not implemented' && exit 1
