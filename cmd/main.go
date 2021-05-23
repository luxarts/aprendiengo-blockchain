package main

import (
	"aprendiengo-blockchain/internal/blockchain"
	"fmt"
	"strconv"
)

func main() {
	// Crea una cadena
	chain := blockchain.NewChain()

	// Agrega un bloque
	chain.AddBlock("bloque 1")

	// Agrega otro bloque
	chain.AddBlock("bloque 2")

	for i, b := range chain.GetBlocks() {
		fmt.Printf("==========Bloque %d==========\n", i)
		fmt.Printf("Hash anterior: %x\n", b.PrevHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)

		pow := blockchain.NewProofOfWork(b)
		valid := pow.Validate()

		fmt.Printf("Pow: %s\n", strconv.FormatBool(valid))
		fmt.Printf("=============================\n")
	}
}
