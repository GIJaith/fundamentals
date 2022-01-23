package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10055.in")
	defer in.Close()
	out, _ := os.Create("10055.out")
	defer out.Close()

	var s1, s2 uint64
	for {
		if _, err := fmt.Fscanf(in, "%d%d\n", &s1, &s2); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(s1, s2)
		fmt.Fprintln(out, s2-s1)
	}
}
