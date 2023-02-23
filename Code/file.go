package main
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"math/rand"
	"time"
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
			p_sta.atk, _ = strconv.Atoi(fields[2])
			p_sta.dif, _ = strconv.Atoi(fields[3])
			p_sta.hp_max, _ = strconv.Atoi(fields[4])
			p_sta.mp, _ = strconv.Atoi(fields[5])
			p_sta.exp, _ = strconv.Atoi(fields[6])
			p_sta.gold, _ = strconv.Atoi(fields[7])
			p_sta.lari, _ = strconv.Atoi(fields[8])
			p_sta.gira, _ = strconv.Atoi(string((fields[0])[1]))	//文字列から数字のみ取り出し、int型に直して格納
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	p_sta.hp = p_sta.hp_max
}

func fileM(m_sta *status) {

	filepass, err := os.Open("../Document/monster_list")	//fopen的な何か
	if err != nil {
		panic(err)
	}
	defer filepass.Close()
	rand.Seed(time.Now().UnixNano())
	seed := "a" + strconv.Itoa(rand.Intn(2))
	scanner := bufio.NewScanner(filepass)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, seed) {
			fields := strings.Split(line, ",")
			m_sta.name = fields[1]
			m_sta.atk, _ = strconv.Atoi(fields[2])
			m_sta.dif, _ = strconv.Atoi(fields[3])
			m_sta.hp_max, _ = strconv.Atoi(fields[4])
			m_sta.mp, _ = strconv.Atoi(fields[5])
			m_sta.exp, _ = strconv.Atoi(fields[6])
			m_sta.gold, _ = strconv.Atoi(fields[7])
			m_sta.lari, _ = strconv.Atoi(fields[8])
			m_sta.gira, _ = strconv.Atoi(fields[9])
			m_sta.avo, _ = strconv.Atoi(fields[10])
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	m_sta.hp = m_sta.hp_max - m_sta.hp_max * rand.Intn(256) / 1024	//モンスターの初期HP設定

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
