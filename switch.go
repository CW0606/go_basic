/**

switch
你可能已经知道 switch 语句会长什么样了。

除非以 fallthrough 语句结束，否则分支会自动终止。
**/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS x.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s", os)
	}
}
