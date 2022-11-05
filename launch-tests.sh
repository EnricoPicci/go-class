#!/bin/bash

# https://stackoverflow.com/questions/63982271/start-server-in-background-run-code-and-stop-server-again-in-single-script

start_servers() {
    echo "Start servers ..."

    echo "Start net-io-bound-processing-server"
    ./bin/net-io-bound-processing-server -sleep 100  -port 8080 &
    server_1_pid=$!
    # wait for the messages produced by the ./bin/net-io-bound-processing-server to be printed before continuing
    sleep 0.1
    echo

    # Wait for the server to start (max 10 seconds)
    for attempt in {1..10}; do
        my_pid=$(lsof -t -i tcp:8080)

        if [[ -n $my_pid ]]; then
            # Make sure the running server is the one we just started.
            if [[ $my_pid -ne $server_1_pid ]]; then
                echo "ERROR: Multiple Servers running."
                echo "Kill the other server with the command"
                echo "lsof -t -i tcp:8080 | xargs kill"
                echo "and then launch again this program"
                exit 1
            fi

            break
        fi

        sleep 1
    done

    if [[ -z $my_pid ]]; then
        echo "ERROR: Timeout while waiting for Server to start"
        exit 1
    fi

    echo "Start interfaces/reader-writer/server/"
    ./bin/reader-writer-server &
    server_2_pid=$!
    # wait for the messages produced by the ./bin/net-io-bound-processing-server to be printed before continuing
    sleep 0.1
    echo

    # Wait for the server to start (max 10 seconds)
    for attempt in {1..10}; do
        my_pid=$(lsof -t -i tcp:8081)

        if [[ -n $my_pid ]]; then
            # Make sure the running server is the one we just started.
            if [[ $my_pid -ne $server_2_pid ]]; then
                echo "ERROR: Multiple Servers running."
                echo "Kill the other server with the command"
                echo "lsof -t -i tcp:8081 | xargs kill"
                echo "and then launch again this program"
                exit 1
            fi

            break
        fi

        sleep 1
    done

    if [[ -z $my_pid ]]; then
        echo "ERROR: Timeout while waiting for Server to start"
        exit 1
    fi
}

stop_servers() {
    echo "Stop Servers ..."
    kill $server_1_pid
    kill $server_2_pid
}

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 

go build -o ./bin/net-io-bound-processing-server ./src/goroutines/io-bound-processing/net-io/server
go build -o ./bin/reader-writer-server ./src/interfaces/reader-writer/server

# create the directories used for the tests
mkdir out
mkdir out/local

start_servers

printf "\n run the tests \n\n"

# -count=1 forces to run also the cached tests
go test ./... -count=1

stop_servers