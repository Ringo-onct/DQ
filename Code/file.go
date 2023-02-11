package main
import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func fileP(p_sta *status, i int) {
	var seed string
	filepass, err := os.Open("../Document/player_list")	//fopen的な何か
	if err != nil {
		panic(err)
	}
	defer filepass.Close()

	scanner := bufio.NewScanner(filepass)
	for scanner.Scan() {
		line := scanner.Text()
		seed = "a" + strconv.Itoa(i)
		if strings.Contains(line, seed) {
			fields := strings.Split(line, ",")
			p_sta.name = fields[1]
			p_sta.hp, _ = strconv.Atoi(fields[2])
			p_sta.atk, _ = strconv.Atoi(fields[3])
			p_sta.atk_min, _ = strconv.Atoi(fields[4])
			p_sta.def, _ = strconv.Atoi(fields[5])
			p_sta.luk, _ = strconv.Atoi(fields[6])
			p_sta.mp, _ = strconv.Atoi(fields[7])

		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func fileM(m_sta *status) {

	filepass, err := os.Open("../Document/monster_list")	//fopen的な何か
	if err != nil {
		panic(err)
	}
	defer filepass.Close()
	seed := Math(1)
	scanner := bufio.NewScanner(filepass)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, seed) {
			fields := strings.Split(line, ",")
			m_sta.name = fields[1]
			m_sta.hp, _ = strconv.Atoi(fields[2])
			m_sta.atk, _ = strconv.Atoi(fields[3])
			m_sta.atk_min, _ = strconv.Atoi(fields[4])
			m_sta.def, _ = strconv.Atoi(fields[5])
			m_sta.luk, _ = strconv.Atoi(fields[6])
			m_sta.mp, _ = strconv.Atoi(fields[7])	//monsterのMPはとして扱う。

		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

//playerデータの量(行数)を調べる
func linecountP() int {
	filepass, err := os.Open("../Document/player_list")
	if err != nil {
		panic(err)
	}
	defer filepass.Close()

	scanner := bufio.NewScanner(filepass)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}
