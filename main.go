package main

import (
	"encoding/json"
	"fmt"
	"github.com/brimstone/go-spa/static"
	"github.com/elazarl/go-bindata-assetfs"
	"net/http"
)

type Response map[string]interface{}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, Response{"args": r.URL.Path[1:]})
}

func main() {
	http.Handle("/",
		http.FileServer(
			&assetfs.AssetFS{Asset: static.Asset, AssetDir: static.AssetDir, Prefix: "data"}))
	http.HandleFunc("/api/", handler)

	http.ListenAndServe(":8080", nil)
}
