PROJECT="todoListApp"
VERSION="v1.0.0"
default:
	@echo Building ${PROJECT} app, you need to run mysql container first\\n *sql: run mysql container \\n *build: build golang binary app \\n *exec: run app
build:
	@go build -o ${PROJECT}-${VERSION} -tags=Funskie
exec:
	@./${PROJECT}-${VERSION}
clean:
	@if [ -f ${PROJECT}-${VERSION} ] ; then rm ${PROJECT}-${VERSION} ; fi && docker stop mysql && docker rm mysql
# .PHONY: default sql build
