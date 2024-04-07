package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const DifficultyLevel = 2

type Block struct {
	transaction  string
	nonce        int
	hash         string
	previousHash string
	timestamp    time.Time
}

type Transaction struct {
	From   string
	To     string
	Amount int
	tPool  []*Transaction
}

func NewTransaction(from string, to string, amount int) *Transaction {
	t := new(Transaction)
	t.From = from
	t.To = to
	t.Amount = amount

	return t
}

type BlockChain struct {
	blocks []*Block
	tPool  []*Transaction
}

func (b *BlockChain) AddTransaction(sender string, recipient string, value int) {
	t := NewTransaction(sender, recipient, value)
	b.tPool = append(b.tPool, t)
}

func mine_block(transaction string, previousHash string, difficulty int) *Block {
	bl := new(Block)
	bl.transaction = transaction
	bl.previousHash = previousHash
	bl.timestamp = time.Now()

	prefix := strings.Repeat("0", difficulty)
	for nonce := 0; ; nonce++ {
		data := strconv.Itoa(nonce) + bl.previousHash + bl.transaction
		hash := CalculateHash(data)
		if strings.HasPrefix(hash, prefix) {
			bl.nonce = nonce
			bl.hash = hash
			return bl
		}
	}
}

func (b *BlockChain) NewBlock(transaction string, previousHash string) *Block {
	minedBlock := mine_block(transaction, previousHash, DifficultyLevel)
	b.blocks = append(b.blocks, minedBlock)
	return minedBlock
}

func ListBlocks(ls *BlockChain) {
	counter := 1
	for _, i := range ls.blocks {
		fmt.Printf("%s Transaction %d %s\n", strings.Repeat("=", 35), counter, strings.Repeat("=", 35))
		fmt.Println("Transaction:   ", i.transaction)
		fmt.Println("Nonce:         ", i.nonce)
		fmt.Println("Previous Hash: ", i.previousHash)
		fmt.Println("Hash:          ", i.hash)
		fmt.Println("Timestamp:     ", i.timestamp)
		fmt.Println()
		counter++
	}

	for _, i := range ls.tPool {
		fmt.Println("Transaction number: ", i)
		fmt.Println("From: ", i.From)
		fmt.Println("To: ", i.To)
		fmt.Println("Amount: ", i.Amount)
	}
}

func VerifyChain(b *BlockChain) int {
	count := 0
	flag := 0

	for count = 0; count < (len(b.blocks) - 1); count++ {
		if b.blocks[count].hash != b.blocks[count+1].previousHash {
			flag = 1
			fmt.Println("Changes detected in Block", count+1)
		}
	}

	return flag
}

func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	Hash := fmt.Sprintf("%x", hash)
	return Hash
}

func DisplayBlockData(b *BlockChain, index int) {
	if index >= 0 && index < len(b.blocks) {
		block := b.blocks[index]
		fmt.Printf("Block %d:\n", index)
		fmt.Printf("  Hash: %s\n", block.hash)
		fmt.Printf("  Previous Hash: %s\n", block.previousHash)
		fmt.Printf("  Nonce: %d\n", block.nonce)
		fmt.Printf("  Timestamp: %s\n", block.timestamp)
		fmt.Printf("  Transactions:\n")
		for _, t := range b.tPool {
			fmt.Printf("    From: %s\n", t.From)
			fmt.Printf("    To: %s\n", t.To)
			fmt.Printf("    Amount: %d\n", t.Amount)
			fmt.Println()
		}
	}
}

func (b *BlockChain) print() {
	for i, block := range b.blocks {

		data := map[string]interface{}{
			"Block Number ":         i,
			"Time Stamp ":           block.timestamp.String(),
			"Nonce ":                block.nonce,
			"Previous Hash ":        block.previousHash,
			"Current Block's Hash ": block.hash,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(jsonData))

	}

}
func main() {
	b := new(BlockChain)

	b.NewBlock("JAke to Qasim", "0")
	b.AddTransaction("Jake", "Qasim", 2)
	b.AddTransaction("Jake", "Jake", 3)

	b.NewBlock("Alice to BoB ", b.blocks[0].hash)
	b.AddTransaction("Alice", "BoB", 4)
	b.AddTransaction("Alice", "Alice", 5)

	ListBlocks(b)
	b.print()

	DisplayBlockData(b, 0)

	if VerifyChain(b) == 0 {
		fmt.Println("No changes detected in the blockchain")
	}
}
