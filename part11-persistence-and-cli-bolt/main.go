package main

import (
	"fmt"
	"publicchain/part11-persistence-and-cli-bolt/BLC"
)

func main() {
	blockchain := BLC.NewBlockchain()
	fmt.Println(blockchain)

	fmt.Printf("tip: %x\n", blockchain.Tip)
	fmt.Println(blockchain.DB)

	// blockchain.AddBlock("send 20 BTC to xiaomin from wangkun")
	// blockchain.AddBlock("send 33 BTC to xiaowang from wangkun")
	// blockchain.AddBlock("send 40 BTC to xiaoxiao from kunkun")

	// for _, block := range blockchain.Blocks {
	// 	fmt.Printf("Data: %s\n", string(block.Data))
	// 	fmt.Printf("PrevBlockHash：%x \n", block.PrevBlockHash)
	// 	fmt.Printf("Timestamp：%s \n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
	// 	fmt.Printf("Hash：%x \n", block.Hash)
	// 	fmt.Printf("Nonce：%d \n", block.Nonce)

	// 	fmt.Println()
	// }

}
