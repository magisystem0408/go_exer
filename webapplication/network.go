package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {

	//html読み込み
	//res,_ :=http.Get("http://google.com")
	//defer res.Body.Close()
	//
	//body,_ :=ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))


	 //URLの生成
	//urlパースをする
	base,_ :=url.Parse("http://example.com")

	//↑の後につけるパスを指定できる
	reference,_ :=url.Parse("/test?a=1&b=2")

	//エンドポイント指定する。
	endpoint:=base.ResolveReference(reference).String()

	fmt.Println(endpoint)

	//リクエストするときは、nilはpostの時はバイト入れないといけない
	req,_ :=http.NewRequest("GET",endpoint,nil)
	//ヘッダー情報
	req.Header.Add("If-None-Match",`W/wyzy`)
}