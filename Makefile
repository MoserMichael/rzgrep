.PHONY: main
main: makego makeJava
	echo "*main build*"

.PHONY: makego
makego:
	GOPATH=$(PWD) GO111MODULE=off go build -o rzgrep cmd/rzgrep/main.go

.PHONY: makeJava
makeJava: 
	cd java-decompiler; ./gradlew build
	cp ./java-decompiler/build/libs/java-decompiler-1.0-SNAPSHOT.jar rzgrep.jar

.PHONY: vet
vet:
	GOPATH=$(PWD) GO111MODULE=off go vet ./cmd/... ./src/...

.PHONY: clean
clean:
	rm -f rzgrep rzgrep-*.tar.gz

rel: clean makego
	mv rzgrep rzgrep
	tar cvfz rzgrep-$(shell uname -s).tar.gz rzgrep rzgrep.jar

allrel : rel
	docker build --progress=plain --rm=true -t rzgrep-builder:latest . 
	docker create -ti --name dummy rzgrep-builder:latest
	docker cp dummy:/go/rzgrep-Linux.tar.gz .
	docker rm -f dummy
