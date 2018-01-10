# MultiProxy

A small multi-domain reverse proxy in Go.

For when Apache/Nginx is too much.

### Usage

```go
func main() {
	fmt.Println("MultiProxy Started.")
	proxy := multi.NewMultiProxy(":8000")
	proxy.AddHost(HostMap{host: "example.com", target: "http://localhost:9000"})
	proxy.AddHost(HostMap{host: "example.org", target: "http://localhost:9001"})
	err := proxy.Run()
	if err != nil {
		log.Fatal(err)
	}
}
```

### License
WTFPL &copy; 2016 Nick Barth
