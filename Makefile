go-spa: *.go static/build.go
	go fmt
	go build
	${GOPATH}/bin/goupx go-spa

static/build.go: static/data
	cd static; ${GOPATH}/bin/go-bindata -pkg=static -o=build.go -nomemcopy=true data 
