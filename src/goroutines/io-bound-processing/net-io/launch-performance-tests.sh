#!/bin/bash

# https://stackoverflow.com/questions/63982271/start-server-in-background-run-code-and-stop-server-again-in-single-script

start_server() {
    echo "Start server ..."
    ./bin/net-io-bound-processing-server -sleep 100  -port 8080 &
    server_pid=$!
    # wait for the messages produced by the ./bin/net-io-bound-processing-server to be printed before continuing
    sleep 0.1
    echo

    # Wait for the server to start (max 10 seconds)
    for attempt in {1..10}; do
        my_pid=$(lsof -t -i tcp:8080)

        if [[ -n $my_pid ]]; then
            # Make sure the running server is the one we just started.
            if [[ $my_pid -ne $server_pid ]]; then
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
}

stop_server() {
    echo "Stop Server ..."
    kill $server_pid
}

# - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - 


echo "Build server"
go build -o ./bin/net-io-bound-processing-server ./src/goroutines/io-bound-processing/net-io/server
echo "Build client"
go build -o ./bin/performance-test ./src/goroutines/io-bound-processing/net-io/client/performance-test

start_server

./bin/performance-test

stop_server