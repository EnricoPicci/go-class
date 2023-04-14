# Drop pattern

In this example we implement the drop pattern.

A certain number of requests is sent to a worker pool.

Each request takes a certain amount of time to be processed.

A buffered channel is put before the worker pool to implement the drop pattern.

These parameters have default values which can be overridden by command line params

- poolSize: number of workers in the worker pool
- chanCap: capacity of the channel
- avgReqInterval: average interval between requests in miliseconds
- procIntervalRatio: ratio between average time spent to process a request and the interval between subsequent requests
- numReq: number of requests coming in to be processed

## build

From the GO-CLASS project folder run the command
`go build -o ./bin/drop-pattern ./src/concurrency-patterns/drop-pattern`

## run

Playing with the parameters passed by command line allows to see how, in different conditions, the drop pattern can behave differently.

### drop about 50% of incoming reuests

If we have a chanCap of 0 (no buffer), a poolSize of 5 (i.e. 5 workers working concurrently) and a procIntervalRatio of 10 (i.e. in the time interval required to process 1 request, 10 new requests come in), the drop pattern drops about 50% of the requests. The reason is that if, for instance, the processing time is P and in this P interval we receive 10 requests, a pool of 5 can process only half of them and therefore half get dropped.

NOTE. Actually in this configuration the percentage of dropped requests tends to be a bit more than 50%, since the time required to process a request is acutally longer that the simple processing time since it has also to consider the time spent in the send/receive operation on the channel. The smaller the value of avgReqInterval (average interval between requests in miliseconds), the smaller is the the processing time (since it is linked to the avgReqInterval by the value of procIntervalRatio) the more the effect is of the time spent in the send/receive operation in the total time taken to process a request, the higher therefore becomes the percentage of dropped requests.

From the GO-CLASS project folder run the command
`./bin/drop-pattern -chanCap 0 -poolSize 5 -procIntervalRatio 10 -avgReqInterval 100 -numReq 100`

### drop few requests increasing the buffer size

If we increase the buffer size we reduce the number of dropped requests since we give to the request a space where they can safely stay while the worker pool is busy processing previous requests. The higher the buffer size the lower the number of drobbed requests.

If we set the buffer size to half the number of requests, than we get close to no dropped requests since, on average, there is space for 50% of the requests to wait for their time to be processed.

From the GO-CLASS project folder run the command
`./bin/drop-pattern -chanCap 50 -poolSize 5 -procIntervalRatio 10 -avgReqInterval 100 -numReq 100`
