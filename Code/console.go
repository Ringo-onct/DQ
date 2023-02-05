package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

/*
p_sta *status	playerのステータス
m_sta *status	monsterのステータス
mode	表示する内容をmodeで選択
*/
func console(p *player, m *monster, mode int) {
	if mode == 0 { //コンソール画面クリア
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
	} else if mode == 1 { //体力表示
		fmt.Println("---------------------")
		fmt.Printf("| PLAYER : %4d     |\n", p.hp)
		fmt.Printf("| MONSTER: %4d     |\n", m.hp)
		fmt.Println("---------------------")
	}
}
