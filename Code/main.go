package main

import (
	"time"
	"os"
	"fmt"
	"github.com/k0kubun/go-ansi"
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
	var m_sta status

	//ゲーム開始待機時のコンソール画面表示
	fmt.Printf("Press ENTER to start")
	ansi.CursorHide()
	for {	//エンター待機
		x := getkey()
		if x == 13 {
			break
		}
	}

	cls()
	for i := 0; i < line; i++ {	//ファイルにある分のplayerデータ読み込み
		fileP(&p_sta[i], i + 1)
	}

	//playerデータ表示
	fmt.Println("--------------------------------")
	for i := 0; i < line; i++ {
		s := ""
		x := 6 - (len(p_sta[i].name) / 3)
		for x > 0 {
			s += "　"
			x--
		}
		fmt.Printf("|    %s%s|HP:%-3d|Lv:%-3d|\n", p_sta[i].name, s, p_sta[i].hp, p_sta[i].lari)
	}
	fmt.Println("|    ぼうけんのしょをつくる　　|")
	fmt.Println("|    ぼうけんのしょをけす　　　|")
	fmt.Println("--------------------------------")

	//player選択
	pl := chose(line + 2) - 1	//配列に使うため-1している。

	if pl == line {				//新規作成
		time.Sleep(1 * time.Second)
		cls()
		makedata(line)
		fileP(&p_sta[line], line + 1)
		lvup(&p_sta[line])
		save(&p_sta[line], pl + 1)
		time.Sleep(1 * time.Second)
	} else if pl == line + 1 {	//データ削除
		time.Sleep(1 * time.Second)
		cls()
		delldata(&p_sta, line)
		fmt.Println("")
		fmt.Println("-----------")
		fmt.Println("|  やめる　|")
		fmt.Println("|  つづける|")
		fmt.Println("-----------")
		if chose(2) == 1 {
			os.Exit(1)
		} else {
			cls()
			goto top
		}
	}

	cls()

	fileP(&p_sta[pl], pl + 1)	//playerデータ読み込みここに置くと再読み込みさせないで体力保持できる
	time.Sleep(1 * time.Second)

	for true {	//戦闘継続ループ
		if battle(&p_sta[pl], &m_sta) == 1 {	//戦闘終了処理
			break
		}
	}

	cls()
	//保存用処理
	save(&p_sta[pl], pl + 1)
	//終了メッセージ
	str := "おつかれさまでした。"
	for _, char1 := range str {
		fmt.Printf("%c", char1)
		time.Sleep(130 * time.Millisecond)
	}

	fmt.Println("")
	time.Sleep(500 * time.Millisecond)
	str = "りせっとぼたんを　おしながら"
	for _, char2 := range str {
		fmt.Printf("%c", char2)
		time.Sleep(130 * time.Millisecond)
	}
	fmt.Println("")
	time.Sleep(500 * time.Millisecond)
	str = "でんげんを　きってください"
	for _, char3 := range str {
		fmt.Printf("%c", char3)
		time.Sleep(130 * time.Millisecond)
	}
	fmt.Println("")

	time.Sleep(2 * time.Second)
}
