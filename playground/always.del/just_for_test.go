package always_del

import (
	"fmt"
	"testing"
)

type Animal interface {
	Speak() string
}

type Cat struct{}

func (c *Cat) Speak() string {
	return "cat"
}

type Dog struct{}

func (d Dog) Speak() string {
	return "dog"
}

func test22(params interface{}) {
	fmt.Println(params)
}
func test444(p Animal) {
	speak := p.Speak()
	fmt.Println(speak)
}
func main() {
	animals := []Animal{&Cat{}, Dog{}}
	for _, animal := range animals {
		test444(animal)
	}

	test22("string")
	test22(123)
	test22(true)
}

func TestM(t *testing.T) {
	main()

}
