package main

/**
 从函数返回局部变量的指针是安全的，编译器会通过逃逸分析来决定是否在堆上分配内存，
 如果不禁用内联，直接在栈上分配内存；

 当前编译器并没有实现尾递归优化，尽管Go执行栈上限是GB规模，轻易不会出现堆栈溢出错误，但依然需要注意拷贝栈复杂成本

 不管是指针，引用类型，还是其他类型参数，都是值拷贝传递，无非是拷贝目标对象，还是拷贝指针自身而已
**/
func test() *int {
	a := 0x100
	return &a
}

func main() {
	var a *int = test()
	println(a, *a)
}