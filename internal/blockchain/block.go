package blockchain

type Block struct {
	PrevHash [32]byte
	Data     []byte
	Hash     [32]byte
	Nonce    int
}

func NewBlock(data string, prevHash [32]byte) *Block {
	// Crea un nuevo bloque
	block := &Block{
		PrevHash: prevHash,
		Data:     []byte(data),
	}

	// Crea una prueba de trabajo para ese bloque
	pow := NewProofOfWork(block)

	// Busca el hash válido
	nonce, hash := pow.Run()

	// Asigna el nonce y el hash al bloque
	block.Nonce = nonce
	block.Hash = hash

	return block
}
