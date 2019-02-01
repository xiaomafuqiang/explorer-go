package interfaces

import "fmt"

type Testers interface {
	Do()
}
type myInt int

func (i *myInt) Do() {
	fmt.Println(i)
}
func testSkill() {
	var myint myInt = myInt(9)
	var myint2 myInt = 8
	fmt.Println(myint)
	myint.Do()
	myint2.Do()
}
