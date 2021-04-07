// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[0:] {
		fmt.Printf("args[%d] is %s\n", i, arg)
	}
}

