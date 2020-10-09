package httpdemo

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func client() {
	transport := &http.Transport{
		Proxy: nil,
		DialContext: (&net.Dialer{
			Timeout:       30 * time.Second,
			LocalAddr:     nil,
			FallbackDelay: 0,
			KeepAlive:     30 * time.Second,
			Resolver:      nil,
			Control:       nil,
		}).DialContext,

		MaxIdleConns:          1000,
		MaxIdleConnsPerHost:   1000,
		MaxConnsPerHost:       1000,
		IdleConnTimeout:       3 * time.Second,
		ExpectContinueTimeout: 3 * time.Second,
	}

	client := http.Client{
		Transport: transport,
		Timeout:   time.Second * 30,
	}

	response, err := client.Get("http://127.0.0.1:1210/bye")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(bytes))

}
