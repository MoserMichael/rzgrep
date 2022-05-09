.PHONY: main
main: makego makeJava
	echo "*main build*"

.PHONY: makego
makego:
	GOPATH=$(PWD) GO111MODULE=off go build -o rzgrep cmd/rzgrep/main.go

.PHONY: makego-arch
makego-arch:
	GOPATH=$(PWD) GO111MODULE=off GOARCH=$(GOARCH) GOOS=$(GOOS) go build -o rzgrep-${GOOS}-${GOARCH} cmd/rzgrep/main.go
	tar cvfz rzgrep-${GOOS}-${GOARCH}.tar.gz rzgrep-${GOOS}-${GOARCH} rzgrep.jar

.PHONY: makego-all
makego-all: makeJava
	GOOS=darwin GOARCH=arm64 make makego-arch
	GOOS=darwin GOARCH=amd64 make makego-arch
	GOOS=linux  GOARCH=arm64 make makego-arch
	GOOS=linux  GOARCH=amd64 make makego-arch
	echo "*** all architectures compiled ***"
	
.PHONY: makeJava
makeJava: 
	cd java-decompiler; ./gradlew build
	cp ./java-decompiler/build/libs/java-decompiler-1.0-SNAPSHOT.jar rzgrep.jar

.PHONY: vet
vet:
	GOPATH=$(PWD) GO111MODULE=off go vet ./cmd/... ./src/...

.PHONY: clean
clean:
	rm -f rzgrep*

#rel: clean makego
#	mv rzgrep rzgrep
#	tar cvfz rzgrep-$(shell uname -s)-$(shell uname -m).tar.gz rzgrep rzgrep.jar

#allrel : rel
#	docker build --progress=plain --rm=true -t rzgrep-builder:latest . 
#	docker create -ti --name dummy rzgrep-builder:latest
#	docker cp dummy:/go/rzgrep-Linux.tar.gz .
#	docker rm -f dummy
