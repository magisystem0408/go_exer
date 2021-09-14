package _struct

import "fmt"

type MyInt int

func (i MyInt) Double() int {
fmt.Printf("%T %v", i, i)
fmt.Printf("%T %v", 1, 1)

//返り値がintなのでintでキャストしないといけない
return int(i % 2)
}

func main() {
myInt :=MyInt(10)
fmt.Println(myInt.Double())
}
