package main
import (
	"fmt"
	"math/rand"
	"time"
)

func action(p *player, m *monster, mode) {
	//乱数発生
	rand.Seed(time.Now().UnixNano())
	
	switch p_sta.action {
	case 0:
		fmt.Println("にげだした。。")

	case 1:
		fmt.Println("プレイヤーのこうげき")
		//ダメージ計算
		p.dmg = rand.Intn(p.atk) + p.atk_min
		fmt.Printf("プレイヤーは%dのダメージをあたえた！\n", p.dmg)

	default:
		fmt.Println("こんらんしている")
	}
}
