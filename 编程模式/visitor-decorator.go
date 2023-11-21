/* Visitor + Decorator */

package main

import "fmt"

type VisitorFunc func(*Info, error) error
type Visitor interface {
    Visit(VisitorFunc) error
}
type Info struct {
    Namespace   string
    Name        string
    OtherThings string
}
func (info *Info) Visit(fn VisitorFunc) error {
  	return fn(info, nil)
}

type NameVisitor struct {
  visitor Visitor
}
func (v NameVisitor) Visit(fn VisitorFunc) error {
  return v.visitor.Visit(func(info *Info, err error) error {
    fmt.Println("NameVisitor() before call function")
    err = fn(info, err)
    if err == nil {
      fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
    }
    fmt.Println("NameVisitor() after call function")
    return err
  })
}

type OtherThingsVisitor struct {
  visitor Visitor
}
func (v OtherThingsVisitor) Visit(fn VisitorFunc) error {
  return v.visitor.Visit(func(info *Info, err error) error {
    fmt.Println("OtherThingsVisitor() before call function")
    err = fn(info, err)
    if err == nil {
      fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
    }
    fmt.Println("OtherThingsVisitor() after call function")
    return err
  })
}

type VisitorDecorator func(VisitorFunc) VisitorFunc
type DecoratedVisitor struct {
  visitor    Visitor
  decorators []VisitorDecorator
}
func NewDecoratedVisitor(v Visitor, fn ...VisitorDecorator) Visitor {
  if len(fn) == 0 {
    return v
  }
  return DecoratedVisitor{v, fn}
}
// Visit implements Visitor
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
  return v.visitor.Visit(func(info *Info, err error) error {
    decoratorLen := len(v.decorators) 
	for i := range v.decorators { 		
d := v.decorators[decoratorLen-i-1] 		
fn = d(fn)
 	} 
	return fn(v.visitor.(*Info), nil)
	})
}

func main() {
	//使用
  info := Info{}
  var v Visitor = &info
 	v = NewDecoratedVisitor(v, NameVisitor, OtherThingsVisitor)

  loadFile := func(info *Info, err error) error {
    info.Name = "Hao Chen"
    info.Namespace = "MegaEase"
    info.OtherThings = "We are running as remote team."
    return nil
  }

	v.Visit(loadFile)

}

/* 
error!
# command-line-arguments
.\visitor-decorator.go:82:29: NameVisitor (type) is not an expression
.\visitor-decorator.go:82:42: OtherThingsVisitor (type) is not an expression

一个 DecoratedVisitor 的结构来存放所有的VistorFunc函数；
NewDecoratedVisitor 可以把所有的 VisitorFunc转给它，构造 DecoratedVisitor 对象；
DecoratedVisitor实现了 Visit() 方法，里面就是来做一个 for-loop，顺着调用所有的 VisitorFunc。
*/

