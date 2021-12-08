# Interchain Accounts
## Local Demo

### Setup

```bash
# Clone this repository and build
git clone https://github.com/cosmos/interchain-accounts.git
cd interchain-accounts
make install 

# Hermes Relayer
[Hermes](https://hermes.informal.systems/) is a Rust implementation of a relayer for the [Inter-Blockchain Communication (IBC)](https://ibcprotocol.org/) protocol.

In order to use the hermes relayer you will need to check out a specific branch that can be used with interchain-accounts. 

git clone https://github.com/informalsystems/ibc-rs
git checkout 97360a0c
cd relayer-cli
cargo build

In the variables.sh file inside /network/hermes/ replace the $HERMES_BINARY variable with a path to the hermes binary build from the previous step. You can find this in the /target/debug/ directory inside ibc-rs. 
# Bootstrap two local chains & create a connection using the hermes relayer
make init

# Wait for the ibc connection & channel handshake to complete and the relayer to start
```

### Demo

```bash
# Open a seperate terminal

# Store the following account addresses within the current shell env
export testwallet1=$(icad keys show testwallet1 -a --keyring-backend test --home ./data/test-1)
echo $testwallet1
export testwallet2=$(icad keys show testwallet2 -a --keyring-backend test --home ./data/test-2)
echo $testwallet2

# Register an IBC Account on chain test-2 
icad tx intertx register --from $testwallet1 --connection-id connection-0 --counterparty-connection-id connection-0 --chain-id test-1 --gas 150000 --home ./data/test-1 --node tcp://localhost:16657 --keyring-backend test -y

# Start the hermes relayer in the first terminal
# This will also finish the channel creation handshake signalled during the register step
make start-rly

# Query the address of the interchain account
icad query intertx interchainaccounts $testwallet1 connection-0 connection-0 --home ./data/test-1 --node tcp://localhost:16657

# Store the interchain account address by parsing the query result
export ICA_ADDR=$(icad query intertx interchainaccounts $testwallet1 connection-0 connection-0 --home ./data/test-1 --node tcp://localhost:16657 -o json | jq -r '.interchain_account_address')

# Check the interchain account's balance on test-2 chain. It should be empty.
icad q bank balances $ICA_ADDR --chain-id test-2 --node tcp://localhost:26657

# Send some assets to $IBC_ACCOUNT.
icad tx bank send $testwallet2 $ICA_ADDR 10000stake --chain-id test-2 --home ./data/test-2 --node tcp://localhost:26657 --keyring-backend test -y

# Check that the balance has been updated
icad q bank balances $ICA_ADDR --chain-id test-2 --node tcp://localhost:26657

# Test sending assets from interchain account via ibc.
icad tx intertx send $ICA_ADDR $testwallet2 5000stake --connection-id connection-0 --counterparty-connection-id connection-0 --chain-id test-1 --gas 90000 --home ./data/test-1 --node tcp://localhost:16657 --from $testwallet1 --keyring-backend test -y

# Wait until the relayer has relayed the packet

# Check if the balance has been changed (it should now be 500stake)
icad q bank balances $ICA_ADDR --chain-id test-2 --node tcp://localhost:26657
```

## Collaboration

Please use conventional commits  https://www.conventionalcommits.org/en/v1.0.0/

```
chore(bump): bumping version to 2.0
fix(bug): fixing issue with...
feat(featurex): adding feature...
```
