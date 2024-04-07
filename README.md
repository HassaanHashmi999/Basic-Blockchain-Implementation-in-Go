# Blockchain Implementation in Go

This is a simple implementation of a blockchain in Go programming language. The blockchain consists of blocks containing transactions and utilizes proof-of-work for mining new blocks. Below is an overview of the main components and functionalities of this implementation.

## Components

### Block
- Structure representing a block in the blockchain.
- Contains transaction data, nonce, hash, previous hash, and timestamp.

### Transaction
- Structure representing a transaction between two parties.
- Contains sender, recipient, amount, and a transaction pool.

### Blockchain
- Structure representing the entire blockchain.
- Consists of blocks and a transaction pool.
- Supports adding transactions, mining new blocks, listing blocks, displaying block data, and verifying the chain.

## Functions

- **mine_block**: Mines a new block with given transaction data and previous hash.
- **AddTransaction**: Adds a new transaction to the transaction pool.
- **NewBlock**: Creates a new block with provided transaction data and adds it to the blockchain.
- **ListBlocks**: Lists all the blocks in the blockchain along with their details.
- **VerifyChain**: Verifies the integrity of the blockchain by checking the consistency of hashes between blocks.
- **CalculateHash**: Calculates the SHA-256 hash of a given string.
- **DisplayBlockData**: Displays the data of a specific block in the blockchain.
- **print**: Prints the blockchain data in JSON format.

## Usage

- The `main` function initializes a new blockchain and performs various operations like adding transactions, mining blocks, and displaying blockchain data.
- Transactions are added using `AddTransaction`, and new blocks are mined using `NewBlock`.
- Blockchain integrity can be verified using `VerifyChain`.
- Detailed block data can be displayed using `DisplayBlockData`.
- Finally, the blockchain data is printed in JSON format using `print`.

## Instructions

1. Clone or download the repository.
2. Run the Go program by executing `go run main.go`.
3. Follow the instructions in the code comments to understand and modify the implementation as needed.

## Note

This implementation serves as a basic demonstration of blockchain concepts and may not be suitable for production use. Further enhancements and optimizations may be required for real-world applications.
