#!/bin/bash
set -e

# Load shell variables
. ./hermes/gaia-variables.sh

### Sleep is needed otherwise the relayer crashes when trying to init
sleep 1s
### Restore Keys
$HERMES_BINARY -c ./hermes/gaia-config.toml keys restore test-1 -m "alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"
sleep 5s

$HERMES_BINARY -c ./hermes/gaia-config.toml keys restore test-2 -m "alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"
sleep 5s
