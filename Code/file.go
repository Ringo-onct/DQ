package main
import (
	"bufio"
	//"fmt"
	"os"
	"strconv"
	"strings"
)

func file (f allfile) {
	f.datafile()
}

func (p *player) datafile() {

	filepass, err := os.Open("../Document/player_list.txt")	//fopen的な何か
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

		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (m *monster) datafile() {

	filepass, err := os.Open("../Document/monster_list.txt")	//fopen的な何か
	if err != nil {
		panic(err)
	}
	defer filepass.Close()

	scanner := bufio.NewScanner(filepass)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "mob1") {
			fields := strings.Split(line, ",")
			m.hp, _ = strconv.Atoi(fields[1])
			m.atk, _ = strconv.Atoi(fields[2])
			m.atk_min, _ = strconv.Atoi(fields[3])

		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
