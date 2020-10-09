package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

var addr = "http://127.0.0.1:8083"

func main() {
	parse, err := url.Parse(addr)
	if err != nil {
		panic(err)
	}

	ar := make([]string, 0, 100)
	ar = append(ar, "select * from")

	proxy := httputil.NewSingleHostReverseProxy(parse)

	//proxy.Director = func(request *http.Request) {
	//	if request.Host == "127.0.0.1:9092" {
	//		url, _ := http.NewRequest(http.MethodGet, "https://www.baidu.com", nil)
	//		request.URL = url.URL
	//		request.Host = url.Host
	//	}

	//}
	proxy.ModifyResponse = func(response *http.Response) error {

		return nil
	}
	http.ListenAndServe(":9092", proxy)
}
