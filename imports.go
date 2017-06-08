/**
导入
这个代码用圆括号组合了导入，这是“打包”导入语句。

同样可以编写多个导入语句，例如：

import "fmt"
import "math"
不过使用打包的导入语句是更好的形式。
**/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.", math.Sqrt(8))
}
