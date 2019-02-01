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

// again test
//type Animals interface {
//	Shark()
//}
type Sayer interface {
	Say(message string)
	SayHi()
}

type Animal struct {
	Name string
}

func (a *Animal) Say(message string) {
	fmt.Printf("Animal[%v] say: %v\n", a.Name, message)
}

func (a *Animal) SayHi() {
	a.Say("Hi")
}

//func (a *Dog) SayHi() {
//
//	fmt.Println(a.Name)
//}

// override Animal.Say
func (d *Dog) Say(message string) {
	fmt.Printf("Dog[%v] say: %v\n", d.Name, message)
}

type Dog struct {
	Animal
}

func runCat() {
	var sayer Sayer

	sayer = &Dog{Animal{Name: "Yoda"}}
	sayer.Say("hello world")

	//var animal *Animal = &Dog{Animal{Name: "Yoda"}} // 不支持父类指针指向子类，下面这种写法是不允许的
}

func runCat2() {
	var sayer Sayer

	sayer = &Dog{Animal{Name: "Yoda"}}
	sayer.Say("hello world") // Dog[Yoda] say: hello world
	sayer.SayHi()            // Animal[Yoda] say: Hi
}

// test inherit
type ParentInterfacer interface {
	myParent() string
}
type ParentSt struct {
	name string
}

func (p *ParentSt) myParent() string {
	return p.name
}

type ChildSt struct {
	*ParentSt
}

func runCat3() {
	//var s ParentSt = &ChildSt{ParentSt{"parent"}}  // 不支持父类指针指向子类，下面这种写法是不允许的
	var s ParentInterfacer = &ChildSt{&ParentSt{"parent"}}
	fmt.Println(s)
}
