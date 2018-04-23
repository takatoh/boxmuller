package main

import (
	"fmt"
	"os"
	"flag"
	"strconv"

	"github.com/takatoh/boxmuller"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
`Usage:
  %s [options] <N>
Options:
`, os.Args[0])
		flag.PrintDefaults()
	}
	opt_mu := flag.Float64("mu", 0.0, "specify mean value.")
	opt_sigma := flag.Float64("sigma", 1.0, "specify standard deviation.")
	flag.Parse()

	n, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "N must be integer.")
		flag.Usage()
		os.Exit(1)
	}

	m := n / 2

	bm := boxmuller.New(*opt_mu, *opt_sigma)
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
