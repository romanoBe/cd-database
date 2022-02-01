package main

import (
	"bufio"
	"fmt"
	"os"
)

const programTitle = "CD Database"

func main() {
	var albums []string
	println(programTitle)

	for {
		println("Type in command: q - quit, l - list CDs, a - add CD, d - delete CD")
		cmd := readCmd()
		switch cmd {
		case 'q', 'Q':
			return
		case 'l', 'L':
			listCDs(albums)
		case 'a', 'A':
			// albums = getNewCD()
			albums = append(albums, getNewCD()) /* append - jak metoda push w JS */
		case 'd', 'D':
			deleteCD()
		}
	}
}

func readCmd() rune {
	reader := bufio.NewReader(os.Stdin)
	cmd, _, _ := reader.ReadRune()
	return cmd
}

func listCDs(albums []string) {
	println("Displaying list of CDs")
	for i, v := range albums { /* odpowiednik for each */
		fmt.Printf("%d. %s", i+1, v)
	}
}

func getNewCD() string {
	println("Adding new CD")
	return readString("Input album title: ")
}

func deleteCD() {
	println("Delete CD")
}

func readString(title string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(title)
	input, _ := reader.ReadString('\n')
	// input = strings.TrimSuffix(input, "\n")
	return input
}
