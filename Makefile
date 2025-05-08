.PHONY: build clean docker test

BUILD_DIR=$(shell pwd)/build

build:
	@( \
		cd app || exit 1; \
		go build -o $(BUILD_DIR)/surveillance-proxy main.go; \
	)

clean:
	@rm -rf ${BUILD_DIR} &>/dev/null || true
	@mkdir ${BUILD_DIR}

docker:
	@echo 'not implemented' && exit 1

test:
	@echo 'not implemented' && exit 1
