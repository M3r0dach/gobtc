package main

func main() {
	bc := NewBlockChain()
	defer bc.db.Close()
	cli := CLI{bc}
	cli.Run()
}
