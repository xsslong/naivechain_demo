package naivechain_demo

import (
	"sync"
	"time"
)

// 用于操作链
type Blockchain struct {
	blocks Blocks
	mu     sync.RWMutex
}

// 创建区块链
func newBlockChain() *Blockchain {
	return &Blockchain{
		blocks: Blocks{genesisBlock},
		mu:     sync.RWMutex{},
	}
}

// 计算链的长度
func (blockchain *Blockchain) len() int64 {
	return int64(len(blockchain.blocks))
	//return blockchain.blocks.Len()
}

// 获得创世块
func (blockchain *Blockchain) getGenesisBlock() *Block {
	return blockchain.getBlock(0)
}

// 获得最新块

func (blockchain *Blockchain) getLatestBlock() *Block {
	return blockchain.getBlock(blockchain.len() - 1)
}

// 根据索引获得块
func (blockchain *Blockchain) getBlock(index int64) *Block {
	blockchain.mu.RLock()
	defer blockchain.mu.RUnlock()
	return blockchain.blocks[index]
}

// 验证创世块是否有效
func (blockchain *Blockchain) isGenesisBlockValid() bool {
	gBlock := blockchain.getGenesisBlock()
	return gBlock.Hash == genesisBlock.Hash &&
		gBlock.isValidHash()
}

// 创建块
func (blockchain *Blockchain) generateBlock(data string) *Block {
	block := &Block{
		Index:        blockchain.getLatestBlock().Index + 1,
		Timestamp:    time.Now().Unix(),
		Data:         data,
		Previoushash: blockchain.getLatestBlock().Hash,
	}
	block.Hash = block.hash()
	return block
}

// 验证区块链是否有效, 只需要把链上相邻的块信息递归验证即可
func (blockchain Blockchain) isBlockchainValid() bool {
	blockchain.mu.RLock()
	defer blockchain.mu.RUnlock()
	if blockchain.len() == 0 || !blockchain.isGenesisBlockValid() {
		return false
	}
	preBlock := blockchain.getGenesisBlock()
	var i int64
	for i = 1; i < blockchain.len(); i++ {
		nowBlock := blockchain.getBlock(i)
		if ok := isValidBlock(nowBlock, preBlock); !ok {
			return false
		}
		preBlock = nowBlock
	}
	return true
}

// 添加块
func (blockchain Blockchain) addBlock(block *Block) {
	blockchain.mu.Lock()
	defer blockchain.mu.Unlock()
	blockchain.blocks = append(blockchain.blocks, block)
}
