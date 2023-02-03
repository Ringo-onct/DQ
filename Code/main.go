package main

import "fmt"

type Status struct {
	name   string
	hp     int
	mp     int
	atk    int
	def    int
	luk    int
	action int
}

func main() {

	var (
	//action int	廃止
	)

	p_sta := new(Status)

	for true {

		fmt.Println("0:にげる")
		fmt.Printf("行動の選択>")
		fmt.Scan(&p_sta.action)

		//実際の行動
		//プレイヤーの行動
		switch p_sta.action {
		case 0:
			fmt.Println("にげだした。。")

		default:
			fmt.Println("こんんらんしている")
		}

		if p_sta.action == 0 {
			break
		}

		//モンスターの行動
		fmt.Println("ようすをみている")

	}
}
