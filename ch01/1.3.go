package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// TODO 做试验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], " "))
}
