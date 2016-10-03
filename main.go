package main

import (
    "bufio"
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
    "strings"
)

func handler(p map[string]*httputil.ReverseProxy) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if _, exists := p[r.Host]; exists {
            p[r.Host].ServeHTTP(w, r)
        } else {
            fmt.Fprintf(w, "Not Found.\n")
        }
    }
}

func main() {
    proxies := map[string]*httputil.ReverseProxy{}

    file, err := os.Open("router.config")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        config := strings.Split(line, "|")
        port, host := config[0], config[1]

        remote, err := url.Parse(host)
        if err != nil {
            log.Fatal(err)
        }

        proxies[remote.Host] = httputil.NewSingleHostReverseProxy(&url.URL{Scheme: remote.Scheme, Host: "localhost:" + port})
        fmt.Println(port + " " + remote.Scheme + " " + remote.Host)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    log.Fatal(http.ListenAndServe(":80", handler(proxies)))
}
