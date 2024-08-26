package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(catalan(7))  // 429
	fmt.Println(catalan(10)) // 16796
}

func catalan(n int) *big.Int {
	
	twoNFact := big.NewInt(1)
	nFact := big.NewInt(1)
	nPlusOneFact := big.NewInt(1)

	for i := 2; i <= 2*n; i++ {
		twoNFact.Mul(twoNFact, big.NewInt(int64(i)))
	}

	for i := 2; i <= n; i++ {
		nFact.Mul(nFact, big.NewInt(int64(i)))
	}

	for i := 2; i <= n+1; i++ {
		nPlusOneFact.Mul(nPlusOneFact, big.NewInt(int64(i)))
	}

	result := big.NewInt(1)
	result.Mul(twoNFact, result)
	result.Div(result, nPlusOneFact)
	result.Div(result, nFact)

	return result
}
