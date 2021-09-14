package statement

import "fmt"

//帰り値二つの場合
func add(x, y int) (int, int) {
	return x + y, x - y
}

//帰り値を指定した場合はreturnを描かなくても良い
func cal(price, item int) (result int) {
	result = price * item
	return
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 :=cal(100,2)
	fmt.Println(r3)


	//関数内関数
	f := func() {
		fmt.Println("inner func")
	}
	f()

	//関数の直後に書いて実行するもある
	func(x int){
		fmt.Println("inner func",x)
	}(1)

}
