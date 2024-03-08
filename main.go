package main

import (
	"fmt"
	"log"

	"github.com/btcsuite/btcd/wire"
)

func main() {

	// 1 - get decode address for my account
	fmt.Print("****** Step 1/7 = getDecodeAddress...")
	address, err := getDecodeAddress()
	if err != nil {
		log.Fatal("Error in getDecodeAddress:", err)
	}
	fmt.Println(" - finished: ", address.EncodeAddress())

	// 2 - get client rpc to comunicate with bitcoin core
	fmt.Print("****** Step 2/7 = get getClientRpc...")
	client, err := getClientRpc()
	defer client.Shutdown()
	if err != nil {
		log.Fatal("Error in getClientRpc:", err)
	}
	fmt.Println(" - finished.")

	all_send := 0
	all_send_correct := 0
	var block *wire.MsgBlock

	for {
		fmt.Println("****** all_send: ", all_send, " - all_send_correct: ", all_send_correct)
		fmt.Print("****** Step 3/7 = get block...        - ")
		block = getBlock(client)
		fmt.Println(" - finished.")

		fmt.Print("****** Step 4/7 = mine block...      ")
		hash, nonce := mineBlock(block)
		if hash == "" {
			fmt.Println(" - max tries...")
			continue
		}
		fmt.Println(" - Block mined nonce:", nonce, "and hash:", hash)

		fmt.Print("****** Step 5/7 = create new block...")
		btcdBlock, err := createNewBlock(block, address, hash, nonce)
		if err != nil {
			fmt.Println("Error in create new block:", err)
			break
		}
		fmt.Println(" - finished correctly!!!")

		fmt.Print("****** Step 6/7 = serializeBlock...  ")
		btcdBlockSerialized, err := serializeBlock(btcdBlock)
		if err != nil {
			fmt.Println("Error in serializeBlock:", err)
			break
		}
		fmt.Println(" - finished correctly!!!")

		fmt.Print("****** Step 7/7 = send block...      ")
		all_send++
		err = sendBlock(client, btcdBlockSerialized)
		if err != nil {
			fmt.Println(" - Error in sendBlock:", err)
		} else {
			all_send_correct++
			fmt.Println(" - finished correctly - we can hope!")
		}
	}

}
