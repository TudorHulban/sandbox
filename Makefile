test: ## Run tests with check race and coverage.
	@go test -failfast -count=1 ./... -json -cover -race | tparse -smallscreen

lint: ## Run golangci-lint with existing yaml configuration.
	@golangci-lint run