package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("12468.in")
	defer in.Close()
	out, _ := os.Create("12468.out")
	defer out.Close()

	var c1, c2 int
	for {
		if fmt.Fscanf(in, "%d%d\n", &c1, &c2); c1 == -1 && c2 == -1 {
			fmt.Println("end of file")
			break
		}
		var clickcnt int

		if c2 > c1 {
			clickcnt = c2 - c1
		} else {
			clickcnt = c1 - c2
		}
		if clickcnt >= 50 {
			clickcnt = 100 - clickcnt
		}
		fmt.Println(c1, c2, clickcnt)
		fmt.Fprintln(out, clickcnt)

	}

}
