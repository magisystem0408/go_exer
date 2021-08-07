package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res,_ :=http.Get("http://google.com")
	defer res.Body.Close()

	body,_ :=ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}