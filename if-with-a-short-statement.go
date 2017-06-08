/**
if 的便捷语句
跟 for 一样， if 语句可以在条件之前执行一个简单语句。

由这个语句定义的变量的作用域仅在 if 范围之内。

（在最后的 return 语句处使用 v 看看。）
**/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func min(x, y int) int {
	if x2, y2 := x*x, y*y; x2 < y2 {
		return x
	}
	return y

}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		min(4, 5),
	)
}
