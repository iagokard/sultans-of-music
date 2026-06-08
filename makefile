BINARY_NAME=som
MAIN_FILE=main.go

.PHONY: run build generate-docs

## run: Run the application
run: generate-docs
	go run $(MAIN_FILE)

## build: Build the application
build: generate-docs
	go build -o $(BINARY_NAME) $(MAIN_FILE)

## generate-docs: Generate Swagger documentation
generate-docs:
	swag init -g $(MAIN_FILE) --parseDependency --parseInternal
