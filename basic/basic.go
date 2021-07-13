package main

import "fmt"
import "strings"

func main() {
	var (
		u8 uint8=255
		i8 int8 =127
		f32 float32 =0.2
		c64 complex64 =-5+12i
	)
	fmt.Println(u8,i8,f32,c64)

	fmt.Printf("type=%T value=%v",u8,u8)
	//実行結果
	//type=uint8 value=255

	//変数の型宣言
	x:=0
	fmt.Println(x)



	//算術シフト
	fmt.Println(1<<0)
	fmt.Println(1<<1)
	fmt.Println(1<<2)
	fmt.Println(1<<3)
	//実行結果
	//1248

	//文字列の指定
	fmt.Println(string("hello would"[0]))
	//実行結果
	//H



	var s string ="Hello World"
	//文字列の置き換え
	//fmt.Println(strings.Replace(s,"H","X",1))
	s =strings.Replace(s,"H","H",1)

	//文字列の中に入っているかのTrue or false
	fmt.Println(strings.Contains(s,"World"))

}