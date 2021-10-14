// Algorand Indexer (v2) example 
// Demonstrate an Indexer Asset Search with pagination on TestNet

const algosdk = require('algosdk');
const baseServer = "https://testnet-algorand.api.purestake.io/idx2";
const port = "";

const token = {
    'X-API-key': 'B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab',
}

let indexerClient = new algosdk.Indexer(token, baseServer, port);

(async () => {

    const txnInfo = await indexerClient.searchForTransactions().txid('H6QGCDZGS64ZD6SXYUHQKFEG5CTF4VB3JCCT4WAGWRH2LTY7UV3A').do()
    console.log(txnInfo)

})().catch(e => {
    console.log(e);
});



