/**
switch 的执行顺序
switch 的条件从上到下的执行，当匹配成功的时候停止。

（例如，

switch i {
case 0:
case f():
}
当 i==0 时不会调用 f 。）

**/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	case today + 3:
		fmt.Println("In three days")
	default:
		fmt.Println("Too far away")
	}
}
