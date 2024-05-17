#!/bin/bash

$BOOKKEEPER_HOME/bin/bookkeeper shell bookieformat -deleteCookie -n -f

$BOOKKEEPER_HOME/bin/bookkeeper shell metaformat -n -f

$BOOKKEEPER_HOME/bin/bookkeeper bookie >>$BOOKKEEPER_HOME/logs/bookkeeper.stdout.log 2>>$BOOKKEEPER_HOME/logs/bookkeeper.stderr.log
