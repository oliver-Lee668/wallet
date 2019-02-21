package main

import (
	"wallet/CLI"
	"wallet/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	defer bc.Db().Close()

	cli := CLI.CLI{}
	cli.Run()

}