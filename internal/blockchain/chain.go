package blockchain

type Chain struct {
	blocks []*Block
}

func NewChain() *Chain {
	// Genera un bloque
	originBlock := NewBlock("origin", [32]byte{})

	// Genera una cadena con el bloque origen
	return &Chain{
		[]*Block{
			originBlock,
		},
	}
}

func (c *Chain) AddBlock(data string) {
	// Obtiene el último bloque de la cadena
	prevBlock := c.blocks[len(c.blocks)-1]

	// Crea un nuevo bloque usando el hash anterior
	newBlock := NewBlock(data, prevBlock.Hash)

	// Agrega el bloque a la cadena
	c.blocks = append(c.blocks, newBlock)
}

func (c *Chain) GetBlocks() []*Block {
	return c.blocks
}
