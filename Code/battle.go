//ここに戦闘関連の処理を置く。
package main
import (
	"time"
	"fmt"
	"math/rand"
)

func battle(p_sta *status, m_sta *status) int{
	var action int

	rand.Seed(time.Now().UnixNano())	//乱数設定

	cls()
	time.Sleep(1 * time.Second)

	fileM(m_sta)	//monsterデータ読み込み

	//エンカウント処理
	if (m_sta.dif * rand.Intn(64)) > (p_sta.dif * rand.Intn(256)) {	//monsterの先制攻撃
		fmt.Printf("%sは、%sがみがまえるまえにおそってきた！\n", m_sta.name, p_sta.name)

		fmt.Println(m_sta.name, "のこうげき")
		time.Sleep(500 * time.Millisecond)
		//ダメージの計算
		if (m_sta.atk - p_sta.dif / 4 ) >= m_sta.atk / 2 + 1 {
			m_sta.dmg = (rand.Intn(256) * (m_sta.atk - p_sta.dif / 4 + 1) / 256 + m_sta.atk - p_sta.dif / 4) / 4
		} else if m_sta.atk - p_sta.dif / 2 < 0 {
			m_sta.dmg = rand.Intn(256) * (m_sta.atk / 2 + 1) / 256 + 2
		} else {
			m_sta.dmg = rand.Intn(256) * (m_sta.atk / 2 + 1) / 256 + 2
		}

		if m_sta.dmg == 0 {
			fmt.Println("ミス！")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("ダメージをうけない！")
		} else {
			fmt.Printf("%sは%dのダメージをうけた！\n", p_sta.name, m_sta.dmg)
		}
		//ダメージ処理
		p_sta.hp -= m_sta.dmg

		//勝敗判定
		if p_sta.hp <= 0 {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("")
			fmt.Println(p_sta.name, "はたおれた。。")
			return 1
		}

	} else {
		fmt.Println(m_sta.name,"があわられた！！")
	}

	time.Sleep(2 * time.Second)

	for true {	//戦闘処理
		cls()	//コンソール画面クリア

		//体力表示
		s1, s2 := "", ""

		x := 6 - (len(m_sta.name) / 3)
		for x > 0 {
			s1 += "　"
			x--
		}
		x = 6 - (len(p_sta.name) / 3)
		for x > 0 {
			s2 += "　"
			x--
		}
		fmt.Println("---------------------")
		fmt.Printf("| %s%s: %-3d |\n", p_sta.name, s2, p_sta.hp)
		fmt.Printf("| %s%s: %-3d |\n", m_sta.name, s1, m_sta.hp)
		fmt.Println("---------------------")

		//プレイヤーの行動選択
		fmt.Println("")
		fmt.Println("-----------")
		fmt.Println("|  にげる　|")
		fmt.Println("|  こうげき|")
		fmt.Println("-----------")
		action = chose(2)

		time.Sleep(1 * time.Second)
		cls()

		if action == 1 {		//戦闘離脱
			fmt.Printf("%sはにげだした。。\n", p_sta.name)
			time.Sleep(1 * time.Second)
			cls()
			return 1
		} else if action == 2 {	//戦闘後のダメージ処理
			fmt.Println(p_sta.name, "のこうげき")
			time.Sleep(500 * time.Millisecond)
			//ダメージ計算
			p_sta.dmg = (rand.Intn(256) * (p_sta.atk - m_sta.dif / 2 + 1) / 256 + p_sta.atk - m_sta.dif / 2) / 4
			if rand.Intn(32) == 0 {	//1 / 32の確立
				fmt.Println("！！！！かいしんのいちげき！！！！")
				time.Sleep(500 * time.Millisecond)
				p_sta.dmg += p_sta.atk - ((p_sta.atk / 2) * rand.Intn(256)) / 256
			}

			if p_sta.dmg == 0 {
				fmt.Println("ミス！")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("ダメージを　あたえられない！")
			} else if rand.Intn(65) < m_sta.avo {
				fmt.Printf("%sがこうげきをよけた！\n", m_sta.name)
				p_sta.dmg = 0
			} else {
				fmt.Printf("%sは%dのダメージをあたえた！\n", p_sta.name, p_sta.dmg)
			}

			//ダメージ処理
			m_sta.hp -= p_sta.dmg
		}

		//勝敗判定
		if m_sta.hp <= 0 {		//勝利処理
			fmt.Println(m_sta.name, "をたおした！")
			time.Sleep(500 * time.Millisecond)
			if p_sta.lari == 30 {
				fmt.Printf("%d Goldを手に入れた！", m_sta.gold)
			} else {
				fmt.Printf("%d Goldと%d Expを手に入れた！\n", m_sta.gold, m_sta.exp)
			}

			p_sta.exp += m_sta.exp
			p_sta.gold += m_sta.gold
			if p_sta.exp > 65535 {
				p_sta.exp = 65535
			}
			//レベルアップ確認
			time.Sleep(500 * time.Millisecond)
			for {
				n := p_sta.lari + 1
				if (3 * (n - 2) * (n - 2) * (n - 2) + 7) <= p_sta.exp {	//レベルアップ処理
					p_sta.lari++
					lvup(p_sta)
					fmt.Printf("%sは、Lv%dにレベルアップした！\n", p_sta.name, p_sta.lari)
					} else {
					break
				}
			}
			break
		}

		time.Sleep(1 * time.Second)

		//モンスターの行動
		fmt.Println(m_sta.name, "のこうげき")
		time.Sleep(500 * time.Millisecond)
		//ダメージの計算
		if (m_sta.atk - p_sta.dif / 4 ) >= m_sta.atk / 2 + 1 {
			m_sta.dmg = (rand.Intn(256) * (m_sta.atk - p_sta.dif / 4 + 1) / 256 + m_sta.atk - p_sta.dif / 4) / 4
		} else if m_sta.atk - p_sta.dif / 2 < 0 {
			m_sta.dmg = rand.Intn(256) * (m_sta.atk / 2 + 1) / 256 + 2
		} else {
			m_sta.dmg = rand.Intn(256) * (m_sta.atk / 2 + 1) / 256 + 2
		}

		if m_sta.dmg == 0 {
			fmt.Println("ミス！")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("ダメージをうけない！")
		} else {
			fmt.Printf("%sは%dのダメージをうけた！\n", p_sta.name, m_sta.dmg)
		}
		//ダメージ処理
		p_sta.hp -= m_sta.dmg

		//勝敗判定
		if p_sta.hp <= 0 {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("")
			fmt.Println(p_sta.name, "はたおれた。。")
			return 1
		}

		time.Sleep(3 * time.Second)
	}

	fmt.Println("")
	fmt.Println("-----------")
	fmt.Println("|  やめる　|")
	fmt.Println("|  つづける|")
	fmt.Println("-----------")
	if chose(2) == 1 {
		return 1
	}

	return 0
}
