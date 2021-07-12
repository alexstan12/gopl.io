package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/weightconv"
)

func main(){

	if len(os.Args) > 1 {
		for _, v:= range os.Args[1:] {
			v, _ := strconv.Atoi(v)
			fmt.Fprintf(os.Stdout, "You have " + weightconv.Kilograms(float64(v)).String() + " and " + weightconv.KgToLbs(weightconv.Kilograms(float64(v))).String())
		}
	}

}
