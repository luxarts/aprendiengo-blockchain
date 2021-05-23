package blockchain

import (
	"aprendiengo-blockchain/internal/utils"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const Difficult = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	// Crea un numero
	target := big.NewInt(1)

	// Desplaza los bits de ese numero hacia la izquierda Difficult veces
	target.Lsh(target, uint(256-Difficult))

	// Devuelve la prueba
	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) NewNonce(nonce int) []byte {
	// Genera la información concatenando el hash anterior con la información del bloque
	data := append(pow.Block.PrevHash[:], pow.Block.Data...)

	// Concatena el número
	data = append(data, utils.IntToHex(int64(nonce))...)

	// Concatena la dificultad
	data = append(data, utils.IntToHex(int64(Difficult))...)

	return data
}

func (pow *ProofOfWork) Run() (int, [32]byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	// Itera todos los números positivos
	for nonce < math.MaxInt64 {
		// Prueba con el nonce
		data := pow.NewNonce(nonce)

		// Genera un hash
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)

		// Guarda el hash como un número
		intHash.SetBytes(hash[:])

		// Si el hash generado es menor al target
		if intHash.Cmp(pow.Target) == -1 {
			break
		}

		// Incrementa el número de prueba
		nonce++
	}
	fmt.Printf("\nNonce encontrado: %d\n", nonce)

	return nonce, hash
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	// Crea un nuevo nonce con el nonce encontrado
	data := pow.NewNonce(pow.Block.Nonce)

	// Genera un hash
	hash := sha256.Sum256(data)

	// Guarda el hash como un número
	intHash.SetBytes(hash[:])

	// Si el hash generado es menor al target el hash es valido
	return intHash.Cmp(pow.Target) == -1
}
