
/* 泛型：interface{} */

// Container is a generic container, accepting anything.
type Container []interface{}
// Put adds an element to the container.
func (c *Container) Put(elem interface{}) {
    *c = append(*c, elem)
}
// Get gets an element from the container.
func (c *Container) Get() interface{} {
    elem := (*c)[0]
    *c = (*c)[1:]
    return elem
}

func main() {
	// usage
	intContainer := &Container{}
	intContainer.Put(7)
	intContainer.Put(42)
	// 在把数据取出来时，因为类型是 interface{}，所以，还要做一个转型，只有转型成功，才能进行后续操作。
	// 因为 interface{}太泛了，泛到什么类型都可以放。

	// Type Assert
	// assert that the actual type is int
	elem, ok := intContainer.Get().(int)
	if !ok {
			fmt.Println("Unable to read an int from intContainer")
	}
	fmt.Printf("assertExample: %d (%T)\n", elem, elem)
}

