package main

import (
	"fmt"

	tm "github.com/buger/goterm"
)

/*
p_sta *status	playerのステータス
m_sta *status	monsterのステータス
mode	表示する内容をmodeで選択
*/
func console(p *player, m *monster, mode int) {
	if mode == 0 { //コンソール画面クリア
		tm.Clear()
		tm.MoveCursor(1, 1)
		tm.Flush()
	} else if mode == 1 { //体力表示
		fmt.Println("---------------------")
		fmt.Printf("| PLAYER : %4d     |\n", p.hp)
		fmt.Printf("| MONSTER: %4d     |\n", m.hp)
		fmt.Println("---------------------")
	}
}
