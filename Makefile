GO111MODULE=on
GOPROXY=https://goproxy.cn

.PHONY: test

test: 
	go test -race -count=1 -cover -v ./...

