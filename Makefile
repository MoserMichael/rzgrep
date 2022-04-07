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
	mv rzgrep rzgrep
	tar cvfz rzgrep-$(shell uname -s).tar.gz rzgrep

allrel : rel
	docker build --progress=plain --rm=true -t rzgrep-builder:latest . 
	docker create -ti --name dummy rzgrep-builder:latest
	docker cp dummy:/go/rzgrep-Linux.tar.gz .
	docker rm -f dummy
