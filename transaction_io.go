package main

import (
	"bytes"
	"encoding/base64"
)

type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

type TxInput struct {
	Txid      []byte //引用TxOut的交易id
	Vout      int    // 引用TxOut在交易中的输出序号
	Signature []byte
	PubKey    []byte
}

func (in *TxInput) UseKey(PubKeyHash []byte) bool {
	return bytes.Compare(HashPubKey(in.PubKey), PubKeyHash) == 0
}

// IsLockedWithKey checks if the output can be used by the owner of the pubkey
func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

func (out *TxOutput) Lock(address string) {
	pubKeyHash, err := base64.StdEncoding.DecodeString(address)
	if err != nil {
		panic(err)
	}
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

func NewTxOutput(subsidy int, address string) *TxOutput {
	out := &TxOutput{subsidy, nil}
	out.Lock(address)
	return out
}
