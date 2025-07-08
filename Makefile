.PHONY: help
help: ## print make commands list 
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## compile the binary
	go build -o ./bin/orichalcum-abacus ./main.go

.PHONY: run
run: ## build and run app.
	go build -o ./tmp/orichalcum-abacus-test ./main.go
	./tmp/orichalcum-abacus-test

.PHONY: test
test: ## run all tests.
	go test -v ./... -count=1

.PHONY: tidy
tidy: ## clean up mod file.
	go mod tidy

.PHONY: act-test
act-test: #Test github actions
	act -j test
