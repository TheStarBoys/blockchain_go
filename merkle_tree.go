package main

import (
	"crypto/sha256"
)

// MerkleTree represent a Merkle tree
// MerkleTree 代表一个默克尔树
type MerkleTree struct {
	RootNode *MerkleNode
}

// MerkleNode represent a Merkle tree node
// MerkleNode 代表一个默克尔树节点
type MerkleNode struct {
	Left  *MerkleNode
	Right *MerkleNode
	Data  []byte
}

// NewMerkleTree creates a new Merkle tree from a sequence of data
// NewMerkleTree 从一个data序列, 创建并返回一个默克尔树
func NewMerkleTree(data [][]byte) *MerkleTree {
	// 这里的MerkleTree算法底层不是采用的二叉树存储节点，而是一个数组
	var nodes []MerkleNode
	// 如果交易数据不为偶数, 拷贝最后一个交易, 凑成偶数个交易
	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}
	// 为nodes填充节点数据
	for _, datum := range data {
		node := NewMerkleNode(nil, nil, datum)
		nodes = append(nodes, *node)
	}
	// 计算MerkleRoot
	for i := 0; i < len(data)/2; i++ { // 外层控制MerkleTree层数
		var newLevel []MerkleNode

		for j := 0; j < len(nodes); j += 2 { // 内层计算出上一层nodes
			node := NewMerkleNode(&nodes[j], &nodes[j+1], nil)
			newLevel = append(newLevel, *node)
		}

		nodes = newLevel
	}

	mTree := MerkleTree{&nodes[0]}

	return &mTree
}

// NewMerkleNode creates a new Merkle tree node
// NewMerkleNode 创建一个新的默克尔树节点
func NewMerkleNode(left, right *MerkleNode, data []byte) *MerkleNode {
	mNode := MerkleNode{}
	// 如果没有左右节点, 直接存储交易数据的hash
	if left == nil && right == nil {
		hash := sha256.Sum256(data)
		mNode.Data = hash[:]
	} else { // 否则, 对左右节点的hash进行拼接后, 再求当前节点的hash
		prevHashes := append(left.Data, right.Data...)
		hash := sha256.Sum256(prevHashes)
		mNode.Data = hash[:]
	}

	mNode.Left = left
	mNode.Right = right

	return &mNode
}
