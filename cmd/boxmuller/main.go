package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/takatoh/boxmuller"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	m := n / 2

	bm := boxmuller.NewBoxMuller(1.0, 0.2)
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
