package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	//note ignoring potential errors from input.err()
	for line, n := range counts {
		if n > 1 {
			fmt.Print("%d\t%s\n", n, line)
		}
	}

}
