package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Block struct {
	Data      string
	PrevHash  string
	Hash      string
	Timestamp int64
	Nonce     int
}

func CalculateHash(block Block) string {
	input := fmt.Sprintf("%s%s%d%d", block.Data, block.PrevHash, block.Nonce, block.Timestamp)
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func MineBlock(data string, prevHash string, difficulty int) Block {
	block := Block{
		Data:      data,
		PrevHash:  prevHash,
		Nonce:     0,
		Timestamp: time.Now().Unix(),
	}
	target := strings.Repeat("0", difficulty)
	for {
		block.Nonce = rand.Intn(10000000)
		block.Hash = CalculateHash(block)
		if strings.HasPrefix(block.Hash, target) {
			break
		}
	}
	return block
}

func MiningCoins() {
	rand.New(rand.NewSource(time.Now().Unix()))
	difficulty := 3
	//Mine a new block
	newBlock := MineBlock("TX1: Alice->Bob", "0000000000", difficulty)
	fmt.Printf("Mined Block:\nData: %s\nHash: %s\nNonce: %d\n",
		newBlock.Data, newBlock.Hash, newBlock.Nonce)
}
