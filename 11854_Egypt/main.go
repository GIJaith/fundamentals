package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	in, _ := os.Open("11854.in")
	defer in.Close()
	out, _ := os.Create("11854.out")
	defer out.Close()

	var l1, l2, h int
	for {
		if fmt.Fscanf(in, "%d%d%d\n", &l1, &l2, &h); l1 == 0 && l2 == 0 && h == 0 {
			break
		}

		arr := []int{l1, l2, h}
		arrSlice := arr[:] //created a slice of the array
		sort.Sort(sort.Reverse(sort.IntSlice(arrSlice)))

		if (arrSlice[1]*arrSlice[1])+(arrSlice[2]*arrSlice[2]) == (arrSlice[0] * arrSlice[0]) {
			fmt.Fprintln(out, "right")
		} else {
			fmt.Fprintln(out, "wrong")
		}
	}
}
