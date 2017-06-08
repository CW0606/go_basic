/**
slice 切片
一个 slice 会指向一个序列的值，并且包含了长度信息。

[]T 是一个元素类型为 T 的 slice。

len(s) 返回 slice s 的长度。
**/

package main

import "fmt"

func main() {
	s := []int{2, 3, 4, 5, 6, 11, 23}
	fmt.Println("s == ", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}
