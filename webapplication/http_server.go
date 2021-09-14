package main

import (
	"html/template"
	"log"
	"net/http"
)


//type MyHandler struct {
//
//}
//
//func (h *MyHandler) ServeHTTP (w http.ResponseWriter,r *http.Request){
//	fmt.Fprintln(w,"hello world")
//}
func top(w http.ResponseWriter, r *http.Request)  {

	//htmlファイルを用意しておく
	//htmlファイルを解析する
	t,err :=template.ParseFiles("templ.html")
	if err !=nil{
		log.Println(err)
	}

	//ファイルを実行する
	t.Execute(w,"hello world1111")

}


func main() {
	http.HandleFunc("/top",top)
	http.ListenAndServe(":8080",nil)
}