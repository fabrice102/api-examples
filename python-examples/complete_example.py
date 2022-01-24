from algosdk.v2client import algod
from algosdk import mnemonic, account
from algosdk.future.transaction import PaymentTxn


# Function from Algorand Inc.
def wait_for_confirmation(client, txid):
    last_round = client.status().get('last-round')
    txinfo = client.pending_transaction_info(txid)
    while not (txinfo.get('confirmed-round') and txinfo.get('confirmed-round') > 0):
        print('Waiting for confirmation')
        last_round += 1
        client.status_after_block(last_round)
        txinfo = client.pending_transaction_info(txid)
    print('Transaction confirmed in round', txinfo.get('confirmed-round'))
    return txinfo


# Setup HTTP client w/guest key provided by PureStake
algod_address = 'https://testnet-algorand.api.purestake.io/ps2'
algod_token = ""
headers = {
    "X-API-Key": "B3SU4KcVKi94Jap2VXkK83xx38bsv95K5UZm2lab",
}

# Initialize throw-away account for this example - check that is has funds before running script
mnemonic_phrase = 'code thrive mouse code badge example pride stereo sell viable adjust planet text close erupt embrace nature upon february weekend humble surprise shrug absorb faint'
account_private_key = mnemonic.to_private_key(mnemonic_phrase)
account_public_key = account.address_from_private_key(account_private_key)

algodclient = algod.AlgodClient(algod_token, algod_address, headers)

# get suggested parameters from Algod
params = algodclient.suggested_params()

send_amount = 10  # amount to send in microAlgos
existing_account = account_public_key
send_to_address = 'AEC4WDHXCDF4B5LBNXXRTB3IJTVJSWUZ4VJ4THPU2QGRJGTA3MIDFN3CQA'

# Create and sign transaction
tx = PaymentTxn(existing_account, params, send_to_address, send_amount)
signed_tx = tx.sign(account_private_key)

try:
    tx_confirm = algodclient.send_transaction(signed_tx)
    print('Transaction sent with ID', signed_tx.transaction.get_txid())
    wait_for_confirmation(algodclient, txid=signed_tx.transaction.get_txid())
except Exception as e:
    print(e)
