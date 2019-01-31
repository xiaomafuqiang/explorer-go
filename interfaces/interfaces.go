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

type superSubers struct {
	name string
	age  int
}

func (s *superSubers) print2() {
	fmt.Println("print2 ")
}

func (s *superSubers) print() {
	fmt.Println("print ")
}

func runSuper() {
	var supers = superSubers{
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
	t := Tester{&User{1, "Tom"}}
	fmt.Println(t.s.String())
}
