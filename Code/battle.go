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
	if float32(p_sta.hp) <= float32(p_sta.hp_max) * 0.25 {
		fmt.Print("\033[31m")
	} else {
		fmt.Print("\033[37m")
	}
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

		if float32(p_sta.hp) <= float32(p_sta.hp_max) * 0.25 {
			fmt.Print("\033[31m")
		} else {
			fmt.Print("\033[37m")
		}
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

	chose:
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
		fmt.Println("------------")
		fmt.Println("|  にげる　|")
		fmt.Println("|  こうげき|")
		if p_sta.lari >= 3 {	//魔法は、lv3以上じゃないと使えないので、表示するレベルを制限する。
			fmt.Println("|  まほう　|")
			fmt.Println("------------")
			action = chose(3)
		} else {
			fmt.Println("-----------")
			action = chose(2)
		}



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
		} else if action == 3 {	//魔法処理
			p_sta.dmg = 0
			i := 1
			cls()
			fmt.Println("")
			fmt.Println("-------------")
			fmt.Println("|  ホイミ　　|")
			switch {
			case p_sta.lari >= 4:
				fmt.Println("|  ギラ　　　|")
				i++
				fallthrough
			case p_sta.lari >= 7:
				fmt.Println("|  ラリホー　|")
				i++
				fallthrough
			case p_sta.lari >= 9:
				fmt.Println("|  レミーラ　|")
				i++
				fallthrough
			case p_sta.lari >= 10:
				fmt.Println("|  マホトーン|")
				i++
				fallthrough
			case p_sta.lari >= 12:
				fmt.Println("|  リレミト　|")
				i++
				fallthrough
			case p_sta.lari >= 13:
				fmt.Println("|  ルーラ　　|")
				i++
				fallthrough
			case p_sta.lari >= 15:
				fmt.Println("|  トヘロス　|")
				i++
				fallthrough
			case p_sta.lari >= 17:
				fmt.Println("|  ベホイミ　|")
				i++
				fallthrough
			case p_sta.lari >= 19:
				fmt.Println("|  ベギラマ　|")
				i++
			}
			fmt.Println("-------------")
			action = chose(i)
			cls()
			time.Sleep(1 * time.Second)

			//mp不足処理。長いのは勘弁して。
			if action == 1 && p_sta.mp < 4 || action == 2 && p_sta.mp < 2 || action == 3 && p_sta.mp < 4 || action == 4 && p_sta.mp < 3 || action == 5 && p_sta.mp < 2 || action == 6 && p_sta.mp < 6 || action == 7 && p_sta.mp < 8 || action == 8 && p_sta.mp < 2 || action == 9 && p_sta.mp < 10 || action == 10 && p_sta.mp < 5 {
				fmt.Println("mpが足りないようだ")
				goto chose	//選択し直し
			}

			switch action {
			case 1:
				p_sta.mp -= 4
				fmt.Printf("%sはホイミを使った！\n", p_sta.name)
				magic := rand.Intn(8) + 10
				time.Sleep(500 * time.Millisecond)

				if p_sta.hp + magic <= p_sta.hp_max {
					p_sta.hp += magic
				} else {
					magic = p_sta.hp_max - p_sta.hp
					p_sta.hp = p_sta.hp_max
				}

				fmt.Printf("hpが%d回復した！\n", magic)

			case 2:
				p_sta.mp -= 2
				fmt.Printf("%sはギラを放った！\n", p_sta.name)
				p_sta.dmg = rand.Intn(8) + 5
				time.Sleep(500 * time.Millisecond)
				if rand.Intn(16) + 1 <= m_sta.gira {
					fmt.Printf("%sはギラをひらりとかわした！\n", m_sta.name)
					time.Sleep(500 * time.Millisecond)
					fmt.Println("ダメージを　あたえられない！")
					p_sta.dmg = 0
				} else {
					fmt.Printf("%dのダメージをあたえた！\n", p_sta.dmg)
				}
				time.Sleep(500 * time.Millisecond)

			case 3:
				p_sta.mp -= 2
				fmt.Printf("%sはラリホーを放った！\n", p_sta.name)
				time.Sleep(500 * time.Millisecond)
				if rand.Intn(16) + 1 <= m_sta.lari {
					fmt.Printf("%sは気合で眠気を吹っ飛ばした！\n", m_sta.name)
					time.Sleep(500 * time.Millisecond)
				} else {
					fmt.Printf("%sはすやすやと眠った！", m_sta.name)
					m_sta.flag += 2	//一回目は寝たまま、二回目から起きる可能性がある1に変化するようにする
				}
				time.Sleep(500 * time.Millisecond)

			case 4:
				p_sta.mp -= 3
				fmt.Printf("%sはレミーラを使った！\n", p_sta.name)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("しかし　何も起こらなかった")
				time.Sleep(500 * time.Millisecond)

			case 5:
				p_sta.mp -= 2
				fmt.Printf("%sはマホトーンを放った！\n", p_sta.name)
				time.Sleep(500 * time.Millisecond)
				if rand.Intn(16) + 1 <= m_sta.mp {
					fmt.Printf("しかし、%sは魔法を見切っていたようだ！\n", m_sta.name)
					time.Sleep(500 * time.Millisecond)
					fmt.Println("マホトーンが防がれてしまった！")
					p_sta.dmg = 0
				} else {
					fmt.Printf("%sはまほうが使えなくなった！\n", m_sta.name)
					m_sta.flag += 20
				}
				time.Sleep(500 * time.Millisecond)

			case 6:
				p_sta.mp -= 6
				fmt.Printf("%sはリレミトを使った！\n", p_sta.name)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("しかし　何も起こらなかった")
				time.Sleep(500 * time.Millisecond)

			case 7:
				p_sta.mp -= 8
				fmt.Printf("%sはルーラを使った！\n", p_sta.name)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("しかし　何も起こらなかった")
				time.Sleep(500 * time.Millisecond)

			case 8:
				p_sta.mp -= 2
				fmt.Printf("%sはトヘロスを使った！\n", p_sta.name)
				time.Sleep(500 * time.Millisecond)
				fmt.Println("本来は弱い敵が居なくなるところだが")
				time.Sleep(130 * time.Millisecond)
				fmt.Println("まだ実装されていないようだ！")
				time.Sleep(500 * time.Millisecond)

			case 9:
				p_sta.mp -= 10
				fmt.Printf("%sはベホイミを使った！\n", p_sta.name)
				magic := rand.Intn(16) + 85
				time.Sleep(500 * time.Millisecond)

				if p_sta.hp + magic <= p_sta.hp_max {
					p_sta.hp += magic
				} else {
					magic = p_sta.hp_max - p_sta.hp
					p_sta.hp = p_sta.hp_max
				}

				fmt.Printf("hpが%d回復した！\n", magic)
				time.Sleep(500 * time.Millisecond)

			case 10:
				p_sta.mp -= 2
				fmt.Printf("%sはベギラマを放った！\n", p_sta.name)
				p_sta.dmg = rand.Intn(8) + 58
				time.Sleep(500 * time.Millisecond)
				if rand.Intn(16) + 1 <= m_sta.gira {
					fmt.Printf("%sはベギラマをひらりとかわした！\n", m_sta.name)
					time.Sleep(500 * time.Millisecond)
					fmt.Println("ダメージを　あたえられない！")
					p_sta.dmg = 0
				} else {
					fmt.Printf("%dのダメージをあたえた！\n", p_sta.dmg)
					time.Sleep(500 * time.Millisecond)
				}
			}
			m_sta.hp -= p_sta.dmg
		}

		if float32(p_sta.hp) <= float32(p_sta.hp_max) * 0.25 {
			fmt.Print("\033[31m")
		} else {
			fmt.Print("\033[37m")
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
					switch p_sta.lari {
					case 3:
						fmt.Println("まほう：ホイミを習得した！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("ホイミは、4mpを消費してhpを回復をするまほうだ！")
					case 4:
						fmt.Println("まほう：ギラを習得した！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("ギラは、2mpを消費して敵にダメージをあたえるまほうだ！")
					case 7:
						fmt.Println("まほう：ラリホーを習得した！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("ラリホーは2mpを消費してまれに敵を眠らせるまほうだ！")
					case 9:
						fmt.Println("まほう：レミーラを習得した！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("レミーラは3mpを消費して周りを照らすまほうだ！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("しかし、まだ実装してないから、間違えて使わないようにしような！")
					case 10:
						fmt.Println("まほう：マホトーンを習得した！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("マホトーンは2mpを消費してまれに敵のまほうを封じるまほうだ！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("しかし、まだ敵にまほうを実装していないようだ！")
					case 12:
						fmt.Println("まほう：リレミトを習得した！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("リレミトは6mpを消費してダンジョンから脱出するまほうだ！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("しかし、まだ実装してないから、間違えて使わないようにしような！")
					case 13:
						fmt.Println("まほう：ルーラを習得した！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("ルーラは8mpを消費してラダトームの城に戻るまほうだ！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("しかし、まだ実装してないから、間違えて使わないようにしような！")
					case 15:
						fmt.Println("まほう：トヘロスを習得した！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("トヘロスは2mpを消費して弱い敵を出さないようにするまほうだ！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("しかし、まだ実装してないから、間違えて使わないようにしような！")
					case 17:
						fmt.Println("まほう：ベホイミを習得した！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("ベホイミは10mpを消費してhpを大回復するまほうだ！")
					case 19:
						fmt.Println("まほう：ベギラマを習得した！")
						time.Sleep(500 * time.Millisecond)
						fmt.Println("ベギラマは5mpを消費して敵に大ダメージを与えるまほうだ！")
					}
					getkey()
					} else {
					break
				}
			}
			break
		}

		time.Sleep(1 * time.Second)

		//モンスターの行動
		cls()
		m_sta.dmg = 0
		if m_sta.flag % 10 >= 2 {
			m_sta.flag = m_sta.flag - (m_sta.flag % 10) + 1
		} else if m_sta.flag % 10 == 1 && rand.Intn(3) == 0{
			m_sta.flag = m_sta.flag - (m_sta.flag % 10)
			fmt.Printf("%sは目を覚ました！\n", m_sta.name)
		} else if m_sta.flag % 10 == 1 {
			fmt.Printf("%sは戦闘中にもかかわらず眠っているようだ。\n", m_sta.name)
		}

		if m_sta.flag % 10 == 0 {
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
		}

		if float32(p_sta.hp) <= float32(p_sta.hp_max) * 0.25 {
			fmt.Print("\033[31m")
		} else {
			fmt.Print("\033[37m")
		}

		//勝敗判定
		if p_sta.hp <= 0 {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("")
			fmt.Println(p_sta.name, "はたおれた。。")
			time.Sleep(2 * time.Second)
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
		fmt.Print("\033[37m")	//色を元に戻す
		return 1
	}

	return 0
}
