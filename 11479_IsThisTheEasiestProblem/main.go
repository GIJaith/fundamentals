package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("11479.in")
	defer in.Close()
	out, _ := os.Create("11479.out")
	defer out.Close()

	var ctr, n, s1, s2, s3 int
	for fmt.Fscanf(in, "%d\n", &n); n > 0; n-- {
		fmt.Printf("n value is %d\n", n)

		fmt.Fscanf(in, "%d%d%d\n", &s1, &s2, &s3)
		fmt.Printf("%d %d %d \n", s1, s2, s3)

		ctr++
		switch {
		case s1 <= 0 || s2 <= 0 || s3 <= 0 || s1+s2 <= s3 || s2+s3 <= s1 || s1+s3 <= s2:
			fmt.Fprintf(out, "Case %d: Invalid\n", ctr) //Invalid
		case s1 == s2 && s1 == s3:
			fmt.Fprintf(out, "Case %d: Equilateral\n", ctr) //Equilateral
		case s1 == s2 || s1 == s3 || s2 == s3:
			fmt.Fprintf(out, "Case %d: Isoceles\n", ctr) //Isosceles
		default:
			fmt.Fprintf(out, "Case %d: Scalene\n", ctr) //Scalene
		}
	}
}
