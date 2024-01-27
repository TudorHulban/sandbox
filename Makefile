test: ## Run tests with check race and coverage.
	@go test -failfast -count=1 ./... -json -cover -race | tparse -smallscreen

lint: ## Run golangci-lint with existing yaml configuration.
	@golangci-lint run

pg-local:
	@docker run -d --name=co-postgres -p 5432:5432 \
	-e POSTGRES_PASSWORD=password \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_DB=sandbox \
	postgres
