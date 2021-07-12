//Echo4 prints its command line arguments
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", true, "selects wheter it applies a new line at the end")
var s = flag.String("s", " ", "the delimeter between printed strings")

func main(){
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *s))

	if *n {
		fmt.Println()
	}
}
