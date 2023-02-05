package main
import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"os/exec"
	"runtime"
)

func console(p_sta *status, m_sta *status, mode int) int {
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
			fmt.Println("---------------------")
			fmt.Printf("| PLAYER : %4d     |\n", p_sta.hp)
			fmt.Printf("| MONSTER: %4d     |\n", m_sta.hp)
			fmt.Println("---------------------")

		case 2:	//戦闘終了判定
			if m_sta.hp <= 0 {
				fmt.Println("モンスターをたおした！")
				return 1	//勝利
			} else if p_sta.hp <= 0 {
				fmt.Println("プレイヤーはたおれた。。")
				return 2	//敗北
			}


	}
	return 0
}

func prompt(p_sta *status, mode int) {
	switch mode {
		case 1:	//行動選択
			fmt.Println("0:にげる")
			fmt.Println("1:こうげき")
			fmt.Printf("行動の選択>")
			fmt.Scan(&p_sta.action)
			fmt.Printf("\n")	//見やすくするための改行
	}

}

func actionP(p_sta *status) {
	//乱数発生
	rand.Seed(time.Now().UnixNano())

	switch p_sta.action {
		case 0:	//戦闘離脱
			fmt.Println("にげだした。。")

		case 1:	//攻撃
			fmt.Println("プレイヤーのこうげき")
			//ダメージ計算
			p_sta.dmg = rand.Intn(p_sta.atk) + p_sta.atk_min
			fmt.Printf("プレイヤーは%dのダメージをあたえた！\n", p_sta.dmg)
		default:
			fmt.Println("こんらんしている")
	}
}

func actionM(m_sta *status) {
	//乱数発生
	rand.Seed(time.Now().UnixNano())

	switch m_sta.action {
		case 1:	//攻撃
			fmt.Println("モンスターのこうげき")
			//ダメージの計算
			m_sta.dmg = rand.Intn(m_sta.atk) + m_sta.atk_min
			fmt.Printf("モンスターは%dのダメージをあたえた。\n", m_sta.dmg)
		default:
			fmt.Println("モンスターはようすをみている")
	}
}
