 #!/bin/bash
HERMES_BINARY=hermes
HERMES_DIRECTORY=./hermes/
CONFIG_DIR=./hermes/gaia-config.toml

echo "Using hermes relayer version: "
$HERMES_BINARY version | sed 's/^/    /'
