package main

func main() {
	s := []int{1, 2, 4, 7}
	// 结果应该是 5, 1, 2, 4, 7
	s = Add(s, 0, 5)
	// 输出s中的元素

	// 结果应该是5, 9, 1, 2, 4, 7
	s = Add(s, 1, 9)
	for i := 0; i < len(s); i++ {
		println(s[i])
	}
	// 结果应该是5, 9, 1, 2, 4, 7, 13
	s = Add(s, 6, 13)

	// 结果应该是5, 9, 2, 4, 7, 13
	s = Delete(s, 2)

	// 结果应该是9, 2, 4, 7, 13
	s = Delete(s, 0)

	// 结果应该是9, 2, 4, 7
	s = Delete(s, 4)

}

func Add(s []int, index int, value int) []int {
	// 具体实现
	// 1. 先扩容
	s = append(s, 0)
	// 2. 把 index 之后的元素往后移动
	copy(s[index+1:], s[index:])
	// 3. 把 value 放到 index 位置
	s[index] = value
	// 4. 返回新的切片
	return s
}

func Delete(s []int, index int) []int {
	// 具体实现
	// 1. 把 index 之后的元素往前移动
	copy(s[index:], s[index+1:])
	// 2. 返回新的切片
	return s
}
