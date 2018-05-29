// $ go build cmdline.go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(len(os.Args))

	for _, params := range os.Args {
		fmt.Println(params)
	}
}
