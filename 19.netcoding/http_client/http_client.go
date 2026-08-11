package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("http://127.0.0.1:/")
	if err != nil {
		panic(err)
	}
	buf, err := io.ReadAll(res.Body)
	fmt.Println(string(buf))
}
