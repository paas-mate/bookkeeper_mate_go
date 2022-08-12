#!/bin/bash

mkdir -p $BOOKKEEPER_HOME/logs
nohup $BOOKKEEPER_HOME/mate/bookkeeper_mate >>$BOOKKEEPER_HOME/logs/bookkeeper_mate.stdout.log 2>>$BOOKKEEPER_HOME/logs/bookkeeper_mate.stderr.log

