package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func file(p *player) {

	filepass, err := os.Open("../Document/player.txt")	//fopen的な何か
	if err != nil {
		panic(err)
	}
	defer filepass.Close()

	scanner := bufio.NewScanner(filepass)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "user1") {
			fields := strings.Split(line, ",")
			p.hp, _ = strconv.Atoi(fields[1])
			p.atk, _ = strconv.Atoi(fields[2])
			p.atk_min, _ = strconv.Atoi(fields[3])

			fmt.Printf("hp: %d, atk: %d, atk_min = %d\n", p.hp, p.atk, p.atk_min)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
