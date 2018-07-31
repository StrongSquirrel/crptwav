test:
	go test -race -timeout 30s `go list ./... | grep -v /vendor/`
