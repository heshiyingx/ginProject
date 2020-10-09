package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

type pxy struct{}

func (p *pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	username, password, _ := req.BasicAuth()
	fmt.Printf("收到请求的数据:\n{method:%v;\n host:%v;\n remoteAddr:%v;\nusername:%v \n;pwd:%v}\n", req.Method, req.Host, req.RemoteAddr, username, password)
	//	step1 浅拷贝对象，然后就再新增属性
	outReq := new(http.Request)
	*outReq = *req

	if req.URL.Host == "www.jd.com" {
		outReq.URL.Host = "www.taobao.com"
		outReq.URL.Scheme = "https"
		outReq.Host = "www.taobao.com"
	}
	all, err := ioutil.ReadAll(req.Body)
	if err == nil {
		outReq.Body = ioutil.NopCloser(bytes.NewReader(all))
	}

	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err != nil {
		if prior, ok := req.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ",") + "," + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}

	//	step2 请求下游
	transport := http.DefaultTransport
	res, err := transport.RoundTrip(outReq)
	if err != nil {
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	defer res.Body.Close()
	//	step3 将请求到的内容给上游
	for key, value := range res.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}

	rw.WriteHeader(res.StatusCode)
	responseData, _ := ioutil.ReadAll(res.Body)
	rw.Write(responseData)

}

func main() {
	err1 := fmt.Errorf("错误1")
	err2 := fmt.Errorf("错误2")
	fmt.Println(err1 == err2)

	http.Handle("/", &pxy{})
	http.ListenAndServe(":9091", nil)
}
