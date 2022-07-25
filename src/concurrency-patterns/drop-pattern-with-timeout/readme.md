# Drop pattern with timeout

In this example we implement the drop pattern with timeout.

A certain number of requests is sent to a worker pool.

Each request takes a certain amount of time to be processed.

Each request is given a context with a timeout when it comes in. The request waits for the pool to start processing it or for its context to time out. This is how the drop pattern with timeout is implemented.

These parameters have default values which can be overridden by command line params

- poolSize: number of workers in the worker pool
- timeout: timeout in miliseconds after which a request waiting to be processed is dropped
- avgReqInterval: average interval between requests in miliseconds
- procIntervalRatio: ratio between average time spent to process a request and the interval between subsequent requests
- numReq: number of requests coming in to be processed
- extraordinaryEvent: if true, it is possible that an extraordinary event occurs that generates a peak in processing time (this simulates for instance a transient problem on a DB that slows down the processing of some requests)

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/drop-pattern-with-timeout ./src/concurrency-patterns/drop-pattern-with-timeout`

## run

Playing with the parameters passed by command line allows to see how, in different conditions, the drop pattern can behave differently.

### drop few (maybe none) incoming requests

Let's consider a poolSize of 5 (i.e. 5 workers working concurrently) and a procIntervalRatio of 10 (i.e. in the time required to process 1 request 10 new requests come in).

Let's consider 100 requests coming in with, on average, a new request every 100ms. When we have received all the 100 request, on average, we can expect to have a "queue" of about 50 requests that are waiting for the pool to start processing them (why 50 requests in the queue? because, if the pool is able to process 5 requests in a time interval where 10 new requests come in, it is logical to expect to have a queue of about 50% of the requests received).

If the pool size is 5, a new request comes in every 100ms and each request takes 1000ms to be processed (this is what a procIntervalRatio of 10 means), then every 1000ms we process 5 requests (the size of the pool) which means that, on average, every 200ms a request gets out of the pool processed.

Let's consider the last request in the queue. When the last request enters the queue, the queue is of about 50 requests (for the reasoning done above). So, the last request in the queue has, on average, to wait 50 (the number of requests in the queue) \* 200ms (the average throughput of the pool) = 10000ms.

So, if we put a timeout of 10000ms, we should expect to see very few, if not none, requests dropped and the max size of the queue to be close to 50.

From the GO-CLASS project folder run the command
`./bin/drop-pattern-with-timeout -poolSize 5 -procIntervalRatio 10 -avgReqInterval 100 -numReq 100 -timeout 10000`

### minimize the wait time by increasing the pool size

If we increase the size of the worker pool, we increase the throughput of the pool and therefore e decrease the time each request waits to be processed.

Run the same example as before with a pool of 100 workers and see the value of "Max waiting queue size" and "Max wait duration" printed on the console at the end of the processing.

While this is a way to reduce the wait time for each request, it has a cost represented by the time spent by workers staying idle waiting for a request to process (there are many free workers for any request coming in) See the value "Average idle time for a worker" printed on the console at the end of the processing and compare it with the value produced by the next example.

From the GO-CLASS project folder run the command
`./bin/drop-pattern-with-timeout -poolSize 100 -procIntervalRatio 10 -avgReqInterval 10 -numReq 1000 -timeout 1000`

### find the right balance between worker pool throughput and frequency of incoming requests

If we increase the size of the worker pool to be able to absord the flow of request withouth having the workers stay idle for long time, then we have reached the right balance.

Run the same example as before with a pool of 10 workers. Given that procIntervalRatio is 10, this means that on average in the time 1 request gets processed 10 new requests arrive. Having a pool of 10 workers should bring us to the point of equilibirum where the flow of requests get processed steadily and the wait time tends to remain stable. Also the time spent idle by workers is low (see the value "Average idle time for a worker" printed on the console at the end of the processing).

If we reduce the pool size to 9, the balance is lost and the wait time tend to increase. With this configuration some requests may be dropped.

From the GO-CLASS project folder run the command
`./bin/drop-pattern-with-timeout -poolSize 10 -procIntervalRatio 10 -avgReqInterval 10 -numReq 1000 -timeout 1000`
`./bin/drop-pattern-with-timeout -poolSize 9 -procIntervalRatio 10 -avgReqInterval 10 -numReq 1000 -timeout 1000`

### drop in case of extraordinary events

The drop pattern is mainly conceived to produce clear signals (to the clients and to the system) when some extraordinary events that slow down the processing occur.

So we typically should point to a configuration that guarantees the right balance between the flow of incoming requests and the throughput of the system. When extraordinary conditions occur that slow down processing, the drop pattern drops requests that have been waiting too long generating a clear signal of an hopefully transient exceptional condition.

From the GO-CLASS project folder run the command
`./bin/drop-pattern-with-timeout -poolSize 10 -procIntervalRatio 10 -avgReqInterval 10 -numReq 1000 -timeout 1000 -extraordinaryEvent`

## build and run with race detector

To build the package with the race detector on, from the GO-CLASS project folder run the command
`o build -race -o ./bin/drop-pattern-with-timeout ./src/concurrency-patterns/drop-pattern-with-timeout`

To run it, just run it with the usual commands simply writing the stdOutput to a file so that race messages (by default written on stdErr) are the only ones left on the console
`./bin/drop-pattern-with-timeout -poolSize 10 -procIntervalRatio 10 -avgReqInterval 10 -numReq 1000 -timeout 1000 > race.txt`
