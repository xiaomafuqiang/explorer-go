package methods

import "fmt"

type User struct {
	id   int
	name string
}

func (self User) TestValue(info string) {
	fmt.Println(&self, &self.name, info)
}

func (self *User) TestPointer(info string) {
	fmt.Println(&self, &self.name, info)
}

// always copy receiver even if [  (*User).TestValue override]
func testValues() {
	fmt.Println("-----------------------------------------")

	u := User{1, "Tom"}
	fmt.Println(&u.name, "u")

	// 方法调用
	u.TestValue("u.TestValue()")
	var testValue = u.TestValue
	testValue("testValue")

	var userTestValue = User.TestValue
	userTestValue(u, "User.TestValue")
	var userXTestValue = (*User).TestValue
	userXTestValue(&u, "*User.TestValue")
	fmt.Println("diff method...")

	var uu = &u
	fmt.Println("res", uu)

}

// cannot get expression by [ User.TestPointer x ] receiver error only by [ (*User).TestPointer √ ]
// always one instance handler...
func testPointers() {
	fmt.Println("-----------------------------------------")

	uuu := &User{1, "Tom"}
	u := *uuu
	fmt.Println(&u.name, "u")

	// 方法调用
	u.TestPointer("u.TestPointer()")
	var testPointer = u.TestPointer
	testPointer("testPointer")

	//var userTestPointer = User.TestPointer // ### invalid method expression User.TestPointer (needs methods receiver: (*User).TestPointer)
	var userXTestPointer = (*User).TestPointer
	userXTestPointer(&u, "*User.TestPointer")
	fmt.Println("diff method...")

	fmt.Println("-----------------------------------------")

	var uu = &u
	fmt.Println("res", &uu.name)
}

type Data struct{}

func (Data) TestValue() {}

func (*Data) TestPointer() {}

// should be diff for *Data or Data get error nil
func testNilBindExpressionOrValue() {

	var p *Data = nil
	p.TestPointer()

	fmt.Println((*Data)(nil), p)

	(*Data)(nil).TestPointer() // method value
	(*Data).TestPointer(nil)   // method expression

	// p.TestValue()            // invalid memory address or nil methods dereference

	// (Data)(nil).TestValue()  // cannot convert nil to type Data
	// Data.TestValue(nil)      // cannot use nil as type Data in function argument
}
