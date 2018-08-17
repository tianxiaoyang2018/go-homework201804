package main

import (
	"net/http"
	"time"
	"fmt"
)

func main() {
	client := http.Client{Timeout: time.Duration(int64(1000000000))}
	resp, err := client.Get("http://127.0.0.1:8081/getlist")
	if err!=nil {
		fmt.Println(err)
	}else {
		byte := make([]byte,resp.ContentLength)
		resp.Body.Read(byte)
		fmt.Println(string(byte))
	}
}
