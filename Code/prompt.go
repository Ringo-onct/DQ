package main
import (
	"fmt"
)

func prompt(p_sta *status, mode int) {
	if mode == 1 {
		fmt.Println("0:にげる")
		fmt.Println("1:こうげき")
		fmt.Printf("行動の選択>")
		fmt.Scan(&p_sta.action)
		fmt.Printf("\n")	//見やすくするための改行
	}
}
