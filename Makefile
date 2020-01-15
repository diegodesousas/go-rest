APPDIR=/go/src/github.com/diegodesousas/go-rest
APPNAME=go-rest
GOPATH=/go
PWD=$(shell pwd)
PORT=9000
GOVERSION=1.13

docker: 
	docker run \
		-it \
		-v ${PWD}:${APPDIR} \
		-v ${PWD}/.pkg:${GOPATH}/pkg \
		-p ${PORT}:${PORT} \
		--name ${APPNAME} \
		--hostname ${APPNAME} \
		--rm  \
		-w ${APPDIR} \
		golang:${GOVERSION} \
		${CMD}

test:
	@$(MAKE) CMD="go test -race ./..." docker

sh:
	@$(MAKE) docker

run:
	@$(MAKE) CMD="go run main.go" docker