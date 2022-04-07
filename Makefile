.PHONY: main
main:
	GOPATH=$(PWD) GO111MODULE=off go build -o rzgrep cmd/rzgrep/main.go

.PHONY: vet
vet:
	GOPATH=$(PWD) GO111MODULE=off go vet ./cmd/... ./src/...

.PHONY: clean
clean:
	rm -f rzgrep

rel: clean main
	tar cvfz rzgrep.tar.gz rzgrep
