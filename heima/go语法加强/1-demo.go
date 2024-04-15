package go语法加强

var a *string

//func main() {
//	p := Person{"mike", 'm', 18}
//	fmt.Println(p)
//	// 索引成员变量 "."
//	fmt.Println(p.sex)
//	fmt.Println(p.age)
//	fmt.Println(p.name)
//}

//func main() {
//	//赋值初始化结构体
//	var man3 Person
//	man3.name = "mike"
//	man3.sex = 'm'
//	man3.age = 18
//
//	//取地址实例化
//	fmt.Println(man3)
//
//	// 结构体的比较和赋值，只能使用== != 比较
//	var p1 = Person{"andy", 'm', 20}
//	var p2 = Person{"andy", 'm', 20}
//	var p3 = Person{"andy", 'm', 20}
//	fmt.Println("p1 == p2", p1 == p2)
//	fmt.Println("p1 == p3", p1 == p3)
//
//	// 相同类型的结构体赋值
//	var p4 Person
//	p4 = p3
//	fmt.Println("p4 = ", p4)
//
//	var temp Person
//	// 结构体变量的大小
//	ptr := unsafe.Sizeof(temp)
//	fmt.Println("ptr = ", ptr)
//	test(temp)
//	fmt.Println("temp = ", temp)
//
//}
//
//// 结构体传参是值传递，不会改变原来的值，但是会浪费内存空间，所以一般传指针
///**
//	type Person struct {
//        name string
//        sex  [1000]byte
//        age  [1000]int // 这样就会拷贝1000*8个字节
//	}
//*/
//func test(man Person) {
//	man.age = 20
//	man.sex = 'a'
//	man.name = "ddd"
//	fmt.Println("ptr = ", unsafe.Sizeof(man))
//}
