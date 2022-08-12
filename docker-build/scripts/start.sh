#!/bin/bash

DIR="$( cd "$( dirname "$0"  )" && pwd  )"
$BOOKKEEPER_HOME/mate/config_gen
if [ $CLUSTER_INIT == "true" ]; then
    $BOOKKEEPER_HOME/mate/bookkeeper_mate 2>&1
else
    bash -x $DIR/start-daemon.sh
    tail -f /dev/null
fi