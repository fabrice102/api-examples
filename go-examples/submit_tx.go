package main

import (
	"context"
	"fmt"

	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/client/v2/common"
	"github.com/algorand/go-algorand-sdk/crypto"
	transaction "github.com/algorand/go-algorand-sdk/future"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

// Function from Algorand Inc. - utility for waiting on a transaction confirmation
func waitForConfirmation(txID string, client *algod.Client) {
	status, err := client.Status().Do(context.Background())
	if err != nil {
		fmt.Printf("error getting algod status: %s\n", err)
		return
	}
	lastRound := status.LastRound
	for {
		pt, _, err := client.PendingTransactionInformation(txID).Do(context.Background())
		if err != nil {
			fmt.Printf("error getting pending transaction: %s\n", err)
			return
		}
		if pt.ConfirmedRound > 0 {
			fmt.Printf("Transaction confirmed in round %d\n", pt.ConfirmedRound)
			break
		}
		fmt.Printf("Waiting for confirmation...\n")
		lastRound++
		status, err = client.StatusAfterBlock(lastRound).Do(context.Background())
	}
}

func main() {
	const algodAddress = "https://testnet-algorand.api.purestake.io/ps2"
	const psTokenKey = "X-API-Key"
	const psToken = "B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab"

	const fromAddr = "UDCFS2TSVK5MM5A3GH4WLZDNGZE4CJT2AOTADUB7MZKW72CXPRZAHKVTVU"
	const mn = "cool online brush identify bean nuclear elder soft fashion mind inside drama camp excess captain window spare oxygen tonight kingdom sustain pigeon predict ability rail"
	const toAddr = "ZHGZZQ2PIWYRK6MIK44GKO3VGQUC7NS2V3UQ63M3DIMFUFGI4BRWK7WDBU"

	fromAddrPvtKey, err := mnemonic.ToPrivateKey(mn)
	if err != nil {
		fmt.Printf("error recovering private key: %s\n", err)
		return
	}

	commonClient, err := common.MakeClient(algodAddress, psTokenKey, psToken)
	if err != nil {
		fmt.Printf("failed to make common client: %s\n", err)
		return
	}
	algodClient := (*algod.Client)(commonClient)

	// Get the suggested transaction parameters
	txParams, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		fmt.Printf("error getting suggested tx params: %s\n", err)
		return
	}
	note := []byte(nil)
	closeRemainderTo := ""
	amount := uint64(10000)

	// Make transaction
	txn, err := transaction.MakePaymentTxn(fromAddr, toAddr, amount, note, closeRemainderTo, txParams)
	if err != nil {
		fmt.Printf("Error creating transaction: %s\n", err)
		return
	}

	// Sign the Transaction
	_, bytes, err := crypto.SignTransaction(fromAddrPvtKey, txn)
	if err != nil {
		fmt.Printf("Failed to sign transaction: %s\n", err)
		return
	}

	// Broadcast the transaction to the network
	sendResponse, err := algodClient.SendRawTransaction(bytes).Do(context.Background())
	if err != nil {
		fmt.Printf("failed to send transaction: %s\n", err)
		return
	}
	fmt.Printf("Transaction sent with ID %s\n", sendResponse)
	waitForConfirmation(sendResponse, algodClient)
}
