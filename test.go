package main

import "fmt"

func main() {
	str := removeDuplicates([]string{"red", "black", "red", "pink", "blue", "pink", "blue"})
	fmt.Println(str)
}

// ：写一个函数，就地消除[]string中重复字符串，如：
// {"red", "black", "red", "pink", "blue", "pink", "blue"}
// 消除后为：{"red", "black", "pink", "blue"}

func removeDuplicates(strs []string) []string {
	// 请在这里完成你的代码
	map1 := make(map[string]int)
	slice := []string{}
	for _, v := range strs {
		map1[v] = map1[v] + 1
	}
	for k, _ := range map1 {
		slice = append(slice, k)
	}
	return slice
}
