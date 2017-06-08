// 程序包名
package main

import "fmt"

const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
)

func main() {
	fmt.Println(PB)
}
