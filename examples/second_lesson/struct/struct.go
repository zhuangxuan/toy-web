package main

import "fmt"

func main() {
	// duck1 是 *ToyDuck
	//duck1 := &ToyDuck{}

	// duck1 := ToyDuck{} duck1 := &ToyDuck{} 两者的区别
	//  值类型与指针类型的区别
	//  值类型：值类型的变量直接存储值，内存通常在栈中分配。值类型的变量的赋值、传参都会产生一次复制动作。
	//  指针类型：指针类型的变量存储的是一个地址，这个地址存储最终的值。内存通常在堆中分配。指针类型的变量的赋值、传参不会产生复制动作，只是传递地址。

	//// 修改duck1中的字段
	//duck1.Color = "红色"
	//duck1.Price = 100
	//
	//duck1.Swim()

	//duck2 := ToyDuck{}
	//duck2.Color = "红色"
	//duck2.Price = 100
	//duck2.Swim()
	//
	//// duck3 是 *ToyDuck
	//duck3 := new(ToyDuck)
	//duck3.Swim()
	//
	//// 当你声明这样的时候，Go 就帮你分配好内存
	//// 不用担心空指针的问题，以为它压根就不是指针
	//var duck4 ToyDuck
	//duck4.Swim()
	//
	////// duck5 就是一个指针了
	//var duck5 *ToyDuck
	//// 这边会直接panic 掉
	//duck5.Swim()
	//
	//// 赋值，初始化按字段名字赋值
	//duck6 := ToyDuck{
	//	Color: "黄色",
	//	Price: 100,
	//}
	//duck6.Swim()
	//
	//// 初始化按字段顺序赋值，不建议使用
	//duck7 := ToyDuck{"蓝色", 1024}
	//duck7.Swim()
	//
	//// 后面再单独赋值
	//duck8 := ToyDuck{}
	//duck8.Color = "橘色"

}

// ToyDuck 玩具鸭
type ToyDuck struct {
	Color string
	Price uint64
}

func (t *ToyDuck) Swim() {
	fmt.Printf("门前一条河，游过一群鸭，我是%s，%d一只\n", t.Color, t.Price)
}
