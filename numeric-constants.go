/**
数值常量
数值常量是高精度的 值 。

一个未指定类型的常量由上下文来决定其类型。

也尝试一下输出 needInt(Big) 吧。

（int 可以存放最大64位的整数，根据平台不同有时会更少。）
**/

package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 1.0
}

func main() {
	// x, y := 1, 1.0
	// z := x * y
	// fmt.Println(z)
	//
	//
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println(needInt(Big))
}
