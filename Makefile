.PHONY: test

test: 
	go test -race -count=1 -cover -v ./...

