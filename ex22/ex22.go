package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("9223372036854775807012311", 10)
	b.SetString("234249328237419234123941237412", 10)

	fmt.Println(a, "+", b, "=", sum(a, b))
	fmt.Println(a, "-", b, "=", sub(a, b))
	fmt.Println(a, "*", b, "=", mul(a, b))
	fmt.Println(b, "/", a, "=", div(a, b))
}

func sum(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int)

	res.Add(a, b)

	return res
}

func sub(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int)

	res.Sub(a, b)

	return res
}

func div(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int)

	res.Div(a, b)

	return res
}

func mul(a *big.Int, b *big.Int) *big.Int {
	res := new(big.Int)

	res.Mul(a, b)

	return res
}
