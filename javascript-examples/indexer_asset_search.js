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

    let i = 5;

    let assetInfo = await indexerClient.searchForAssets().limit(1).name('test').do()
    console.log(assetInfo)

    while (i && ('next-token' in assetInfo)) {
        assetInfo = await indexerClient.searchForAssets().limit(1).name('test').nextToken(assetInfo['next-token']).do();
        i--;
        console.log(assetInfo) // Array of assets object
        console.log(assetInfo.assets[0]) // An asset returned
    }

})().catch(e => {
    console.log(e);
});



