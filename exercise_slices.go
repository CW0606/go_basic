/**练习：slice
实现 Pic 。它返回一个长度为 dy 的 slice，其中每个元素是一个长度为 dx 且元素类型为8位无符号整数的 slice。当你运行这个程序时， 它会将每个整数作为对应像素的灰度值（好吧，其实是蓝度）并显示这个 slice 所对应的图像。

计算每个像素的灰度值的方法由你决定；几个有意思的选择包括 (x+y)/2、x*y 和 x^y 。

（需要使用循环来分配 [][]uint8 中的每个 []uint8 。）

（使用 uint8(intValue) 来在类型之间进行转换。）
**/
package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		slice_x = &pic[y]
		*slice_x = make([]uint8, dx)
		for x := range *slice_x {
			pic[y][x] = unint8((x + y) * x * x / y * y)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
