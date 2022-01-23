package main

import (
	"fmt"
	"os"
)

var langMap = map[string]string{"HELLO": "ENGLISH", "HOLA": "SPANISH", "HALLO": "GERMAN", "BONJOUR": "FRENCH", "CIAO": "ITALIAN", "ZDRAVSTVUJTE": "RUSSIAN"}

func gsearch(greet string) string {
	if language, i := langMap[greet]; i {
		return language
	}
	return "UNKNOWN"
}

func main() {
	in, _ := os.Open("12250.in")
	defer in.Close()
	out, _ := os.Create("12250.out")
	defer out.Close()

	var greet string

	for ctr := 1; ; ctr++ {
		if fmt.Fscanf(in, "%s\n", &greet); greet == "#" {
			fmt.Println(greet)
			break
		}
		fmt.Println(greet)
		fmt.Fprintf(out, "Case %d: %s\n", ctr, gsearch(greet))

	}
}
