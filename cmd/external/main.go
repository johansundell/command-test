package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	command := "./added"
	dir := "/home/johan/go/src/github.com/johansundell/command-test/cmd/added"
	args := []string{}
	cmd := exec.Command(command, args...)
	cmd.Dir = dir

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
