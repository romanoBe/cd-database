package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/romanoBe/cd-database/model"
)

const programTitle = "CD Database"

func main() {
	var myCollection model.Collection
	println(programTitle)

	for {
		println("Type in command: q - quit, l - list CDs, a - add CD, d - delete CD")
		cmd := readCmd()
		switch cmd {
		case 'q', 'Q':
			return
		case 'l', 'L':
			listCDs(myCollection.Albums)
		case 'a', 'A':
			artist, title, label, year := getNewCD()
			album := model.Album{
				Artist:      artist,
				Title:       title,
				RecordLabel: label,
				Year:        year,
			}
			myCollection.Albums = append(myCollection.Albums, album)
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

func listCDs(albums []model.Album) {
	println("Displaying list of CDs")
	for i, v := range albums {
		fmt.Printf("%d. %s \"%s\" %s %d \n", i+1, v.Artist, v.Title, v.RecordLabel, v.Year)
	}
}

func getNewCD() (string, string, string, int) {
	println("Adding new CD")
	artist := readString("Input artist: ")
	title := readString("Input album title: ")
	label := readString("Input record label: ")
	year, err := strconv.Atoi(readString("Input release year: "))
	if err != nil {
		println("ERROR!", err.Error())
		panic("Wychodzę")
	}
	return artist, title, label, year
}

func deleteCD() {
	println("Delete CD")
}

func readString(title string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(title)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\r\n")
	return input
}
