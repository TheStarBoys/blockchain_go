package main

import "bytes"

// TXInput represents a transaction input
// TXInput 代表一个交易的输入
type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte
}

// UsesKey checks whether the address initiated the transaction
// UsesKey 检查输入使用了指定密钥来解锁一个输出
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
