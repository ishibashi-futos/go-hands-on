package main

import "fmt"

func main() {
	var num int = 10
	switch num {
	case 10:
		fmt.Println("10")
		fallthrough
	case 11:
		x, y := calc(10, 20)
		fmt.Println("total", x)
		fmt.Println("Average", y)
		fallthrough
	case 12:
		z := func() string { return "Hello, World" }
		fmt.Println(z())
		fallthrough
	case 13:
		fmt.Println(calc2(10, 20))
	default:
		fmt.Println("oh...")
	}

	fmt.Println(arrays(0, 3))
	user1 := user{
		Name: "ishibashi.futoshi",
		Age:  26,
	}
	fmt.Println(user1)
	pointer()
	var a = 10
	var b = 10
	// 値渡しで計算するfunction
	add1 := func(x int) { x = x + 1 }
	add1(a)
	// 参照渡しで破壊的関数？のような作りに・・・
	add2 := func(x *int) { *x = *x + 1 }
	add2(&b)
	fmt.Println(a, b)
	deRefarence()
}

func calc(x int, y int) (int, int) {
	return x + y, (x + y) * 2
}

func calc2(x int, y int) (sum int, average int) {
	sum = x + y
	average = (x + y) * 2
	return
}

// slise
func arrays(start int, end int) []string {
	var arr = []string{"a", "b", "c", "d", "e", "f", "g"}
	return arr[start:end]
}

type user struct {
	Name string
	Age  int
}

func pointer() {
	var p *int
	var n = 10
	p = &n
	fmt.Println(*p)
}

func deRefarence() {
	user1 := user{
		Name: "takeshi",
	}
	user2 := user{
		Name: "takashi",
	}
	rename1 := func(u user) { u.Name = "yamada" }
	rename2 := func(u *user) { u.Name = "yamada" }
	rename1(user1)
	rename2(&user2)
	fmt.Println(user1, user2)
}
