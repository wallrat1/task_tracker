package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)
	r := []rune(text)
	res := text
	if width < len(r) {
		res = string(r[0:width]) + "..."
	}
	fmt.Println(res)
}
