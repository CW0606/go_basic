/**
map 的文法
map 的文法跟结构体文法相似，不过必须有键名。

**/

package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := map[string]Vertex{
		"Bell": {
			40.0, 50.0,
		},
		"Google": {
			44.0, 50.0,
		},
	}

	fmt.Println(m)
}
