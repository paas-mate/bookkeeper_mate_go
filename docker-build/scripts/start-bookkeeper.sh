#!/bin/bash

mkdir -p $BOOKKEEPER_HOME/logs
$BOOKKEEPER_HOME/bin/bookkeeper bookie >>$BOOKKEEPER_HOME/bookkeeper.stdout.log 2>>$BOOKKEEPER_HOME/bookkeeper.stderr.log

