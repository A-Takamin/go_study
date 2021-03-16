package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	sum(3, 5)

	// for文
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while文
	n := 0
	for n < 10 {
		fmt.Println(n)
	}

	// 無限ループとbreakとcontinue
	n = 0
	for {
		n++
		if n < 10 {
			break
		}
		if n%2 == 0 {
			continue
		}
		fmt.Print(n)
	}

	// switch。fallthroughで次のcaseへまたぐ。
	n = 10
	switch n {
	case 5:
		fmt.Println("hoge")
	case 10, 15:
		fmt.Println("fuga")
		fallthrough
	case 20, 25:
		fmt.Println("piyo")
	}
}

func sum(x, y int) int {
	return x + y
}
