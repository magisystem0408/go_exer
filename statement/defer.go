package main

//deferはmain関数の処理が終わったら実行される
import "fmt"

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("foo")
}

func main() {

	foo()
	defer fmt.Println("world")
	fmt.Println("hollo")

	//実行結果
	//foo
	//world foo
	//hollo
	//world



	//スタッキングdefer最初に入れたものは最後になる
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")



}
