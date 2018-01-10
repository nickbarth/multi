package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type HostMap struct {
	host   string
	target string
}

type MultiProxy struct {
	addr  string
	hosts map[string]*httputil.ReverseProxy
}

func NewMultiProxy(addr string) *MultiProxy {
	return &MultiProxy{
		addr:  addr,
		hosts: make(map[string]*httputil.ReverseProxy),
	}
}

func (t *MultiProxy) AddHost(p HostMap) {
	fmt.Printf("`%v` :: `%v`\n", p.host, p.target)
	server, _ := url.Parse(p.target)
	t.hosts[p.host] = httputil.NewSingleHostReverseProxy(&url.URL{Scheme: server.Scheme, Host: server.Host})
}

func (t MultiProxy) Run() error {
	err := http.ListenAndServe(t.addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, exists := t.hosts[r.Host]; exists {
			t.hosts[r.Host].ServeHTTP(w, r)
		}
	}))

	return err
}

func main() {
	fmt.Println("MultiProxy Started.")
	proxy := NewMultiProxy(":8000")
	proxy.AddHost(HostMap{host: "example.com", target: "http://localhost:9000"})
	proxy.AddHost(HostMap{host: "example.org", target: "http://localhost:9001"})
	err := proxy.Run()
	if err != nil {
		log.Fatal(err)
	}
}
