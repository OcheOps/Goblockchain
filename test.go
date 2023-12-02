package main

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce         int
	previousHash  string
	timestamp     int64
	transactions  []string
}

func NewBlock(nonce int, previousHash string, transactions []string) *Block {
	b := new(Block)
	b.nonce = nonce
	b.timestamp = time.Now().UnixNano()
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func (b *Block) Print() {
	fmt.Printf("timestamp: %d\n", b.timestamp)
	fmt.Printf("nonce: %d\n", b.nonce)
	fmt.Printf("previousHash: %s\n", b.previousHash)
	fmt.Printf("transactions: %v\n", b.transactions)
}

func main() {
	// Example usage
	transactions := []string{"transaction1", "transaction2", "transaction3"}
	b := NewBlock(123, "init hash", transactions)
	b.Print()
}