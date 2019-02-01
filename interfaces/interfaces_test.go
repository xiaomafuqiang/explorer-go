package interfaces

import (
	"fmt"
	"testing"
)

func TestInterfaces(t *testing.T) {

	main()

	runSuper()

	main2()

	fmt.Println(" runCopy ------------->> ")
	runCopy()
	fmt.Println(" multipleImplements ------------->> ")
	multipleImplements()
}
