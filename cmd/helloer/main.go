package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("What is your name")
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello", str)
}
