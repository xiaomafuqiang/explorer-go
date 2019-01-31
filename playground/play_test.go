package playground

import (
	"fmt"
	"runtime"
	"testing"
)

func TestPlay(tt *testing.T) {

	sliceCopyWillBeValueAnother()

	goDeferMistake()

	getCurrentInfo()

}

func sliceCopyWillBeValueAnother() {
	var sl = []int{1, 2, 3, 4, 5, 6}
	var sl2 = make([]int, 3)

	copy(sl2, sl[:3])

	sl2[2] = 9
	fmt.Println(sl, sl2)
}

func goDeferMistake() {
	var i = 1

	// second call defer ⬆️⬆️
	defer fmt.Println("result: ", func() int { return i * 2 }()) // not func(){}()  inline reserve vars directly

	// first call defer ⬆️
	defer func() { // call func delay also vars calculated delay
		fmt.Println("result: ", func() int {
			return i * 2
		}())
	}()
	i++
}

func getCurrentInfo() {
	cpu := runtime.NumCPU()
	goroutine := runtime.NumGoroutine()
	call := runtime.NumCgoCall()
	fmt.Println(cpu, goroutine, call)
}
