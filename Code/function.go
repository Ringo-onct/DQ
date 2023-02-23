package main
import (
	"log"
	"fmt"

	"github.com/mattn/go-tty"
	"github.com/k0kubun/go-ansi"
)
func getkey() int {
	key := 0
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer tty.Close()

	for {
		r, _ := tty.ReadRune()
		if r != 0 {
			if int(r) == 27 {
				key = 1
			} else if r == 91 {
				key = 2
			} else if key == 2 {
				switch r {
				case 65:
					return 128
				case 66:
					return 129
				case 67:
					return 130
				case 68:
					return 131
				}
				key = 0
			} else {
				return int(r)
			}
		}
	}
}

func chose(line int, mode int) int{
	i := 0
	if mode == 1 {
		ansi.CursorNextLine(1)
		ansi.CursorUp(line + 2)
		ansi.CursorForward(2)
		fmt.Printf(" >")
		line++
	} else {
		ansi.CursorNextLine(1)
		ansi.CursorUp(line - 2)
		fmt.Printf(" >")
	}

	for {
		x := getkey()
		if x == 128 {
			i--
			if i >= 0 {
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(2)
				ansi.CursorUp(1)
				fmt.Printf(" >")
			} else {
				i++
			}
		} else if x == 129 {
			i++
			if i <= (line - 1) {
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(1)
				ansi.CursorDown(1)
				fmt.Printf(">")
			} else {
				i--
			}
		} else if x == 13 {
			x := i
			for x <= 1 {
				ansi.CursorDown(1)
				x++
			}
			fmt.Println("")
			break
		}
	}
	return i
}
