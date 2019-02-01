package interfaces

import "fmt"

type Man struct {
	id   int
	name string
}

func (man *Man) String() string {
	return man.name
}

func testTransfer() {
	var o interface{} = &Man{1, "tom"}
	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println(i)
	}
	u := o.(*Man)
	fmt.Println(u.name, u.id, u)

	//  switch 做批量类型判断，不支持 fallthrough
	switch v := o.(type) {
	case nil:
		fmt.Println("nil")
	//case fmt.Stringer:
	//	fmt.Println(v)
	case func() string:
		fmt.Println(v())
	case Man:
		fmt.Println("man")
	case *Man:
		fmt.Println("*man")
	default:
		fmt.Println("default not know type...")
	}
}

// 超集接口对象可转换为子集接口，反之出错
type Singer interface {
	String() string
}

type Artisters interface {
	Singer
	print() string
}

type APerson struct {
	name string
}

func (person *APerson) String() string {
	return "String: " + person.name
}

func (person *APerson) print() string {

	return "print: " + person.name
}

func testTransInterface() {

	var p Artisters = &APerson{"singers"}
	fmt.Println(p)

	var o Singer = p
	//var pp Artisters = o  // error transfer...
	fmt.Println(o)
}
