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
	fmt.Println("Diegesting words from", filename)
	t := CreateTrie(filename)

	var query string
	for true {
		fmt.Print("\nEnter a string to search for: ")
		fmt.Scanf("%s", &query)
		fmt.Printf("===== Words Containing '%s' =====\n\n", query)
		fmt.Println(t.WordsContaining(query))
	}
}
