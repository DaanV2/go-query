

reports:
	mkdir -p reports

build:
	go build ./...

coverage: test
	go tool cover -html=reports/coverage

test:
	go test -v -coverprofile=reports/coverage ./...
