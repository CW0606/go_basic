/**
练习：斐波纳契闭包
现在来通过函数做些有趣的事情。

实现一个 fibonacci 函数，返回一个函数（一个闭包）可以返回连续的斐波纳契数。
**/

package main

import (
	"fmt"
)

func fibonacci() func() int {
	sum_zero := 0
	sum_one := 1
	return func() int {
		cur := sum_zero + sum_one
		sum_zero = sum_one
		sum_one = cur
		return sum_zero
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Println(f())
	}
}
