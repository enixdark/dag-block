#--------------------------------------
# Stage: Building dag block project
#--------------------------------------
FROM golang:alpine

ENV GOPATH /usr/local/go/

RUN apk update && apk add curl

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mkdir -p /usr/local/go/src/github.com/enixdark/dag-block
ADD . /usr/local/go/src/github.com/enixdark/dag-block
WORKDIR /usr/local/go/src/github.com/enixdark/dag-block

RUN dep ensure

RUN go build -o dag_block


CMD ["./dag_block"]
