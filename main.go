package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	url := "http://e1656207s5.iok.la:5555/containers/e480af73ee6c59508d0d0073f04f93634a09b3c3fddf76be661edb88ab9a36cb/stats"
	HttpGet(url)
}

func HttpGet(uri string) {
	res, err := http.Get(uri)
	if err != nil {
		fmt.Println(err)
	}
	//buf := make([]byte, 10240)

	for {
		bufReader := bufio.NewReader(res.Body)
		bodyBuf, isP, err := bufReader.ReadLine()
		fmt.Println(string(bodyBuf), isP, err)
		if err != nil {
			break
		}
	}
}
