APPDIR=/home/root
APPNAME=go-rest
GOPATH=/go
PWD=$(shell pwd)

docker: 
	docker run -it -v ${PWD}:${APPDIR} -v ${PWD}/pkg:${GOPATH}/pkg --name ${APPNAME} --hostname ${APPNAME} --rm  -w ${APPDIR} golang:1.13 ${CMD}

test:
	@$(MAKE) CMD="go test -race" docker

sh:
	@$(MAKE) docker