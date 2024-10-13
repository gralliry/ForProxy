package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	listenHost := flag.String("h", "0.0.0.0", "")
	listenPort := flag.String("p", "8080", "")
	targetUrls := flag.String("t", "http://example.com", "")
	flag.Parse()

	targetURL, err := url.Parse(*targetUrls)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 设置URL
		r.URL.Scheme = targetURL.Scheme
		r.URL.Host = targetURL.Host
		// 设置header里面的
		r.Host = targetURL.Host
		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.ServeHTTP(w, r)
	})

	log.Printf("Starting reverse proxy server on %s:%s", *listenHost, *listenPort)
	err = http.ListenAndServe(*listenHost+":"+*listenPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
