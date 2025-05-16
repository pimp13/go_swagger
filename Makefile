MAIN_FILE = ./main.go
BUILD_FILE = ./tmp/main.exe


.PHONY: dev swag build start swag-fmt

dev:
	@echo "Running server in development mode"
	@air -c .air.toml

swag:
	@echo "Creating swagger document..."
	@swag init -g main.go --output ./docs

build: swag swag-fmt
	@echo "Building application..."
	@go build -o $(BUILD_FILE) $(MAIN_FILE)

start: build
	@echo "Running server in production mode"
	@$(BUILD_FILE)

swag-fmt:
	@echo "Formatting swagger docs"
	@swag fmt

clean:
	@rm -rf ./docs
	@rm -rf ./tmp