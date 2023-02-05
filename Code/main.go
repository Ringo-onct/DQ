package main

import (
	"fmt"
	"math/rand"
	"time"
)

type status struct { //小文字にしたら、Goのパッケージ内の関数に小文字から始まる関数内から、被らないらしい。
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

type allfile interface {
	datafile()
}

func main() {
	//この形にしろってchatgptに言われた
	p_sta := status{}
	m_sta := status{}
	var p player
	var m monster

	file(&p) //playerデータ読み込み
	file(&m) //monsterデータ読み込み

	for true {

		console(&p, &m, 0) //コンソール画面クリア
		console(&p, &m, 1) //体力表示
		prompt(&p_sta, 1)  //行動選択

		time.Sleep(1 * time.Second)

		console(&p, &m, 0) //コンソール画面クリア

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

		time.Sleep(3 * time.Second)

	}
}
