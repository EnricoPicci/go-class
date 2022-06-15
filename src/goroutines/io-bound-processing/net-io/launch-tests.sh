#!/bin/bash

# https://stackoverflow.com/questions/63982271/start-server-in-background-run-code-and-stop-server-again-in-single-script

start_server() {
    echo "Start server ..."
    ./bin/net-io-bound-processing-server -sleep 100  -port 8080 &
    server_pid=$!

    # Wait for the server to start (max 10 seconds)
    for attempt in {1..10}; do
        my_pid=$(lsof -t -i tcp:8080)

        if [[ -n $my_pid ]]; then
            # Make sure the running server is the one we just started.
            if [[ $my_pid -ne $server_pid ]]; then
                echo "ERROR: Multiple Servers running."
                echo "→ lsof -t -i tcp:8080 | xargs kill"
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

go build -o ./bin/net-io-bound-processing-server ./server
go build -o ./bin/net-io-bound-processing-client ./client


start_server

printf "\n test with 1 core and no concurrency \n\n"
./bin/net-io-bound-processing-client -requests 100 -maxprocs 1 -concurrent 1

printf "\n test with ALL cores and no concurrency \n\n"
./bin/net-io-bound-processing-client -requests 100 -concurrent 1

printf "\n Compare the results to check how much the variation in number of cores affects the time spent to process all requests \n\n"


printf "\n test with 1 core and 10 concurrent goroutines \n\n"
./bin/net-io-bound-processing-client -requests 100 -maxprocs 1 -concurrent 10

printf "\n test with ALL cores and 10 concurrent goroutines \n\n"
./bin/net-io-bound-processing-client -requests 100 -concurrent 10

printf "\n Compare the results to check how much the variation in number concurrency affects the time spent to process all requests \n\n"


printf "\n test with 1 core and 100 concurrent goroutines \n\n"
./bin/net-io-bound-processing-client -requests 100 -maxprocs 1 -concurrent 100

printf "\n test with ALL cores and 100 concurrent goroutines \n\n"
./bin/net-io-bound-processing-client -requests 100 -concurrent 100

printf "\n Compare the results to check how much the variation in number concurrency affects the time spent to process all requests \n\n"


stop_server