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





type Blockchain struct {
	transactionPool []string 	
	chain []*Block
}


func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	genesisBlock := bc.CreateBlock(0, "init hash")
	bc.chain = append(bc.chain, genesisBlock)
	return bc
}	
func(bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	block := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.transactionPool = []string{}
	bc.chain = append(bc.chain, block)
	return block

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