package main
import "fmt"

type player struct {
	player_hp int
	player_dmg int
	player_atk int
	player_atk_c int
}

type monster struct {
	monster_hp int
	monster_dmg int
	monster_atk int
	monster_atk_c int
}

func main(){

	var(
		action int
	)

	for true {

		fmt.Println("0:にげる")
		fmt.Printf("行動の選択>")
		fmt.Scan(&action)

		//実際の行動
		//プレイヤーの行動
		switch action {
		case 0:
			fmt.Println("にげだした。。")

		default:
			fmt.Println("こんんらんしている")
		}

		if action == 0 {
			break
		}

		//モンスターの行動
		fmt.Println("ようすをみている\n")

	}
}
