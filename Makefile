APPDIR=/go/src/github.com/diegodesousas/go-rest
APPNAME=go-rest
IMAGE=go-currency-converter
GOPATH=/go
PWD=$(shell pwd)
PORT=9000

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
		${IMAGE} \
		${CMD} \

test:
	@$(MAKE) CMD="go test -race ./..." docker

sh:
	@$(MAKE) docker

run:
	@$(MAKE) CMD="dogo" docker

build:
	 docker build -t ${IMAGE} .