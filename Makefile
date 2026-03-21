build:
	go build -o bin/gendiff ./cmd/gendiff
run:
	bin/gendiff -h
lint:
	golangci-lint run
lint-fix:
	golangci-lint run --fix
test:
	go test ./pkg/... -coverprofile=coverage.out
