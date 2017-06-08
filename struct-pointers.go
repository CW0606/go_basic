/**
	结构体指针
结构体字段可以通过结构体指针来访问。

通过指针间接的访问是透明的。
**/

package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{2, 3}
	p := &v
	p.X = 100
	fmt.Println(v)
}
