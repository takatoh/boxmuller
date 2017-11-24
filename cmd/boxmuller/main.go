package main

import (
	"fmt"
//	"os"
	"flag"
	"strconv"

	"github.com/takatoh/boxmuller"
)

func main() {
	opt_mean := flag.Float64("mean", 0.0, "specify mean value.")
	flag.Parse()

	n, _ := strconv.Atoi(flag.Args()[0])
	m := n / 2

	bm := boxmuller.NewBoxMuller(*opt_mean, 0.2)
	for i := 0; i < m; i++ {
		z1, z2 := bm.Rand()
		fmt.Println(z1)
		fmt.Println(z2)
	}
	if n % 2 == 1 {
		z1, _ := bm.Rand()
		fmt.Println(z1)
	}
}
