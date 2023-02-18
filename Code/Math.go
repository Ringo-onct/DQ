package main
import (
	"time"
	"math/rand"
	"strconv"
)

func math(mode int) string {
	rand.Seed(time.Now().UnixNano())	//randのシード
	var (
		seed string
		x int
	)
	switch mode {
		case 1:	//monsterシードランダム
			x = (rand.Intn(2))
			seed = "a" + strconv.Itoa(x)
		default:
	}
	return seed
}
