package interfaces

import (
	"fmt"
)

type humaner interface {
	run() (speed int)
}

type gender interface {
	getGender() string
}

type man struct {
	name string
	age  int
}

func (human *man) run() int {
	return human.age
}

func (huaman *man) getGender() string {

	return "man"
}

func main() {

	var m = &man{
		name: "mark",
		age:  222,
	}
	speed := m.run()
	fmt.Println(speed)
	getGender := m.getGender()
	fmt.Println(getGender)

	fmt.Println("============ type ================")

}

// 多态
type super interface {
	print()
}

type suber interface {
	super

	print2()
}

type superSuber struct {
	name string
	age  int
}

func (s *superSuber) print2() {
	fmt.Println("print2 ")
	//s.name = "====================== new super.name"
}

func (s *superSuber) print() {
	fmt.Println("print ")
}

func runSuper() {
	var supers = superSuber{
		"mark", 3434,
	}
	supers.print()
	supers.print2()
}

// test type inner
type Tester struct {
	s interface { // can be type struct which implements it
		String() string
	}
}

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user %d, %s", self.id, self.name)
}

func main2() {
	var t = Tester{&User{1, "Tom"}}
	fmt.Println(t.s.String())
}

// test copy receiver ? no for *class but for value always yes
type copyier interface {
	run2()
}
type CopierClass struct {
	name string `json:"for pointer_check"`
}

func (c *CopierClass) run2() {
	c.name = "new name in interface method..."
	fmt.Println("*CopierClass", &c.name, c.name)
}

func runCopy() {
	var copyInstance = CopierClass{
		name: "for check address name",
	}
	fmt.Println("copyInstance", &copyInstance.name, copyInstance.name)
	copyInstance.run2()
	fmt.Println("copyInstance twice...", &copyInstance.name, copyInstance.name)

}

// test for multiple implements
type aInterface interface {
	me()
}
type bInterface interface {
	me()
}

// implements class
type abStruct struct {
	name string `info:"for info"`
}

func (abIs *abStruct) me() {
	fmt.Println("who am i?")
}

func multipleImplements() {
	var ab = abStruct{
		name: "is ab",
	}
	var abPointer = &abStruct{
		name: "is ab pointer",
	}
	var abA aInterface = &abStruct{
		name: "is aInterface",
	}
	var abB bInterface = &abStruct{
		name: "is bInterface",
	}

	fmt.Println(ab, abPointer, abA, abB)

	fmt.Println("multiple")
	type abStructParent struct {
		aInterface
		bInterface
	}

	var abParent = abStructParent{
		abA, // use interface but cannot get any field about abA.name
		abB,
	}
	fmt.Println(abParent)
}
