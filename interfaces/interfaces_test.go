package interfaces

import (
	"fmt"
	"testing"
)

func TestInterfaces(t *testing.T) {
	fmt.Println("interfaces.go")
	main()
	runSuper()
	main2()
	fmt.Println(" runCopy ------------->> ")
	runCopy()
	fmt.Println(" multipleImplements ------------->> ")
	multipleImplements()

}

func TestInterfaceTable(tt *testing.T) {
	fmt.Println("interfaceTable.go")

	user2CopyInterface()

	nilInterfaceCondition()
}

func TestInterfaceTransfer(t *testing.T) {
	//testTransfer()
	testTransInterface()
}
