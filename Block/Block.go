package Block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// 設置 block 的 hash
func (b *Block) setHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10)) // 把 Timestamp 換算成 10 進位 (藉由CreateBlock賦予Timestamp)
	payload := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(payload) // 產生 Hash
	b.Hash = hash[:]               // 設置 Hash
}

func CreateBlock(Data []byte, PrevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(), // Unix timestamp
		Data:          Data,
		PrevBlockHash: PrevBlockHash,
		Hash:          []byte{},
	}
	block.setHash()
	return block
}

func CreateGenesisBlock() *Block {
	return CreateBlock([]byte("Genesis Block"), []byte{})
}
