package main

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
	fmt.Println("  createblockchain -address AddRESS -  Create a blockchain and send genesis block reward to ADDRESS")
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
	createBlockChainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockChainAddress := createBlockChainCmd.String("address", "", "The address to send genesis block reward to")
	var err error
	switch os.Args[1] {
	case "createblockchain":
		err = createBlockChainCmd.Parse(os.Args[2:])
	case "printchain":
		err = printChainCmd.Parse(os.Args[2:])
	default:
		cli.printUsage()
		return
	}
	if err != nil {
		panic(err)
	}
	if createBlockChainCmd.Parsed() {
		if *createBlockChainAddress == "" {
			createBlockChainCmd.Usage()
			return
		}
		//cli.bc.AddBlock(*addBlockData)
		fmt.Println("Success!")
	}
	if printChainCmd.Parsed() {
		bc := NewBlockChain()
		defer bc.db.Close()
		bc.Browse()
	}
}
