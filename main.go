package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/elazarl/goproxy"
)

type myProxyWrapper struct {
	proxy http.Handler
}

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false
	proxy.OnResponse().DoFunc(
		func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			r.Header.Set("X-Proxy", "fvin_vf_pbby")
			log.Println(r.Request.Method, r.Request.URL)
			return r
		})

	// address := ":" + os.Args[1]
	address := ":8080"
	fmt.Println("starting http proxy at ", address)
	log.Fatal(http.ListenAndServe(address, proxy))
}
