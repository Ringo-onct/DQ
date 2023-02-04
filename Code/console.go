package main
import (
	"fmt"
)
/*
p_sta *status	playerのステータス
m_sta *status	monsterのステータス
mode	表示する内容をmodeで選択
*/
func console(p_sta *status, m_sta *status, mode) {
	if mode == 1 {
		fmt.Println("---------------------")
		fmt.Printf("| PLAYER : %4d     |\n", p.hp)
		fmt.Printf("| MONSTER: %4d     |\n", m.hp)
		fmt.Println("---------------------")
	}
}
