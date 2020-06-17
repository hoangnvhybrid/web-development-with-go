package main

import (
	"fmt"
)

type First func(z int) int

func squareSum(x int) First {
	return func(y int) int {
		return y * x
	}
}

type Operator func(x float64) float64

func Map(op Operator, a []float64) []float64 {
	res := make([]float64, len(a))
	for i, x := range a {
		res[i] = op(x)
	}
	return res
}
func roro(x float64) float64 {
	return 100 * x
}
func roro1() Operator {
	//return func(x float64) float64 {
	//	return 100 * x
	//}
	return roro
}

func New() (Count func()) {
	n := 0
	return func() {
		n++
		fmt.Println(n)
	}
}

func main() {
	var bien1 Operator
	{
		bien1 = roro1()
	}
	fmt.Println(bien1(10))

	f1, f2 := New(), New()
	f1() // 1
	f2() // 1 (different n)
	f1() // 2
	f2() // 2

	/*people := []string{"lan", "thuy anh", "thu", "nhung"}
	fmt.Println("truoc", people)
	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})
	fmt.Println("sau:", people)*/

	/*op := math.Abs
	a := []float64{1, -2}
	b := Map(op, a)
	fmt.Println(b)

	c := Map(func(x float64) float64 { return 10 * x }, b)
	fmt.Println(c)

	d := Map(roro, b)
	fmt.Println(d)*/

	/*	var dautien First
		{
			dautien = squareSum(10)
		}
		var dauHai First
		dauHai = squareSum(2)
		fmt.Println(dauHai(3))

		fmt.Println(dautien(10))
		a := squareSum(10)
		fmt.Println(a(3))*/
}
