/* 接口完整性检查 */

package main

import "fmt"

type Shape interface {
	Sides() int
	Area() int
}
type Square struct {
	len int
}

func (s *Square) Sides() int {
	return 4
}
func main() {
	// 声明一个 _ 变量（没人用）会把一个 nil 的空指针从 Square 转成 Shape。
	// 如果没有实现完相关的接口方法，编译器就会报错：
	// cannot use (*Square)(nil) (value of type *Square) as Shape value in variable declaration: *Square does not implement Shape (missing method Area)
	var _ Shape = (*Square)(nil)

	s := Square{len: 5}
	fmt.Printf("%d\n", s.Sides())
}

/*
资源清理
出错后是需要做资源清理的，不同的编程语言有不同的资源清理的编程模式。
C 语言：使用的是 goto fail; 的方式到一个集中的地方进行清理。
C++ 语言：一般来说使用 RAII 模式，通过面向对象的代理模式，把需要清理的资源交给一个代理类，然后再析构函数来解决。
  RAII = Resource Acquisition Is Initialization 资源获取即初始化：在构造函数中申请分配资源，在析构函数中释放资源。
Java 语言：可以在 finally 语句块里进行清理。
Go 语言：使用 defer 关键词进行清理。
*/
