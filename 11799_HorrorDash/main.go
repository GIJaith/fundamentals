package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	in, _ := os.Open("11799.in")
	defer in.Close()
	out, _ := os.Create("11799.out")
	defer out.Close()

	var ctr, n, r1, r2, r3, r4, r5, r6 int
	for fmt.Fscanf(in, "%d\n", &n); n > 0; n-- {
		fmt.Printf("value of N is %d\n", n)

		fmt.Fscanf(in, "%d%d%d%d%d%d\n", &r1, &r2, &r3, &r4, &r5, &r6)
		fmt.Printf("%d %d %d %d %d %d \n", r1, r2, r3, r4, r5, r6)

	
		arr := []int{r2, r3, r4, r5, r6}
		arrSlice := arr[:] //created a slice of the array
		sort.Sort(sort.Reverse(sort.IntSlice(arrSlice)))

		ctr++

		fmt.Fprintf(out, "Case %d: %d \n", ctr, arrSlice[0])

	}
}
