/* 泛型：reflect */

type Container struct {
    s reflect.Value
}
func NewContainer(t reflect.Type, size int) *Container {
    if size <=0  { size=64 }
    return &Container{
        s: reflect.MakeSlice(reflect.SliceOf(t), 0, size), 
    }
}
func (c *Container) Put(val interface{}) error {
    if reflect.ValueOf(val).Type() != c.s.Type().Elem() {
        return fmt.Errorf("Put: cannot put a %T into a slice of %s", 
				val, c.s.Type().Elem()))
    }
    c.s = reflect.Append(c.s, reflect.ValueOf(val))
    return nil
}
func (c *Container) Get(refval interface{}) error {
    if reflect.ValueOf(refval).Kind() != reflect.Ptr ||
        reflect.ValueOf(refval).Elem().Type() != c.s.Type().Elem() {
        return fmt.Errorf("Get: needs *%s but got %T", c.s.Type().Elem(), refval)
    }
    reflect.ValueOf(refval).Elem().Set( c.s.Index(0) )
    c.s = c.s.Slice(1, c.s.Len())
    return nil
}

/*
在 NewContainer() 时，会根据参数的类型初始化一个 Slice。
在 Put() 时，会检查 val 是否和 Slice 的类型一致。
在 Get() 时，需要用一个入参的方式，因为没有办法返回 reflect.Value 或 interface{}，不然还要做 Type Assert。
不过有类型检查，所以，必然会有检查不对的时候，因此，需要返回 error
*/

func main() {
	// usage
	f1 := 3.1415926
	f2 := 1.41421356237
	c := NewMyContainer(reflect.TypeOf(f1), 16)
	if err := c.Put(f1); err != nil {
		panic(err)
	}
	if err := c.Put(f2); err != nil {
		panic(err)
	}
	g := 0.0
	if err := c.Get(&g); err != nil {
		panic(err)
	}
	fmt.Printf("%v (%T)\n", g, g) //3.1415926 (float64)
	fmt.Println(c.s.Index(0)) //1.4142135623
}
