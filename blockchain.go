package main

import (
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

const (
	dbFile               = "data.bolt"
	blocksBucket         = "blocksBucket"
	genesisCoinnbaseData = "Power is the only true thing"
)

type BlockChain struct {
	db  *bolt.DB
	tip []byte
}

func dbExist() bool {
	_, err := os.Stat(dbFile)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

/*
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
*/
func CreateBlockChain(address string) *BlockChain {
	if dbExist() {
		fmt.Println("BlockChain already exists.")
		os.Exit(1)
	}

	var tip []byte
	cbtx := NewCoinbaseTX(address, genesisCoinnbaseData)
	genesis := NewGenesisBlock(cbtx)
	tip = genesis.Hash
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			return err
		}
		err = b.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			return err
		}
		err = b.Put([]byte("l"), genesis.Hash)
		return err
	})
	if err != nil {
		panic(err)
	}
	return &BlockChain{db, tip}
}

func NewBlockChain() *BlockChain {
	if dbExist() == false {
		fmt.Println("No existing blockchain found. Create one first.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open("data.bolt", 0600, nil)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		tip = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		panic(err)
	}
	return &BlockChain{db, tip}
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.tip, bc.db}
}

func (bc *BlockChain) Browse() {
	i := bc.Iterator()
	for len(i.currentHash) > 0 {
		block := i.Next()
		fmt.Println(block)
	}
}
