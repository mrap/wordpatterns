package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	filename := os.Args[1]
	fmt.Println("Digesting words from", filename)
	t := CreateWordmap(filename)

	var (
		query   string
		results []string
	)

	for true {
		fmt.Print("\nEnter a string to search for: ")
		fmt.Scanf("%s", &query)
		results = t.WordsContaining(query)
		fmt.Printf("===== %d Words Containing '%s' =====\n\n", len(results), query)
		fmt.Println(results)
	}
}
