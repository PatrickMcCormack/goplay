#!/usr/bin/env bash

#
# launch.sh - start a server on the local machine on a
#  specified port. Ports are configured in servers.cfg.
#

source raft.cfg

NUM_SERVERS=${#SERVERS[@]}

if [ -z $1 ]; then
        echo "Usage: $0 <server-number>"
        exit 1
fi

if [ $1 -le 0 ] || [ $1 -gt $NUM_SERVERS ]; then
	echo "Launch failed, server $1 is outside the range of the number of configured servers."
	exit 1
fi

echo "Launching server" $1

COUNTER=1
THIS_SERVER=""
REMOTE_SERVERS=""
for i in ${SERVERS[@]}; do
	if [ $COUNTER -eq $1 ]; then
    THIS_SERVER=$i
	else
    if [ "$REMOTE_SERVERS" = "" ]; then
      REMOTE_SERVERS=$i
    else
      REMOTE_SERVERS=$REMOTE_SERVERS","$i
    fi
	fi
	COUNTER=$[COUNTER + 1]
done

LOCAL_PORT=`echo $THIS_SERVER | sed -e 's/^.*://g'`
LOCAL_SERVER="localhost:"$LOCAL_PORT

# note - this is a little error prone. If you give the wrong
# index into the config for the node you want to start it will
# start but on the wrong port. Trying to be too clever here,
# clean up by requiring a config file per instance when running
# everything on a single box (like ZK does).

# run the raft server
go run raftmain.go --verbose --local $LOCAL_SERVER --cluster $REMOTE_SERVERS
