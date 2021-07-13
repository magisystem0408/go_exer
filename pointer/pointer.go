package main

import "fmt"

	//アドレスを渡す
func one(x *int) {
	//実体の中に入れる
	*x=1

}

func main() {
	var n int = 100

	one(&n)
	fmt.Println(n)



	fmt.Println(n)

	fmt.Println(&n)


	var p *int = &n
	fmt.Println(p)



	fmt.Println(*p)


}