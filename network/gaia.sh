#!/bin/bash

BINARY=gaiad
MNEMONIC_1="alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"
CHAINID=test-1

# Stop if it is already running 
if pgrep -x  "$BINARY" >/dev/null; then
    echo "Terminating $BINARY..."
    killall gaiad
fi

echo "Removing previous data..."
rm -rf ~/.gaia &> /dev/null

# Add directories for both chains, exit if an error occurs
if ! mkdir -p ~/.gaia 2>/dev/null; then
    echo "Failed to create gaiad folder. Aborting..."
    exit 1
fi

echo "Initializing $CHAINID..."
gaiad init test --chain-id=$CHAINID

echo "Adding genesis accounts..."
$BINARY keys add val --keyring-backend=test
$BINARY keys add demowallet --keyring-backend=test
echo $MNEMONIC_1 | $BINARY keys add rly --recover --keyring-backend=test 

$BINARY add-genesis-account $($BINARY keys show val --keyring-backend test -a) 100000000000stake 
$BINARY add-genesis-account $($BINARY keys show demowallet --keyring-backend test -a) 100000000000stake 
$BINARY add-genesis-account $($BINARY keys show rly --keyring-backend test -a) 100000000000stake 

echo "Creating and collecting gentx..."
$BINARY gentx val 7000000000stake --chain-id $CHAINID --keyring-backend test
$BINARY collect-gentxs

echo "Changing defaults and ports in app.toml and config.toml files..."
sed -i -e 's/"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:26657/g' ~/.gaia/config/config.toml
sed -i -e 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.gaia/config/config.toml
sed -i -e 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.gaia/config/config.toml
sed -i -e 's/index_all_keys = false/index_all_keys = true/g' ~/.gaia/config/config.toml
sed -i -e 's/enable = false/enable = true/g' ~/.gaia/config/app.toml
sed -i -e 's/swagger = false/swagger = true/g' ~/.gaia/config/app.toml

# Update host chain genesis to allow x/bank/MsgSend ICA tx execution
sed -i -e 's/\"allow_messages\":.*/\"allow_messages\": [\"\/cosmos.bank.v1beta1.MsgSend\", \"\/cosmos.staking.v1beta1.MsgDelegate\", \"\/ibc.applications.transfer.v1.MsgTransfer\"]/g' ~/.gaia/config/genesis.json

echo "Starting $CHAINID in ~/.gaia ..."
echo "Creating log file at gaia.log"
$BINARY start --log_level trace --log_format json --pruning=nothing > gaia.log 2>&1 &
