package main

import (
	gochains "github.com/Liberontissauri/blockchains-in-go/blockchain"
)

type list_blocks_response struct {
	Blocks []*gochains.Block `json:"blocks"`
}