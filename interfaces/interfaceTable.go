package interfaces

import (
	"fmt"
	"reflect"
	"unsafe"
)

func user2CopyInterface() {
	u := User{1, "Tom"}
	var vi, pi interface{} = u, &u

	// vi.(User).name = "Jack"       // Error: cannot assign to vi.(User).name 只有使用指针才能修改其状态。
	pi.(*User).name = "Jack"

	fmt.Println(vi.(User), vi, u)
	fmt.Println(pi.(*User), pi)
}

var a interface{} = nil         // tab = nil, data = nil
var b interface{} = (*int)(nil) // tab 包含 *int 类型信息, data = nil

type iface struct {
	itab, data uintptr
}

func nilInterfaceCondition() {
	ia := (*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))

	fmt.Println("nilInterfaceCondition")

	fmt.Println(a == nil, ia, reflect.ValueOf(nil))          // reflect.ValueOf(nil) == reflect.ValueOf(a) <invalid reflect.Value>
	fmt.Println(b == nil, b, ib, reflect.ValueOf(b).IsNil()) // false <nil> {5241632 0} true

}
