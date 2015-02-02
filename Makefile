go-spa: *.go
	go fmt
	go build
	$GOPATH/bin/goupx go-spa
