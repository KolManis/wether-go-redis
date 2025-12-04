.PHONY: run 

# Variables
APP_NAME=server
BIN_DIR=bin
CMD_PATH=cmd/server


## run: Run the application
run:
	@echo "Starting server..."
	go run $(CMD_PATH)/main.go

