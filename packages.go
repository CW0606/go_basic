/**
包
每个 Go 程序都是由包组成的。

程序运行的入口是包 main 。

这个程序使用并导入了包 "fmt" 和 "math/rand" 。

按照惯例，包名与导入路径的最后一个目录一致。例如，"math/rand" 包由 package rand 语句开始。

注意：这个程序的运行环境是确定性的，因此 rand.Intn 每次都会返回相同的数字。 （为了得到不同的随机数，需要提供一个随机数种子，参阅 rand.Seed。）
**/

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My Faorite number is", rand.Intn(100))
	fmt.Println("My Faorite number is", rand.Intn(100))
	fmt.Println("My Faorite number is", rand.Intn(100))
	fmt.Println("My Faorite number is", rand.Intn(100))
}
