package main

import (
	"fmt"
	"github.com/takatoh/boxmuller"
)

func main() {
	bm := boxmuller.NewBoxMuller(1.0, 0.2)
	for i := 0; i < 5000; i++ {
		z1, z2 := bm.Rand()
		fmt.Println(z1)
		fmt.Println(z2)
	}
}
