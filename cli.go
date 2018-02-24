package main

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct {
	bc *BlockChain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
	fmt.Println("  addblock -data DATA - Add new block to the blockchain")
}

func (cli *CLI) validateArgs() bool {
	if len(os.Args) < 2 {
		cli.printUsage()
		return false
	}
	return true
}

func (cli *CLI) Run() {
	if cli.validateArgs() == false {
		cli.printUsage()
		return
	}
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")
	var err error
	switch os.Args[1] {
	case "addblock":
		err = addBlockCmd.Parse(os.Args[2:])
	case "printchain":
		err = printChainCmd.Parse(os.Args[2:])
	default:
		cli.printUsage()
		return
	}
	if err != nil {
		panic(err)
	}
	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			return
		}
		cli.bc.AddBlock(*addBlockData)
		fmt.Println("Success!")
	}
	if printChainCmd.Parsed() {
		cli.bc.Browse()
	}
}
