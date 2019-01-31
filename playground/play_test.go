package playground

import (
	"fmt"
	"testing"
)

func TestPlay(tt *testing.T) {
	var sl = []int{1, 2, 3, 4, 5, 6}
	var sl2 = make([]int, 3)

	copy(sl2, sl[:3])

	sl2[2] = 9
	fmt.Println(sl, sl2)

}
