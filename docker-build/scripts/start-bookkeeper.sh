#!/bin/bash

OPTS="OPTS="

if ! [ -z ${BK_METADATA_BOOKIE_DRIVERS_PROPERTY} ]; then
   OPTS="${OPTS} -Dbookkeeper.metadata.bookie.drivers=org.apache.bookkeeper.metadata.etcd.EtcdMetadataBookieDriver"
fi

if ! [ -z ${BK_METADATA_CLIENT_DRIVERS_PROPERTY} ]; then
    OPTS="${OPTS} -Dbookkeeper.metadata.client.drivers=org.apache.bookkeeper.metadata.etcd.EtcdMetadataClientDriver"
fi

export ${OPTS}

$BOOKKEEPER_HOME/bin/bookkeeper bookie >>$BOOKKEEPER_HOME/logs/bookkeeper.stdout.log 2>>$BOOKKEEPER_HOME/logs/bookkeeper.stderr.log
