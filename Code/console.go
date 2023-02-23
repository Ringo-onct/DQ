package main
import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"os/exec"
	"runtime"
	"github.com/k0kubun/go-ansi"
)

func console(p_sta *status, m_sta *status, mode int) int {	//何かしらの表示
	switch mode {
		case 0:	//コンソール画面クリア
			os_which := runtime.GOOS
			switch os_which {
				case "windows":
					cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
					cmd.Stdout = os.Stdout
					cmd.Run()

				case "linux":
					cmd := exec.Command("clear") //Linux example, its tested
					cmd.Stdout = os.Stdout
					cmd.Run()
				}

		case 1:	//体力表示
			var (
				s1, s2 string
				x int
			)

			x = 6 - (len(m_sta.name) / 3)
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

		case 2:	//戦闘終了判定
			if m_sta.hp <= 0 {
				fmt.Println("モンスターをたおした！")
				return 1	//勝利
			} else if p_sta.hp <= 0 {
				time.Sleep(500 * time.Millisecond)
				fmt.Println("")
				fmt.Println("プレイヤーはたおれた。。")
				return 2	//敗北
			}
		case 3:	//エンカウント表示！！
			fmt.Println(m_sta.name,"があわられた！！")
		case 4:	//終了時メッセージ
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
		case 5://ゲーム開始待機のコンソール画面
			fmt.Printf("Press ENTER to start")
			ansi.CursorHide()
			for {
				x := getkey()
				if x == 13 {
					break
				}
			}


	}
	return 0
}

func prompt(p_sta *status, mode int) int{	//選択画面
	var action int
	switch mode {
	case 0:	//継続選択
		fmt.Println("")
		fmt.Println("やめる　")
		fmt.Printf("つづける")
		action = chose(2)
		return action

	case 1:	//行動選択
		fmt.Println("")
		fmt.Println("にげる　")
		fmt.Printf("こうげき")
		action = chose(2)
		return action
	case 2:	//player選択
		fmt.Println("")
		fmt.Printf("playerの選択>")
		fmt.Scan(&action)
		return action
	}
	return 0
}

func actionP(p_sta *status, m_sta *status, action int) {
	//乱数発生
	rand.Seed(time.Now().UnixNano())

	switch action {
		case 0:	//戦闘離脱
			fmt.Printf("%sはにげだした。。\n",p_sta.name)

		case 1:	//攻撃
			fmt.Println("プレイヤーのこうげき")
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
			} else {
				fmt.Printf("%sは%dのダメージをあたえた！\n", p_sta,name, p_sta.dmg)
			}
		default:
			fmt.Println("こんらんしている")
	}
}

func actionM(p_sta *status, m_sta *status) {
	//乱数発生
	rand.Seed(time.Now().UnixNano())

	fmt.Println("モンスターのこうげき")
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

}

func player_UI(p_sta *[]status, line int) {
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
		fmt.Printf("| %d.%s%s|HP:%-3d|ATK:%-3d|\n", i + 1, (*p_sta)[i].name, s, (*p_sta)[i].hp, (*p_sta)[i].atk)
	}
	fmt.Println("|　ぼうけんのしょをつくる　　　|")
	fmt.Println("--------------------------------")
}
