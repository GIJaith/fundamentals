package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	in, _ := os.Open("11727.in")
	defer in.Close()
	out, _ := os.Create("11727.out")
	defer out.Close()

	var ctr, n, s1, s2, s3 int
	for fmt.Fscanf(in, "%d\n", &n); n > 0; n-- {
		//fmt.Printf("n value is %d\n", n)

		fmt.Fscanf(in, "%d%d%d\n", &s1, &s2, &s3)
		//fmt.Printf("%d %d %d \n", s1, s2, s3)

		srt := []int{s1, s2, s3}
		sort.Ints(srt)
		fmt.Println(srt[1])

		ctr++

		fmt.Fprintf(out, "Case %d: %d \n", ctr, srt[1])

	}
}
