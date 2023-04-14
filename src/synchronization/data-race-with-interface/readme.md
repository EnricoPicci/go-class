# Examples of data races that can occur when using variables of interface types

## interface-data-race-doers

Example of data races (and potential data corruption) caused by an unprotected assignement (write operation) of concrete values of different types to a varable of interface type

## voters examples

Examples of data races (and potential data corruption) a bit more complex than the one above.

The data race is double, caused by

- an unprotected assignement (write operation) of concrete values of different types to a varable of interface type
- methods of such concrete types that update shared values

The examples show what can heppen when

- there is no control of synchronization ([voters](./voters/) folder)
- there is partial control of synchronization ([voters-vote-synchronized](./voters-vote-synchronized/) and [voters-assignement-synchronized](./voters-assignement-synchronized/) folders)
- there is full control of synchrnonization ([voters-fully-synchronized/](./voters-fully-synchronized/) folder)
