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
	m.atk = 1

	for true {

		//HP表示
		fmt.Println("---------------------")
		fmt.Println("| PLAYER : %4d     |", p.hp)
		fmt.Println("| MONSTER: %4d     |", m.hp)
		fmt.Println("---------------------")

		fmt.Println("0:にげる")
		fmt.Println("1:こうげき")
		fmt.Printf("行動の選択>")
		fmt.Scan(&p_sta.action)

		//実際の行動
		//プレイヤーの行動
		switch p_sta.action {
		case 0:
			fmt.Println("にげだした。。")

		case 1:
			fmt.Println("プレイヤーのこうげき")
			//ダメージ計算
			p.dmg = rand.intn(100)%p.atk + p.atk_min
			fmt.Println("プレイヤーは%dのダメージをあたえた！", p.dmg)

		default:
			fmt.Println("こんらんしている")
		}

		//プレイヤーの行動
		switch p_sta.action {
		//今ここ書いてた
		}

		//モンスターの行動
		fmt.Println("ようすをみている")

	}
}
