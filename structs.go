/**
结构体
一个结构体（ struct ）就是一个字段的集合。

（而 type 的含义跟其字面意思相符。）
**/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

}
