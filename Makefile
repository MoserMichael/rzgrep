.PHONY: main
main:
	GOPATH=$(PWD) GO111MODULE=off go build -o rzgrep cmd/rzgrep/main.go

.PHONY: clean
clean:
	rm -f rzgrep
