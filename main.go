package main

import (
	"Go_Blockchain/Blockchain"
	"bufio"
	"fmt"
	"os"
)

func Help() {
	fmt.Println("There are 3 operations:")
	fmt.Println("Type cre for adding a new Block")
	fmt.Println("Type view for printing the Blockchain")
	fmt.Println("Type exit for exiting")
}

func main() {
	fmt.Println("Welcome to our Blockchain project.")
	fmt.Println("Enter h for help")

	var (
		operation string
	)

	NewBlockchain := Blockchain.CreateBlockchain() // 新增一個區塊鏈
	fmt.Println("Now we have our first block (or be specifically its pointer)")
	fmt.Println("You can type 'cre' for create new block")
	for { // If the condition for "for" is absent, it is equivalent to the boolean value true.
		_, _ = fmt.Scanln(&operation) // 讀取輸入值
		if operation == "h" {
			fmt.Println("Printing the help")
			Help() // 顯示使用方法
		} else if operation == "cre" {
			fmt.Println("Enter your data:")
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine() // 讀一整行 input
			NewBlockchain.AddBlock(data)    // 利用 input 作為 data 來創建區塊鏈
			fmt.Println("block added successfully, please enter help, cre, view or exit")
		} else if operation == "view" {
			for _, block := range NewBlockchain.Blocks {
				fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
				fmt.Printf("Data: %s\n", block.Data)
				fmt.Printf("Hash: %x\n", block.Hash)
				fmt.Println()
			} // 查詢資料
		} else if operation == "exit" {
			break
		} else {
			fmt.Println("Please enter help, cre, view or exit")
		}
	}
}
