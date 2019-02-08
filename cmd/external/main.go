package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Run first program
	command := "./added"
	dir := "/home/johan/go/src/github.com/johansundell/command-test/cmd/added"

	if err := runCommand(command, dir); err != nil {
		log.Fatal(err)
	}

	// Run second program
	command = "./helloer"
	dir = "/home/johan/go/src/github.com/johansundell/command-test/cmd/helloer"
	if err := runCommand(command, dir); err != nil {
		log.Fatal(err)
	}

}

func runCommand(command, dir string) error {
	args := []string{}
	cmd := exec.Command(command, args...)
	cmd.Dir = dir

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
