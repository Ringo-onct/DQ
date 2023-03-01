//繰り返し呼び出す形の処理をここに書く。

package main
import (
	"log"
	"fmt"
	"runtime"
	"os/exec"
	"os"

	"github.com/mattn/go-tty"
	"github.com/k0kubun/go-ansi"
)

//キー入力を認識する関数
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
			} else if key == 1 && r == 91 {
				key = 2
			} else if key == 2 {
				switch r {
				case 65:
					return 128	//うえ
				case 66:
					return 129	//した
				case 67:
					return 130	//みぎ
				case 68:
					return 131	//ひだり
				}
				key = 0
			} else {
				return int(r)
			}
		}
	}
}

//選択させる関数。引数に選択肢の数を入れることで動く。
func chose(line int) int{
	i := 1
	ansi.CursorUp(line + 1)
	ansi.CursorForward(1)
	fmt.Printf(" >")

	for {
		x := getkey()
		if x == 128 {
			i--
			if i >= 1 {
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
			if i <= line {
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(2)
				ansi.CursorDown(1)
				fmt.Printf(" >")
			} else {
				i--
			}
		} else if x == 13 {
			x := i - 1
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

//文字を入力させる関数
func namewrite() string {
	str := ""
	fmt.Println("　　　　　　なまえをきめてね")
	fmt.Println("　　　　　 ＊ ＊ ＊ ＊ ＊ ＊")
	fmt.Println("　あ　い　う　え　お　は　ひ　ふ　へ　ほ")
	fmt.Println("　か　き　く　け　こ　ま　み　む　め　も")
	fmt.Println("　さ　し　す　せ　そ　や　　　ゆ　　　よ")
	fmt.Println("　た　ち　つ　て　と　ら　り　る　れ　ろ")
	fmt.Println("　な　に　ぬ　ね　の　わ　　　を　　　ん")
	fmt.Println("　っ　ゃ　ゅ　ょ　゛　゜　もどる　おわる")

	ansi.CursorUp(6)
	fmt.Printf(" >")
	y := 6
	x := 1
	i := 1	//文字指定
	for {
		key := getkey()

		if key == 128 {	//Up
			if (x == 7 || x == 9) && (y == 1 || y == 3) {	//やゆよとわをんの間を避ける
				y += 2
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(2)
				ansi.CursorUp(2)
				fmt.Printf(" >")
			} else if y < 6 {
				y++
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(2)
				ansi.CursorUp(1)
				fmt.Printf(" >")
			}
		} else if key == 129 {	//Down
			if (x == 7 || x == 9) && (y == 5 || y == 3) {	//やゆよとわをんの間を避ける
				y -= 2
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(2)
				ansi.CursorDown(2)
				fmt.Printf(" >")
			} else if (x == 8 || x == 10) && y == 2 {	//もどるとおわるにも下がれるようにする
				y--
				x--
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(2)
				ansi.CursorDown(1)
				ansi.CursorBack(4)
				fmt.Printf(" >")
			} else if y > 1 {
				y--
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(2)
				ansi.CursorDown(1)
				fmt.Printf(" >")
			}
		} else if key == 130 {	//right
			if (x == 6 || x == 8) && (y == 4 || y ==2) {	//やゆよとわをんの間を避ける
				x += 2
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(2)
				ansi.CursorForward(8)
				fmt.Printf(" >")
			} else if x == 7 && y == 1 {	//もどるを避ける
					x += 2
					ansi.CursorBack(1)
					fmt.Printf(" ")
					ansi.CursorBack(2)
					ansi.CursorForward(8)
					fmt.Printf(" >")
			} else if x == 9 && y == 1 {
				//何もしない
			} else if x < 10 {
				x++
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(2)
				ansi.CursorForward(4)
				fmt.Printf(" >")
			}
		} else if key == 131 {	//left
			if (x == 10 || x == 8) && (y == 2 || y == 4) {	//やゆよとわをんの間を避ける
				x -= 2
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(2)
				ansi.CursorBack(8)
				fmt.Printf(" >")
			} else if x == 9 && y == 1 {	//もどるを避ける
				x -= 2
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(2)
				ansi.CursorBack(8)
				fmt.Printf(" >")
			} else if x > 1 {
				x--
				ansi.CursorBack(1)
				fmt.Printf(" ")
				ansi.CursorBack(2)
				ansi.CursorBack(4)
				fmt.Printf(" >")
			}
		} else if key == 13 {	//文字入力・編集処理

			if 10 * x + y == 71 {
				if i > 1 {
					str = str[:len(str) - 3]
					i--
					ansi.CursorNextLine(7 - y)
					ansi.CursorForward(9 + (i * 2) + i - 1)
					fmt.Printf("＊")
					ansi.CursorNextLine(0)
					ansi.CursorDown(7 - y)
					ansi.CursorForward((x - 1) * 4 + 2)
				}
			} else if 10 * x + y == 91 {
				if i > 1 {
					ansi.CursorBack(1)
					fmt.Printf(" ")
					ansi.CursorPreviousLine(1)
					fmt.Println("")
					fmt.Println("なまえは", str, "でいいですか？")
					fmt.Println("-----------")
					fmt.Println("|  はい　　|")
					fmt.Println("|  いいえ　|")
					fmt.Println("-----------")
					if chose(2) == 1 {
						fmt.Println("勇者", str,"爆誕！")
						fmt.Println("")
						return str
					} else {
						ansi.EraseInLine(1)
						ansi.CursorNextLine(1)
						ansi.EraseInLine(1)
						ansi.CursorNextLine(1)
						ansi.EraseInLine(1)
						ansi.CursorNextLine(1)
						ansi.EraseInLine(1)
						ansi.CursorNextLine(1)
						ansi.EraseInLine(1)
						ansi.CursorUp(6)
						fmt.Printf(" >")
						x = 1
						y = 6
					}
				}
			} else {
				switch 10 * x + y {
				case 16:
					str += "あ"
				case 26:
					str += "い"
				case 36:
					str += "う"
				case 46:
					str += "え"
				case 56:
					str += "お"
				case 66:
					str += "は"
				case 76:
					str += "ひ"
				case 86:
					str += "ふ"
				case 96:
					str += "へ"
				case 106:
					str += "ほ"
				case 15:
					str += "か"
				case 25:
					str += "き"
				case 35:
					str += "く"
				case 45:
					str += "け"
				case 55:
					str += "こ"
				case 65:
					str += "ま"
				case 75:
					str += "み"
				case 85:
					str += "む"
				case 95:
					str += "め"
				case 105:
					str += "も"
				case 14:
					str += "さ"
				case 24:
					str += "し"
				case 34:
					str += "す"
				case 44:
					str += "せ"
				case 54:
					str += "そ"
				case 64:
					str += "や"
				case 84:
					str += "ゆ"
				case 104:
					str += "よ"
				case 13:
					str += "た"
				case 23:
					str += "ち"
				case 33:
					str += "つ"
				case 43:
					str += "て"
				case 53:
					str += "と"
				case 63:
					str += "ら"
				case 73:
					str += "り"
				case 83:
					str += "る"
				case 93:
					str += "れ"
				case 103:
					str += "ろ"
				case 12:
					str += "な"
				case 22:
					str += "に"
				case 32:
					str += "ぬ"
				case 42:
					str += "ね"
				case 52:
					str += "の"
				case 62:
					str += "わ"
				case 82:
					str += "を"
				case 102:
					str += "ん"
				case 11:
					str += "っ"
				case 21:
					str += "ゃ"
				case 31:
					str += "ゅ"
				case 41:
					str += "ょ"
				case 51:
					str += "゛"
				case 61:
					str += "゜"
				}
				ansi.CursorNextLine(7 - y)
				ansi.CursorForward(9 + (i * 2) + i - 1)
				fmt.Printf("%s", str[3 * i - 3:3 * i])
				ansi.CursorNextLine(0)
				ansi.CursorDown(7 - y)
				ansi.CursorForward((x - 1) * 4 + 2)
				i++
			}

		}
	}

	return ""
}

//画面クリアする関数
func cls() {
	fmt.Print("\033[2J", "\033[1;1H")
}
