#!/bin/bash

NAMEROOT="root"
USERNAME="node_1"
# List addresses of servers delimeted by spaces
SERVERS="127.0.0.1"



LOCAL_BIN_FILE_PATH="/home/egorka/oton/q/tendermint_89"


REMOTE_WORKDIR="/home/node_1/.tendermint"
REMOTE_BIN_FILE_PATH="/home/node_1/.tendermint/tendermint_89"

REMOTE_CONFIG_FILE_PATH="~/.tendermint/config/config.toml"
REMOTE_GENESIS_FILE_PATH="~/.tendermint/config/genesis.json"
REMOTE_ADDRBOOK_FILE_PATH="~/.tendermint/config/addrbook.json"

SSH_PARAMS="-o StrictHostKeyChecking=no"

for ADDRESS in $SERVERS
do
    echo
    echo "Working with host $ADDRESS"
    echo

    SCREEN_PID=$(ssh $SSH_PARAMS $NAMEROOT@$ADDRESS "screen -ls" | awk '/\.NODE\t/ {print strtonum($1)}')

    if [ -n "$SCREEN_PID" ]
    then
        echo "Stoping screen instance with name NODE and PID $SCREEN_PID"
        ssh $SSH_PARAMS $NAMEROOT@$ADDRESS "kill -s SIGINT $SCREEN_PID"

        while true
        do
            if [ $(ssh $SSH_PARAMS $NAMEROOT@$ADDRESS "ps aux" | grep $SCREEN_PID | wc -l) -gt 1 ]
            then
                echo "Node is still stoping..."
                sleep 1
            else
                break
            fi
        done
    else
        echo "Node is not running on this server. Nothing to stop!"
    fi


    echo "Install mc"
    ssh $SSH_PARAMS $NAMEROOT@$ADDRESS "apt -y install mc"


    echo "Fix Kill"
    ssh $SSH_PARAMS $NAMEROOT@$ADDRESS "pkill -f NODE"
    ssh $SSH_PARAMS $NAMEROOT@$ADDRESS "pkill -f tendermint_89"

    echo "Cpy old tendermint file"
    ssh $SSH_PARAMS $NAMEROOT@$ADDRESS "runuser -l node_1 -c 'cp -R /.tendermint  /home/node_1'"

    echo "Uploading tendermint bin file"
    scp $SSH_PARAMS $LOCAL_BIN_FILE_PATH $NAMEROOT@$ADDRESS:$REMOTE_BIN_FILE_PATH

    echo "Starting tendermint in detached screen with name NODE"
    ssh $SSH_PARAMS $NAMEROOT@$ADDRESS "runuser -l node_1 -c 'screen -dmS NODE  -L -Logfile /home/node_1/.tendermint/work-$(date +%Y-%m-%d_%H%M).log  $REMOTE_BIN_FILE_PATH start --proxy-app=otoncoin' "
done



