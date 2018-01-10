package naivechain_demo

import (
	"fmt"
	"crypto/sha256"
)

// 用于描述块结构信息
// 这里只描述最必要的块信息: index(索引)、timestamp(时间戳)、data(数据)、hash(哈希值)和previoushash(前置哈希值)
type Block struct {
	Index        int64  `json:"index"`
	Timestamp    int64  `json:"timestamp"`
	Data         string `json:"data"`
	Previoushash string `json:"previoushash"`
	Hash         string `json:"hash"`
}

type Blocks []*Block

// 创世块, 起源块
var genesisBlock = &Block{
	Index:        0,
	Timestamp:    1517457600, // 2018/2/1 12:00:00
	Data:         "这是劳资的创世块",
	Previoushash: "0",
	Hash:         "e43db0d2d1aec23682d596ccac44be5cd6e4653404738d73bc66a91dedd32132", // 创世块的hash值硬编码
}

// 获取块的数量长度
func (blocks Blocks) Len() int {
	return len(blocks)
}

// 判断i块长度是否小于j块长度
// 任何时候在链中都应该只有一组明确的块, 万一冲突了（例如：两个结点都生成了72号块时），会选择有最大数目的块的链(最长的链)
func (blocks Blocks) less(i, j int64) bool {
	return blocks[i].Index < blocks[j].Index
}

// 确认块的完整性, 相邻块信息递归验证
func isValidBlock(block, preBlock *Block) bool {
	return block.Index == preBlock.Index+1 &&
		block.Previoushash == preBlock.Hash &&
		block.isValidHash()
}

// 确认块的Hash是否有效
func (block *Block) isValidHash() bool {
	return block.Hash == block.hash()
}

// 根据块的其他信息(不包括hash)计算出自己的hash
func (block *Block) hash() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf(
		"%d%s%d%s",
		block.Index, block.Previoushash, block.Timestamp, block.Data,
	))))
}
