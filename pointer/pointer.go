package pointer

import "fmt"

type Code struct {
	num int
}

func normalAddOne(c Code) {
	c.num++
}

func pointerAddOne(c *Code) {
	c.num++
}

func Study1() {
	var code Code = Code{num: 7}
	fmt.Println(code) // 7
	normalAddOne(code)
	fmt.Println(code) // 7
	pointerAddOne(&code)
	fmt.Println(code) // 8
}
