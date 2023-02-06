package main
import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

func Math(mode int) string{
	rand.Seed(time.Now().UnixNano())	//randのシード

	switch mode 1:	//monsterシードランダム
		x := rand.Int() % 2 + 1
		seed := "a" + strconv.Itoa(x)

	return seed
}
