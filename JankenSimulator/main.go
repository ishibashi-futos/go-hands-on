package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Hands int

const (
	Rock Hands = iota
	Papaer
	Scissors
)

func (h Hands) String() string {
	switch h {
	case Rock:
		return "Rock"
	case Papaer:
		return "Papaer"
	case Scissors:
		return "Scissors"
	default:
		return "Unknown"
	}
}

type Result struct {
	yourWin int
	meWin   int
	draw    int
}

func (r *Result) addYwin() {
	r.yourWin++
}

func (r *Result) addMwin() {
	r.meWin++
}

func (r *Result) addDraw() {
	r.draw++
}

func main() {
	COUNT := 1000000
	rand.Seed(time.Now().UnixNano())
	result := Result{
		yourWin: 0,
		meWin:   0,
		draw:    0,
	}
	for i := 0; i < COUNT; i++ {
		yourHands, meHands := poi()
		switch {
		case yourHands == meHands:
			result.addDraw()
		case yourHands.String() == "Rock" && meHands.String() == "Scissors",
			yourHands.String() == "Scissors" && meHands.String() == "Papaer",
			yourHands.String() == "Papaer" && meHands.String() == "Rock":
			result.addYwin()
		default:
			result.addMwin()
		}
	}

	fmt.Printf("draw:%d\n", result.draw)
	fmt.Printf("yourWin:%d\n", result.yourWin)
	fmt.Printf("meWin:%d\n", result.meWin)
}

func poi() (your Hands, me Hands) {
	your = yourHands()
	me = meHands()
	return
}

func yourHands() Hands {
	return Hands(rand.Intn(3))
}

func meHands() Hands {
	return Hands(rand.Intn(3))
}
