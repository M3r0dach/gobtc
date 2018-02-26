package main

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("   printchain - Print all the blocks of the blockchain")
	fmt.Println("   createblockchain -address ADDRESS -  Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("   getbalance -address ADDRESS - Get the balance of ADDRESS")
	fmt.Println("   send -from FROM -to TO -amount AMOUNT - send AOUNT coins from FROM to TO")
}

func (cli *CLI) validateArgs() bool {
	if len(os.Args) < 2 {
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
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)

	createBlockChainAddress := createBlockChainCmd.String("address", "", "The address to send genesis block reward to")
	getBalanceAddress := getBalanceCmd.String("address", "", "the address to get balance for")
	sendFrom := sendCmd.String("from", "", "Source address")
	sendTo := sendCmd.String("to", "", "Destination address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send must be greater than 0")
	var err error
	switch os.Args[1] {
	case "createblockchain":
		err = createBlockChainCmd.Parse(os.Args[2:])
	case "printchain":
		err = printChainCmd.Parse(os.Args[2:])
	case "getbalance":
		err = getBalanceCmd.Parse(os.Args[2:])
	case "send":
		err = sendCmd.Parse(os.Args[2:])
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
		bc := CreateBlockChain(*createBlockChainAddress)
		defer bc.db.Close()
		fmt.Println("Done!")
	}
	if printChainCmd.Parsed() {
		bc := NewBlockChain()
		defer bc.db.Close()
		bc.Browse()
	}
	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		bc := NewBlockChain()
		defer bc.db.Close()
		UTXOS := bc.FindUTXO(*getBalanceAddress)
		balance := 0
		for _, out := range UTXOS {
			balance += out.Value
		}
		fmt.Printf("Balance of '%s': %d\n", *getBalanceAddress, balance)
	}
	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}
		bc := NewBlockChain()
		defer bc.db.Close()
		tx := NewUTXOTransaction(*sendFrom, *sendTo, *sendAmount, bc)
		bc.MineBlock([]*Transaction{tx})
		fmt.Println("Success")
	}
}
