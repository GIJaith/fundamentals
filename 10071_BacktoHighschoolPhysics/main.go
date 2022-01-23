package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10071.in")
	defer in.Close()
	out, _ := os.Create("10071.out")
	defer out.Close()

	var v, t int
	for {
		if _, err := fmt.Fscanf(in, "%d%d\n", &v, &t); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(v, t)
		fmt.Fprintln(out, 2*v*t)
	}
}
