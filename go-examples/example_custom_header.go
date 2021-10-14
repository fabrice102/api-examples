package main

import (
	"context"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/client/v2/common"
)

func main() {
	const algodAddress = "https://testnet-algorand.api.purestake.io/ps2"
	const psTokenKey = "X-API-Key"
	const psToken = "B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab"

	commonClient, err := common.MakeClient(algodAddress, psTokenKey, psToken)
	if err != nil {
		fmt.Printf("failed to make common client: %s\n", err)
		return
	}
	algodClient := (*algod.Client)(commonClient)

	fmt.Println("Status")
	nodeStatus, err := algodClient.Status().Do(context.Background())
	if err != nil {
		fmt.Printf("error getting algod status: %s\n", err)
		return
	}

	fmt.Printf("algod last round: %d\n", nodeStatus.LastRound)
	fmt.Printf("algod time since last round: %d\n", nodeStatus.TimeSinceLastRound)
	fmt.Printf("algod catchup: %d\n", nodeStatus.CatchupTime)
	fmt.Printf("algod latest version: %s\n", nodeStatus.LastVersion)
}
