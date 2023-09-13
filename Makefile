.PHONY: build clean clean-all delete

build:
	@echo "Building the app..."
	GOFLAGS=-mod=mod go build -o bin/csvjsonconverter ./cmd

clean:
	@if [ -z "$(file)" ]; then \
		echo "Usage: make clean file=<filename>"; \
	else \
		echo "Removing specified file..."; \
		rm -f $(file); \
	fi

clean-all:
	@echo "Removing all generated CSV and JSON files..."
	@rm -f *.csv *.json

delete:
	@if [ -z "$(type)" ]; then \
		echo "Usage: make delete type=<file_extension>"; \
	else \
		echo "Removing all files with specified extension..."; \
		rm -f *.$(type); \
	fi

test:
	@echo "Running tests..."
	@go test -v ./...