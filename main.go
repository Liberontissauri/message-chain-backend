package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	gochains "github.com/Liberontissauri/blockchains-in-go/blockchain"
	"github.com/gorilla/mux"
)

var message_blockchain *gochains.Blockchain

func main() {
	message_blockchain = gochains.CreateBlockchain(2, 60 * 10)
	router := mux.NewRouter()

	router.HandleFunc("/api/blocks", listBlocks).Methods("GET")
	router.HandleFunc("/api/blocks", addBlock).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func listBlocks(w http.ResponseWriter, r *http.Request) {
	encoded_block_list, _ := json.Marshal(list_blocks_response{message_blockchain.GetBlocks()})
	w.Write(encoded_block_list)
	w.Header().Set("Content-Type", "application/json")
}

func addBlock(w http.ResponseWriter, r *http.Request) {
	var req_block *gochains.Block
	req_body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(req_body, &req_block)
	block := gochains.CreateNewBlock(req_block.Header.Data, req_block.Header.PrevHash, req_block.Header.Nonce, req_block.Header.Target, req_block.Hash)
	block.Header.Timestamp = req_block.Header.Timestamp
	if(block.IsValid()) {
		message_blockchain.AddBlock(block)
		if(message_blockchain.ValidateBlockchain()) {
			w.WriteHeader(http.StatusOK)
		} else {
			message_blockchain.RemoveTopBlock()
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}