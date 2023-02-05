package main
import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func fileP(p_sta *status) {

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
			p_sta.hp, _ = strconv.Atoi(fields[1])
			p_sta.atk, _ = strconv.Atoi(fields[2])
			p_sta.atk_min, _ = strconv.Atoi(fields[3])

		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func fileM(m_sta *status) {

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
			m_sta.hp, _ = strconv.Atoi(fields[1])
			m_sta.atk, _ = strconv.Atoi(fields[2])
			m_sta.atk_min, _ = strconv.Atoi(fields[3])

		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
