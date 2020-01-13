APPDIR=/home/root
PWD=$(shell pwd)

test:
	docker run -v ${PWD}:${APPDIR} --name go-modules --hostname go-modules --rm  -w ${APPDIR} golang:1.13 go test -race

sh:
	docker run -it -v ${PWD}:${APPDIR} --name go-modules --hostname go-modules --rm  -w ${APPDIR} golang:1.13
