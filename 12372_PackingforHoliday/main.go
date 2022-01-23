package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	in, _ := os.Open("12372.in")
	defer in.Close()
	out, _ := os.Create("12372.out")
	defer out.Close()

	var ctr, n, l, w, h int

	for fmt.Fscanf(in, "%d\n", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d%d%d\n", &l, &w, &h)

		arr := []int{l, w, h}
		arrSlice := arr[:] //created a slice of the array
		sort.Sort(sort.Reverse(sort.IntSlice(arrSlice)))

		ctr++
		if arrSlice[0] <= 20 {
			fmt.Fprintf(out, "Case %d: good \n", ctr)
		} else {
			fmt.Fprintf(out, "Case %d: bad \n", ctr)
		}
	}
}
