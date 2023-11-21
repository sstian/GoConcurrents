/*
Go Generator
需要三个东西：
1. 一个函数模板。在里面设置好相应的占位符。
2. 一个函数生成脚本。用于按规则来替换文本并生成新的代码。
3. 一行注释代码。用于生成代码。
*/

// 3. 注释代码
// 生成包名 gen，类型是 uint32，目标文件名以 container 为后缀
//go:generate ./gen.sh ./template/container.tmp.go gen uint32 container
func generateUint32Example() {
    var u uint32 = 42
    c := NewUint32Container()
    c.Put(u)
    v := c.Get()
    fmt.Printf("generateExample: %d (%T)\n", v, v)
}

// 生成包名 gen，类型是 string，目标文件名是以 container 为后缀
//go:generate ./gen.sh ./template/container.tmp.go gen string container
func generateStringExample() {
    var s string = "Hello"
    c := NewStringContainer()
    c.Put(s)
    v := c.Get()
    fmt.Printf("generateExample: %s (%T)\n", v, v)
}

// 在工程目录中直接执行 go generate 命令，生成两份代码：
// uint32_container.go：
package gen
type Uint32Container struct {
    s []uint32
}
func NewUint32Container() *Uint32Container {
    return &Uint32Container{s: []uint32{}}
}
func (c *Uint32Container) Put(val uint32) {
    c.s = append(c.s, val)
}
func (c *Uint32Container) Get() uint32 {
    r := c.s[0]
    c.s = c.s[1:]
    return r
}

// string_container.go：
package gen
type StringContainer struct {
    s []string
}
func NewStringContainer() *StringContainer {
    return &StringContainer{s: []string{}}
}
func (c *StringContainer) Put(val string) {
    c.s = append(c.s, val)
}
func (c *StringContainer) Get() string {
    r := c.s[0]
    c.s = c.s[1:]
    return r
}


/*
第三方工具：
Genny
Generic
GenGen
Gen
*/