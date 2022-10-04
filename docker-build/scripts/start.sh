#!/bin/bash

DIR="$( cd "$( dirname "$0"  )" && pwd  )"
$BOOKKEEPER_HOME/mate/config_gen
if [ $CLUSTER_INIT == "true" ]; then
    bash -x $DIR/init-bookkeeper.sh
else
    bash -x $DIR/start-daemon.sh
    tail -f /dev/null
fi
