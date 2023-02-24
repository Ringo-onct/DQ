package main
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"math/rand"
	"time"
	"fmt"
	"log"
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
			p_sta.hp, _ = strconv.Atoi(fields[4])
			p_sta.hp_max, _ = strconv.Atoi(fields[5])
			p_sta.mp, _ = strconv.Atoi(fields[6])
			p_sta.exp, _ = strconv.Atoi(fields[7])
			p_sta.gold, _ = strconv.Atoi(fields[8])
			p_sta.lari, _ = strconv.Atoi(fields[9])
			p_sta.gira, _ = strconv.Atoi(fields[10])
			p_sta.avo, _ = strconv.Atoi(fields[11])
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
	rand.Seed(time.Now().UnixNano())
	seed := "a" + strconv.Itoa(rand.Intn(2) + 1)
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

func save(p_sta *status, line int) {
	file, err := os.Open("../Document/player_list")
	var (
		upline		string
		downline	string
		i 			int
	)
	addline := "a" + strconv.Itoa(line) + "," + p_sta.name + "," + strconv.Itoa(p_sta.atk) + "," + strconv.Itoa(p_sta.dif) + "," + strconv.Itoa(p_sta.hp) + "," + strconv.Itoa(p_sta.hp_max) + "," + strconv.Itoa(p_sta.mp) + "," + strconv.Itoa(p_sta.exp) + "," + strconv.Itoa(p_sta.gold) + "," + strconv.Itoa(p_sta.lari) + "," + strconv.Itoa(p_sta.gira) + "," + strconv.Itoa(p_sta.avo) + "\n"
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i = 0; scanner.Scan() && i < (line - 1); i++ { // 選択したデータの前のデータまで読み込んでuplineに保存する。
		upline += scanner.Text()

		if i != line - 1 {
			upline += "\n"
		}
	}

	i = 0	//カウント用変数を初期化
	for scanner.Scan() {
		i++
		if i >= 1 {
			downline += scanner.Text() + "\n"
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// 指定したファイルをオープンして、中身をクリアする。
	file, err = os.OpenFile("../Document/player_list", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(upline)
	file.WriteString(addline)
	file.WriteString(downline)
}

func makedata(line int) {
	rand.Seed(time.Now().UnixNano())
	file, err := os.OpenFile("../Document/player_list", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		//エラー処理
		log.Fatal(err)
	}
	defer file.Close()
	name := string(namewrite())
	name = "a" + strconv.Itoa(line + 1) + "," + name + ",4,4,15,15,0,0,120,1,0," + strconv.Itoa(rand.Intn(16))
	fmt.Fprintln(file, name)
}

func delldata(p_sta *[]status, line int) {

	fmt.Println("どのデータを消しますか？")
	var (
		x, i	int
		s		string
	)
	fmt.Println("--------------------------------")

	for i = 0; i < line; i++ {
		s = ""
		x = 6 - (len((*p_sta)[i].name) / 3)
		for x > 0 {
			s += "　"
			x--
		}
		fmt.Printf("|    %s%s|HP:%-3d|Lv:%-3d|\n", (*p_sta)[i].name, s, (*p_sta)[i].hp, (*p_sta)[i].lari)
	}
	fmt.Println("--------------------------------")

	fmt.Println("")
	pl := chose(line - 2, 1)

	file, err := os.OpenFile("../Document/player_list", os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i = 0; i < line - 1; i++ {

		if i < pl {
			s = "a" + strconv.Itoa(i + 1) + "," + (*p_sta)[i].name + "," + strconv.Itoa((*p_sta)[i].atk) + "," + strconv.Itoa((*p_sta)[i].dif) + "," + strconv.Itoa((*p_sta)[i].hp) + "," + strconv.Itoa((*p_sta)[i].hp_max) + "," + strconv.Itoa((*p_sta)[i].mp) + "," + strconv.Itoa((*p_sta)[i].exp) + "," + strconv.Itoa((*p_sta)[i].gold) + "," + strconv.Itoa((*p_sta)[i].lari) + "," + strconv.Itoa((*p_sta)[i].gira) + "," + strconv.Itoa((*p_sta)[i].lari)
		} else if i == pl {
			i++
			s = "a" + strconv.Itoa(i) + "," + (*p_sta)[i].name + "," + strconv.Itoa((*p_sta)[i].atk) + "," + strconv.Itoa((*p_sta)[i].dif) + "," + strconv.Itoa((*p_sta)[i].hp) + "," + strconv.Itoa((*p_sta)[i].hp_max) + "," + strconv.Itoa((*p_sta)[i].mp) + "," + strconv.Itoa((*p_sta)[i].exp) + "," + strconv.Itoa((*p_sta)[i].gold) + "," + strconv.Itoa((*p_sta)[i].lari) + "," + strconv.Itoa((*p_sta)[i].gira) + "," + strconv.Itoa((*p_sta)[i].lari)
		} else {
			s = "a" + strconv.Itoa(i) + "," + (*p_sta)[i].name + "," + strconv.Itoa((*p_sta)[i].atk) + "," + strconv.Itoa((*p_sta)[i].dif) + "," + strconv.Itoa((*p_sta)[i].hp) + "," + strconv.Itoa((*p_sta)[i].hp_max) + "," + strconv.Itoa((*p_sta)[i].mp) + "," + strconv.Itoa((*p_sta)[i].exp) + "," + strconv.Itoa((*p_sta)[i].gold) + "," + strconv.Itoa((*p_sta)[i].lari) + "," + strconv.Itoa((*p_sta)[i].gira) + "," + strconv.Itoa((*p_sta)[i].lari)
		}
		fmt.Fprintln(file, s)
	}
}

func lvup(p_sta *status) {
	if p_sta.lari == 1 {	//タイプ別レベルアップ処理
		if p_sta.avo % 2 == 0 {
			p_sta.atk = int(float32(p_sta.atk) * 0.9) + p_sta.avo / 4
			if p_sta.avo % 4 == 0 {
				p_sta.dif = int(float32(p_sta.dif) * 0.9) + p_sta.avo / 4
			} else {
				p_sta.hp_max = int(float32(p_sta.hp_max) * 0.9) + p_sta.avo / 4
			}
		} else {
			p_sta.gira = int(float32(p_sta.gira) * 0.9)
			if p_sta.lari > 2 {
				p_sta.gira += p_sta.avo / 4
			}

			if p_sta.avo % 4 == 1 {
				p_sta.dif = int(float32(p_sta.dif) * 0.9) + p_sta.avo / 4
			} else {
				p_sta.hp_max = int(float32(p_sta.hp_max) * 0.9) + p_sta.avo / 4
			}
		}
		p_sta.hp = p_sta.hp_max
		p_sta.mp = p_sta.gira
	} else {
		var seed string
		filepass, err := os.Open("../Document/lvup_status")	//fopen的な何か
		if err != nil {
			panic(err)
		}
		defer filepass.Close()

		scanner := bufio.NewScanner(filepass)
		for scanner.Scan() {
			line := scanner.Text()
			seed = "a" + strconv.Itoa(p_sta.lari)
			if strings.Contains(line, seed) {
				fields := strings.Split(line, ",")
				p_sta.atk, _ = strconv.Atoi(fields[1])
				p_sta.dif, _ = strconv.Atoi(fields[2])
				p_sta.hp_max, _ = strconv.Atoi(fields[3])
				p_sta.gira, _ = strconv.Atoi(fields[4])
			}
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}

		//タイプ別レベルアップ処理
		if p_sta.avo % 2 == 0 {
			p_sta.atk = int(float32(p_sta.atk) * 0.9) + p_sta.atk / 4
			if p_sta.avo % 4 == 0 {
				p_sta.dif = int(float32(p_sta.dif) * 0.9) + p_sta.dif / 4
			} else {
				p_sta.hp_max = int(float32(p_sta.hp_max) * 0.9) + p_sta.hp_max / 4
			}
		} else {
			p_sta.gira = int(float32(p_sta.gira) * 0.9)
			if p_sta.lari > 2 {
				p_sta.gira += p_sta.gira / 4
			}

			if p_sta.avo % 4 == 1 {
				p_sta.dif = int(float32(p_sta.dif) * 0.9) + p_sta.dif / 4
			} else {
				p_sta.hp_max = int(float32(p_sta.hp_max) * 0.9) + p_sta.hp_max / 4

			}
		}
		p_sta.hp = p_sta.hp_max
		p_sta.mp = p_sta.gira
	}
}
