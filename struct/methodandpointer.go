package _struct

import "fmt"

type Vertex struct {
	X, Y int
}

//上のvertexに紐付けされる
func (v Vertex) Area() int {
	return v.X * v.Y
}

//structの値を変えたい時は、アスタリスクをつける
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func Area(v Vertex) int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))

	v.Scale(10)
	//structを継承しているので使用できる
	fmt.Println(v.Area())
	v.Area()
}
