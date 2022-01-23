package main

import (
	"fmt"
	"os"
)

func main() {

	in, _ := os.Open("12577.in")
	defer in.Close()
	out, _ := os.Create("12577.out")
	defer out.Close()

	var a string

	for ctr := 1; ; ctr++ {
		if fmt.Fscanf(in, "%s", &a); a == "*" {
			break
		}
		fmt.Fscanf(in, "%s\n", a)
		if a == "Hajj" {
			fmt.Fprintf(out, "Case %d: Hajj-e-Akbar\n", ctr)
		} else {
			fmt.Fprintf(out, "Case %d: Hajj-e-Asghar\n", ctr)
		}
	}

}
