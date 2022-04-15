// Use this program to build, load, and test new neural networks
package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	VERSION = "Version 0.01: Workbench is a shell application that listens for input.\n" +
			  "Functions: Build new NN, Test old NN, Train and Save NNs." +
			  "Type -help for menu."
)

func main() {
	fmt.Println("[WORKBENCH]")

	reader := bufio.NewReader(os.Stdin)
	active := ""
	for {
		fmt.Printf("[%v] -> ",active)
		
	}
}

