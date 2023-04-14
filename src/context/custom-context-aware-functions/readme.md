# Context aware functions

This folder contains 2 examples of implementations of context aware functions.

- The [first example](./context-aware-scanning-function/) is the implementation of a function that process a file line by line using an bufio.Scanner. The function is passed a context. If the context timeouts or is cancelled during the processing of the file, then the processing is terminated and the resources are freed.

- The [second example](./context-aware-io-reader/) is the implementation of a context-aware Reader. This Reader stops reading if its context timeouts or is cancelled.
