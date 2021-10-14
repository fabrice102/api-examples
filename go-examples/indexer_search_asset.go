package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/common"
	"github.com/algorand/go-algorand-sdk/client/v2/indexer"
)

func main() {
	const indexerAddress = "https://testnet-algorand.api.purestake.io/idx2"
	const psTokenKey = "X-API-Key"
	const psToken = "B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab"

	commonClient, err := common.MakeClient(indexerAddress, psTokenKey, psToken)
	if err != nil {
		fmt.Printf("failed to make common client: %s\n", err)
		return
	}
	indexerClient := (*indexer.Client)(commonClient)

	// Lookup asset
	name := "test"
	limit := uint64(1)
	result, err := indexerClient.SearchForAssets().Name(name).Limit(limit).Do(context.Background())

	// Search asset by name
	JSON, err := json.MarshalIndent(result, "", "\t")
	fmt.Printf(string(JSON) + "\n")
}
