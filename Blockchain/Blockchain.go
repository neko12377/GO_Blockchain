package Blockchain

import (
	"Go_Blockchain/Block"
)

type Blockchain struct {
	Blocks []*Block.Block
}

func (blockchain *Blockchain) AddBlock(Data []byte) {
	PrevBlock := blockchain.Blocks[len(blockchain.Blocks)-1] // 取出當前鏈上最後一個Block(前提是ㄧ定存在 Genesis Block 才不會有Error)
	NewBlock := Block.CreateBlock(Data, PrevBlock.Hash)      // 創建 Block
	blockchain.Blocks = append(blockchain.Blocks, NewBlock)  // 把新創建的 Block 接上去
}

func CreateBlockchain() *Blockchain {
	return &Blockchain{[]*Block.Block{Block.CreateGenesisBlock()}}
}
