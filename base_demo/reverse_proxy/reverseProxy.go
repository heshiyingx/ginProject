package main

import (
	"bufio"
	"net/http"
	"net/url"
)

const proxy_addr = "https://www.taobao.com"

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9092", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//	1.解析代理地址，并更改请求体的协议和主机
	proxy, err := url.Parse(proxy_addr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.Header.Set("Referer", "http://www.aihuaju.com")
	r.URL.Scheme = proxy.Scheme
	r.URL.Host = proxy.Host
	r.Host = proxy.Host

	//	2.请求下游
	transport := http.DefaultTransport
	response, err := transport.RoundTrip(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for k, v := range response.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	defer response.Body.Close()
	bufio.NewReader(response.Body).WriteTo(w)
}
