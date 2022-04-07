FROM golang:1.18.0-alpine3.15 AS rzgrepbuilder

RUN apk --no-cache add make

WORKDIR /workspace
RUN ls -al
COPY . .
RUN pwd
RUN ls -al
RUN make rel

FROM golang:1.18.0-alpine3.15 as rzgreprunner
COPY --from=rzgrepbuilder /workspace/rzgrep-Linux.tar.gz .


