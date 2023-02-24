package main

import (
	"time"
	"os"
)

type status struct {	//小文字にしたら、Goのパッケージ内の関数に小文字から始まる関数内から、被らないらしい。
	name	string	//名前
	atk		int		//攻撃力 (pの攻撃力は、武器の攻撃力 + atk)
	dmg		int		//与えるダメージ
	dif		int		//守備力・素早さ (pの守備力は、防具の防御力 + dif / 2)
	hp		int		//体力
	hp_max	int		//最大体力
	mp		int		//pはそのままmp,mはマホの回避率
	exp		int		//pは総Exp,mは獲得exp
	gold	int		//pは総gold,mは獲得gold
	lari	int		//pはレベル,mはラリの回避率
	gira	int		//pは最大mp,mはギラの回避率
	avo		int		//pは成長タイプ。mは回避率
}

func main() {
	top:
	line := linecountP()
	p_sta := make([]status, line + 1)
	//一応、今作ってるのはplayerのデータ読み込みだから、monsterはまだ配列対応させない。
	var (
		m_sta status
		action int
	)
	//ゲーム開始待機時のコンソール画面表示
	console(&p_sta[0], &m_sta, 5)
	console(&p_sta[0], &m_sta, 0)


	for i := 0; i < line; i++ {	//ファイルにある分のplayerデータ読み込み
		fileP(&p_sta[i], i + 1)
	}

	//playerデータ表示
	player_UI(&p_sta, line)

	//player選択
	pl := prompt(&p_sta[0], -line) - 1
	if pl == line {				//新規作成
		time.Sleep(1 * time.Second)
		console(&p_sta[0], &m_sta, 0)
		makedata(line)
		fileP(&p_sta[line], line + 1)
		lvup(&p_sta[line])
		save(&p_sta[line], pl + 1)
		time.Sleep(1 * time.Second)
	} else if pl == line + 1 {	//データ削除
		time.Sleep(1 * time.Second)
		console(&p_sta[0], &m_sta, 0)
		delldata(&p_sta, line)
		if prompt(&p_sta[0], 1) == 0 {
			os.Exit(1)
		} else {
			console(&p_sta[0], &m_sta, 0)
			goto top
		}
	}

	console(&p_sta[pl], &m_sta, 0)


	fileP(&p_sta[pl], pl + 1)	//playerデータ読み込みここに置くと再読み込みさせないで体力保持できる
	time.Sleep(1 * time.Second)

	for true {	//戦闘継続ループ
		console(&p_sta[pl], &m_sta, 0)
		time.Sleep(1 * time.Second)
		fileM(&m_sta)	//monsterデータ読み込み
		if console(&p_sta[pl], &m_sta, 3) == 3 {	//先制攻撃処理
			actionM(&p_sta[pl], &m_sta)
			p_sta[pl].hp -= m_sta.dmg
		}
		time.Sleep(2 * time.Second)

		for true {	//戦闘処理
			console(&p_sta[pl], &m_sta, 0)	//コンソール画面クリア
			console(&p_sta[pl], &m_sta, 1)	//体力表示

			//playerの行動選択

			//プレイヤーの行動選択・結果
			action = prompt(&p_sta[pl], 2)

			time.Sleep(1 * time.Second)
			console(&p_sta[pl], &m_sta, 0)

			if action == 0 {		//戦闘離脱
				actionP(&p_sta[pl], &m_sta, 0)
				time.Sleep(1 * time.Second)
				console(&p_sta[pl], &m_sta, 0)
				break
			} else if action == 1 {	//戦闘後のダメージ処理
				actionP(&p_sta[pl], &m_sta, 1)
				m_sta.hp -= p_sta[pl].dmg
			}

			//勝敗判定
			if console(&p_sta[pl], &m_sta, 2) == 1 {	//勝ち
				p_sta[pl].exp += m_sta.exp
				p_sta[pl].gold += m_sta.gold
				//レベルアップ確認
				time.Sleep(500 * time.Millisecond)
				for {
					n := p_sta[pl].lari + 1
					if (3 * (n - 2) * (n - 2) * (n - 2) + 7) <= p_sta[pl].exp {	//レベルアップ処理
						p_sta[pl].lari++
						lvup(&p_sta[pl])
						console(&p_sta[pl], &m_sta, 6)
					} else {
						break
					}
				}


				break
			}

			time.Sleep(1 * time.Second)
			//モンスターの行動
			actionM(&p_sta[pl], &m_sta)

			//モンスターの行動の結果
			p_sta[pl].hp -= m_sta.dmg

			//勝敗判定
			if console(&p_sta[pl], &m_sta, 2) == 2 {	//負け
				time.Sleep(1 * time.Second)
				break
			}

			time.Sleep(3 * time.Second)
		}

		if p_sta[pl].hp <= 0 {
			break
		} else if prompt(&p_sta[pl], 1) == 0 {
			break
		}
	}
	console(&p_sta[pl], &m_sta, 0)
	//保存用処理
	save(&p_sta[pl], pl + 1)
	//終了メッセージ
	console(&p_sta[pl], &m_sta, 4)

	time.Sleep(2 * time.Second)
}
