package statement

import "fmt"

func main() {

	//ifぶん
	num := 4
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {

	} else {
		fmt.Println("else")
	}

	x, y := 11, 12
	if x == 10 && y == 10 {
		fmt.Println("ff")
	}

	//for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//省略記法
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

}
