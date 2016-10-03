package main

import(
    "log"
    "net/url"
    "net/http"
    "net/http/httputil"
)

func handler(p map[string]*httputil.ReverseProxy) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if _, exists := p[r.Host]; exists {
            p[r.Host].ServeHTTP(w, r)
        }
    }
}

func main() {
    proxies := map[string]*httputil.ReverseProxy{}
    proxies["example.com"] = httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: "localhost:9000"})
    proxies["example.org"] = httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: "localhost:9001"})

    log.Fatal(http.ListenAndServe(":80", handler(proxies)))
}
