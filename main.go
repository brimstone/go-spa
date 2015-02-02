package main

import (
	"fmt"
	"github.com/brimstone/go-spa/static"
	"github.com/elazarl/go-bindata-assetfs"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.Handle("/",
		http.FileServer(
			&assetfs.AssetFS{Asset: static.Asset, AssetDir: static.AssetDir, Prefix: "data"}))
	http.ListenAndServe(":8080", nil)
}
