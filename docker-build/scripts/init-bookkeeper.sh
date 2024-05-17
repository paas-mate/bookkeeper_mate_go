#!/bin/bash

$BOOKKEEPER_HOME/bin/bookkeeper shell bookieformat -deleteCookie -n -f

$BOOKKEEPER_HOME/bin/bookkeeper shell metaformat -n -f

$BOOKKEEPER_HOME/bin/bookkeeper bookie
