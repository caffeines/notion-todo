build:
	@go build -o todo main.go
	@echo "Build complete."

build-all:
	@echo "Building for all platforms..."
	@GOOS=darwin GOARCH=amd64 go build -o bin/notion-todo-darwin-amd64 main.go
	@GOOS=darwin GOARCH=arm64 go build -o bin/notion-todo-darwin-arm64 main.go
	@GOOS=linux GOARCH=amd64 go build -o bin/notion-todo-linux-amd64 main.go
	@GOOS=windows GOARCH=amd64 go build -o bin/notion-todo-windows-amd64.exe main.go
	@echo "Cross-compilation complete. Binaries in ./bin/"

clean:
	@rm -f todo
	@rm -rf bin/
	@echo "Clean complete."

release-test:
	@goreleaser release --snapshot --clean

.PHONY: build build-all clean release-test