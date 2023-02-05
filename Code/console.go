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
			fmt.Printf("| %4s: %4d     |\n", p_sta.name, p_sta.hp)
			fmt.Printf("| %4s: %4d     |\n", m_sta.name, m_sta.hp)
			fmt.Println("---------------------")

		case 2:	//戦闘終了判定
			if m_sta.hp <= 0 {
				fmt.Println("モンスターをたおした！")
				return 1	//勝利
			} else if p_sta.hp <= 0 {
				fmt.Println("プレイヤーはたおれた。。")
				return 2	//敗北
			}
		case 3://エンカウント表示！！
			fmt.Println(m_sta.name,"があわられた！！")
		case 4://終了時メッセージ
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


	}
	return 0
}

func prompt(p_sta *status, mode int) int{
	switch mode {
		case 0:	//継続選択
			fmt.Println("0:やめる")
			fmt.Println("1:つづける")
			fmt.Printf("行動の選択>")
			fmt.Scan(&p_sta.action)
			fmt.Printf("\n")
			return p_sta.action

		case 1:	//行動選択
			fmt.Println("0:にげる")
			fmt.Println("1:こうげき")
			fmt.Printf("行動の選択>")
			fmt.Scan(&p_sta.action)
			fmt.Printf("\n")	//見やすくするための改行
	}
	return 0
}

func actionP(p_sta *status) {
	//乱数発生
	rand.Seed(time.Now().UnixNano())

	switch p_sta.action {
		case 0:	//戦闘離脱
			fmt.Printf("%sはにげだした。。\n",p_sta.name)

		case 1:	//攻撃
			fmt.Println("プレイヤーのこうげき")
			//ダメージ計算
			p_sta.dmg = rand.Intn(p_sta.atk) + p_sta.atk_min
			if rand.Int() % 20 == 0 {	//5%くらい
				fmt.Println("！！！！かいしんのいちげき！！！！")
				p_sta.dmg *= 2	//実はかいしんのいちげきって、防御力無視で二倍なんですよね。
			}
			fmt.Printf("プレイヤーは%dのダメージをあたえた！\n", p_sta.dmg)
		default:
			fmt.Println("こんらんしている")
	}
}

func actionM(m_sta *status) {
	//乱数発生
	rand.Seed(time.Now().UnixNano())
	var x float32	//仮でおいてるやつ。
	var y float32 = 1.5
	switch m_sta.action {
		case 1:	//攻撃
			fmt.Println("モンスターのこうげき")
			//ダメージの計算
			m_sta.dmg = rand.Intn(m_sta.atk) + m_sta.atk_min
			x = float32(m_sta.dmg)
			if rand.Int() % 25 == 0 {	//4%くらい
				fmt.Println("！！！！つうこんのいちげき！！！！")
				m_sta.dmg = int(x * y)	//実はかいしんのいちげきって、防御力無視で二倍なんですよね。
			}
			fmt.Printf("モンスターは%dのダメージをあたえた。\n", m_sta.dmg)
		default:
			fmt.Println("モンスターはようすをみている")
	}
}
