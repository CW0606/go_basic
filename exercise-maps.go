/**
	实现 WordCount。它应当返回一个含有 s 中每个 “词” 个数的 map。函数 wc.Test 针对这个函数执行一个测试用例，并输出成功还是失败。

你会发现 strings.Fields 很有帮助。
**/
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	tokens := strings.Fields(s)
	counter := make(map[string]int)
	for _, token := range tokens {
		elem, ok := counter[token]
		if ok {
			counter[token] = elem + 1
		} else {
			counter[token] = 1
		}
	}
	return counter
}

func main() {
	wc.Test(WordCount)
}
