package main

import (
	"time"
)

type status struct {	//小文字にしたら、Goのパッケージ内の関数に小文字から始まる関数内から、被らないらしい。
	name	string	//名前
	hp		int		//体力
	mp		int		//MP
	atk		int		//攻撃力
	atk_min int		//最低攻撃力
	def		int		//防御力
	luk		int		//運
	action	int		//行動
	dmg		int		//与えたダメージ
}

func main() {
	//この形にしろってchatgptに言われた
	var p_sta status
	var m_sta status

	fileP(&p_sta)	//playerデータ読み込みここに置くと再読み込みさせないで体力保持できる

	for true {	//戦闘継続ループ
	console(&p_sta, &m_sta, 0)
	time.Sleep(1 * time.Second)
	fileM(&m_sta)	//monsterデータ読み込み
	console(&p_sta, &m_sta, 3)
	time.Sleep(2 * time.Second)

	for true {

		console(&p_sta, &m_sta, 0)	//コンソール画面クリア
		console(&p_sta, &m_sta, 1)	//体力表示
		prompt(&p_sta, 1)	//行動選択

		time.Sleep(1 * time.Second)

		console(&p_sta, &m_sta, 0)	//コンソール画面クリア

		//playerの行動
		actionP(&p_sta)

		//プレイヤーの行動の結果
		if p_sta.action == 0 {
			time.Sleep(1 * time.Second)
			console(&p_sta, &m_sta, 0)
			break
		} else if p_sta.action == 1 {
			m_sta.hp -= p_sta.dmg
		}

		//勝敗判定
		if console(&p_sta, &m_sta, 2) == 1 {
			break
		}

		//モンスターの行動
		m_sta.action = 1
		actionM(&m_sta)

		//モンスターの行動の結果
		if m_sta.action == 1 {
			p_sta.hp -= m_sta.dmg
		} else {
			break
		}

		//勝敗判定
		if console(&p_sta, &m_sta, 2) == 2 {
			break
		}

		time.Sleep(3 * time.Second)

	}

	if prompt(&p_sta, 0) == 0 {
		break
	} else {
		p_sta.hp += 10
	}
	}
	console(&p_sta, &m_sta, 0)
	console(&p_sta, &m_sta, 4)
	time.Sleep(2 * time.Second)
}
