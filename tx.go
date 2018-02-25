package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

const subsidy = 50

type Transaction struct {
	ID   []byte // 交易id
	Vin  []TxInput
	Vout []TxOutput
}

type TxOutput struct {
	Value        int
	ScriptPubKey string
}

type TxInput struct {
	Txid      []byte //交易id
	Vout      int    // 输出序号
	ScriptSig string
}

func (tx *Transaction) Serialize() []byte {
	var encoded bytes.Buffer
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		panic(err)
	}
	return encoded.Bytes()
}

func (tx *Transaction) Hash() []byte {
	var hash [32]byte
	txCopy := *tx
	txCopy.ID = []byte{}
	hash = sha256.Sum256(txCopy.Serialize())
	return hash[:]
}

func (tx *Transaction) String() string {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "Transaction %x:\n", tx.ID)
	for i, input := range tx.Vin {
		fmt.Fprintf(&buff, "  Input %d:\n", i)
		fmt.Fprintf(&buff, "    TXID:    %x\n", input.Txid)
		fmt.Fprintf(&buff, "    Out:     %d\n", input.Vout)
		fmt.Fprintf(&buff, "    Script:  %x\n", input.ScriptSig)
	}
	for i, output := range tx.Vout {
		fmt.Fprintf(&buff, "  Output %d:\n", i)
		fmt.Fprintf(&buff, "    Value:  %d\n", output.Value)
		fmt.Fprintf(&buff, "    Script: %x\n", output.ScriptPubKey)
	}
	return buff.String()
}

func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{subsidy, to}
	tx := Transaction{[]byte{},
		[]TxInput{txin},
		[]TxOutput{txout},
	}
	tx.ID = tx.Hash()
	return &tx
}