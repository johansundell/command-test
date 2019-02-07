package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Give me 2 numbers and I will add them")
	reader := bufio.NewReader(os.Stdin)
	str := "lte"
	err := errors.New("test")
	for str[0:1] != "0" {
		str, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("You wrote", str)
	}

}
