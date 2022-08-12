#!/bin/bash

mkdir -p $BOOKKEEPER_HOME/logs
$BOOKKEEPER_HOME/bin/bookkeeper standalone >>$BOOKKEEPER_HOME/logs/bookkeeper.stdout.log 2>>$BOOKKEEPER_HOME/logs/bookkeeper.stderr.log
