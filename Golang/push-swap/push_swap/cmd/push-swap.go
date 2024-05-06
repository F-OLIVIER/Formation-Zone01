package main

import (
	internal "push-swap/push_swap/internal"
)

func main() {
	pileA, pileB := internal.VerifyAndConvert()
	internal.RetourDePileAetPileB(pileA, pileB)
}
