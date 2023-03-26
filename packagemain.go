package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var transactions [][]byte
	lines := bytes.Split(data, []byte("\n"))
	for _, line := range lines {
		if len(line) > 0 {
			transaction, err := hex.DecodeString(string(line))
			if err != nil {
				log.Fatal(err)
			}
			transactions = append(transactions, transaction)
		}
	}
	root := merkleRoot(transactions)
	fmt.Printf("Merkle tree root: %x\n", root)
}

func merkleRoot(transactions [][]byte) []byte {
	if len(transactions) == 1 {
		return hash(transactions[0])
	}
	var hashes [][]byte
	for i := 0; i < len(transactions); i += 2 {
		if i+1 == len(transactions) {
			transactions = append(transactions, transactions[i])
		}
		hashes = append(hashes, hash(append(transactions[i], transactions[i+1]...)))
	}
	return merkleRoot(hashes)
}

func hash(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}
