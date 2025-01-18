golint:
	golangci-lint run -c .golangci.yaml --issues-exit-code 0
.PHONY:golint