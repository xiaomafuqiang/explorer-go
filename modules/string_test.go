package modules

import (
	"fmt"
	"testing"
)

type Pr struct {
	name string
	v    int
}

func Test(tt *testing.T) {

	var s = []*Pr{{"aa", 1}}
	fmt.Println(s)
	var ss = []Pr{{"aa", 1}}
	fmt.Println(ss)
}
