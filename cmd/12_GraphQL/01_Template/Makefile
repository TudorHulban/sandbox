reset:
	@rm -rf .vscode graph vendor
	@rm -rf server.go go.sum go.mod tools.go gqlgen.yml
	@printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
	@go mod init testq
	@go get github.com/99designs/gqlgen

init:
	go run github.com/99designs/gqlgen init

generate:
	@echo "Removing all '*resolver*.go' files in graph folder."
	@rm -f graph/*resolver*.go
	@echo "Generating new files."
	@go run github.com/99designs/gqlgen
	@echo "Generation was succesfull."

run:
	@go run ./server.go

test:
	go test -v ./...