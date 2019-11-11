package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

// TXOutput represents a transaction output
// TXOutput 代表一个交易的输出
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

// Lock signs the output
// Lock 签名输出
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// IsLockedWithKey checks if the output can be used by the owner of the pubkey
// IsLockedWithKey 检查是否提供的公钥哈希被用于锁定输出
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// NewTXOutput create a new TXOutput
// NewTXOutput 创建一个新的交易输出
func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

// TXOutputs collects TXOutput
// TXOutputs 是TXOutput的集合
type TXOutputs struct {
	Outputs []TXOutput
}

// Serialize serializes TXOutputs
// Serialize 序列化交易输出的集合
func (outs TXOutputs) Serialize() []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(outs)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// DeserializeOutputs deserializes TXOutputs
// DeserializeOutputs 反序列化交易输出的集合
func DeserializeOutputs(data []byte) TXOutputs {
	var outputs TXOutputs

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}

	return outputs
}
