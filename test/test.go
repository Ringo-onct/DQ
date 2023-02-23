package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("../Document/test")
	var (
		uptext		string
		downtext	string
	)
	startLine := 1 // 4行目から読み込む例
    lineCount := 0
	line := 2
	addline := "test_text"
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for i := 0; scanner.Scan() && i < line; i++ { // 2行目まで読み込む
        uptext += scanner.Text()

		if i != line - 1 {
			uptext += "\n"
		}
    }

    for scanner.Scan() {
        lineCount++
        if lineCount >= startLine {
            downtext += scanner.Text() + "\n"
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
    fmt.Printf("%s\n%s\n%s", uptext, addline, downtext)


    // 指定したファイルをオープンして、中身をクリアする。
    file, err = os.OpenFile("../Document/test", os.O_TRUNC|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()


	addline += "\n"
	uptext += "\n"
    file.WriteString(uptext)
    file.WriteString(addline)
    file.WriteString(downtext)
}
