package main
import (
	"bufio"
	//"fmt"
	"os"
	"strconv"
	"strings"
)

func file (f allfile) {
	f.playerfile()
	//f.monsterfile()
}

func (g *player) playerfile() {

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
			g.hp, _ = strconv.Atoi(fields[1])
			g.atk, _ = strconv.Atoi(fields[2])
			g.atk_min, _ = strconv.Atoi(fields[3])

		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
