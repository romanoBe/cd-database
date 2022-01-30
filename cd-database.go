package main

import (
	"bufio"
	"os"
)

const programTitle = "CD Database"

func main() {
	println(programTitle)

	for {
		println("Type in command: q - quit, l - list CDs, a - add CD, d - delete CD")
		cmd := readCmd()
		switch cmd {
		case 'q':
			return
		case 'l':
			listCDs()
		case 'a':
			addCD()
		case 'd':
			deleteCD()
		}
	}
}

func readCmd() rune {
	reader := bufio.NewReader(os.Stdin)
	cmd, _, _ := reader.ReadRune()
	return cmd
}

func listCDs() {
	println("Displaying list of CDs")
}

func addCD() {
	println("Adding new CD")
}

func deleteCD() {
	println("Delete CD")
}
