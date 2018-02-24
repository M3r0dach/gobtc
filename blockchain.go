package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

const (
	blocksBucket = "blocksBucket"
)

type Block struct {
	TimeStamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
		0,
	}
	pow := NewProofOfWork(block)
	block.Nonce, block.Hash = pow.Run()
	return block
}
func NewGenesisBlock() *Block {
	block := NewBlock("This is the study of blockchain", []byte{})
	return block
}

func (b *Block) String() string {
	var buff bytes.Buffer
	fmt.Fprintf(&buff, "Timestamp: %d\n", b.TimeStamp)
	fmt.Fprintf(&buff, "Data: %s\n", b.Data)
	fmt.Fprintf(&buff, "Prev: %x\n", b.PrevBlockHash)
	fmt.Fprintf(&buff, "Hash: %x\n", b.Hash)
	fmt.Fprintf(&buff, "Nonce: %d\n", b.Nonce)
	pow := NewProofOfWork(b)
	fmt.Fprintf(&buff, "Pow: %v\n", pow.Validate())
	return buff.String()
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	encoder.Encode(b)
	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	decoder.Decode(&block)
	return &block
}

type BlockChain struct {
	db  *bolt.DB
	tip []byte
}

func (bc *BlockChain) AddBlock(data string) {
	var lastHash []byte
	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("1"))
		return nil
	})
	if err != nil {
		panic(err)
	}
	newBlock := NewBlock(data, lastHash)
	bc.tip = newBlock.Hash
	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put([]byte(newBlock.Hash), newBlock.Serialize())
		if err != nil {
			return err
		}
		err = b.Put([]byte("1"), newBlock.Hash)
		return err
	})
	if err != nil {
		panic(err)
	}
}

func NewBlockChain() *BlockChain {
	var tip []byte
	db, err := bolt.Open("data.bolt", 0600, nil)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			if err != nil {
				return err
			}
			tip = genesis.Hash
			err = b.Put([]byte("1"), tip)
			return err
		}
		tip = b.Get([]byte("1"))
		return nil
	})
	if err != nil {
		panic(err)
	}
	return &BlockChain{db, tip}
}

type BlockChainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.tip, bc.db}
}

func (i *BlockChainIterator) Next() *Block {
	var block *Block
	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		if encodedBlock == nil {
			return errors.New("cannot get block")
		}
		block = DeserializeBlock(b.Get(i.currentHash))
		return nil
	})
	if err != nil {
		panic(err)
	}
	i.currentHash = block.PrevBlockHash
	return block
}

func (bc *BlockChain) Browse() {
	i := bc.Iterator()
	for len(i.currentHash) > 0 {
		block := i.Next()
		fmt.Println(block)
	}
}
