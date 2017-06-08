/**
	for (续)
	循环初始化语句和后置语句都是可选的。

**/

package main

import "fmt"

func main() {
	sum := 1
	for sum < 10000 {
		sum += sum
	}
	fmt.Println(sum)

}
