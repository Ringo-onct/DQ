package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Status struct {
	name   string
	hp     int
	mp     int
	atk    int
	def    int
	luk    int
	action int
}

type player struct {
	hp      int
	dmg     int
	atk     int
	atk_min int
}

type monster struct {
	hp      int
	dmg     int
	atk     int
	atk_min int
}

func main() {

	p_sta := new(Status)
	m_sta := new(Status)
	p := new(player) //この二つは仮置き。後で名前変える予定
	m := new(monster)

	//乱数発生
	rand.Seed(time.Now().UnixNano())

	//仮でそれぞれのステータス割り振る。これは後で別関数で読み込めるようにする。
	p.hp = 10
	p.atk = 8
	p.atk_min = 3
	m.hp = 50
	m.atk = 3
	m.atk_min = 1

	for true {

		//HP表示
		fmt.Println("---------------------")
		fmt.Printf("| PLAYER : %4d     |\n", p.hp)
		fmt.Printf("| MONSTER: %4d     |\n", m.hp)
		fmt.Println("---------------------")

		fmt.Println("0:にげる")
		fmt.Println("1:こうげき")
		fmt.Printf("行動の選択>")
		fmt.Scan(&p_sta.action)
		fmt.Printf("\n")	//見やすくするための改行
		//実際の行動
		//プレイヤーの行動
		switch p_sta.action {
		case 0:
			fmt.Println("にげだした。。")

		case 1:
			fmt.Println("プレイヤーのこうげき")
			//ダメージ計算
			p.dmg = rand.Intn(p.atk) + p.atk_min
			fmt.Printf("プレイヤーは%dのダメージをあたえた！\n", p.dmg)

		default:
			fmt.Println("こんらんしている")
		}

		//プレイヤーの行動の結果
		if p_sta.action == 0 {
			break
		} else if p_sta.action == 1 {
			m.hp -= p.dmg
		}

		//戦闘終了の判定
		if m.hp <= 0 {
			fmt.Println("モンスターをたおした！")
			break
		}

		//モンスターの行動
		m_sta.action = 1
		switch m_sta.action {
		case 1:
			fmt.Println("モンスターのこうげき")
			//ダメージの計算
			m.dmg = rand.Intn(m.atk) + m.atk_min
			fmt.Printf("モンスターは%dのダメージをあたえた。\n", m.dmg)

		default:
			fmt.Println("モンスターはようすをみている")
		}

		//モンスターの行動の結果
		if m_sta.action == 1 {
			p.hp -= m.dmg
		} else {
			break
		}

		//戦闘終了の判定
		if p.hp <= 0 {
			fmt.Println("プレイヤーはたおれた。。")
			break
		}

	}
}
