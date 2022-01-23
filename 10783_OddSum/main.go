package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10783.in")
	defer in.Close()
	out, _ := os.Create("10783.out")
	defer out.Close()

	var caseno, c, a, b int
	for fmt.Fscanf(in, "%d\n", &c); c > 0; c-- {

		fmt.Fscanf(in, "%d\n%d\n", &a, &b)

		fmt.Printf("n value is %d\n l value is %d\n h value is %d\n", c, a, b)
		total := 0
		for i := a; i <= b; i++ {
			if i%2 != 0 {
				total += i
			}
		}
		caseno++
		fmt.Fprintf(out, "Case %d: %d\n", caseno, total)
		//fmt.Fprintln(out, caseno, total)
	}
}
