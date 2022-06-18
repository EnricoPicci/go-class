# User likes

Show how operations which operate on shared memory but are not correctly synchronized can cause problems

## Unprotected (i.e. non sychronized) updates of likes

Many users concurrently send likes to a user very much loved. Each user sending likes is a goroutine, which is very reasonable if we think at
a web scenario where different users logged in are served by different goroutines runnin in parallel.

The field likes of the user which is sent likes is incremented without any synchronization protection.

### build

From the GO-CLASS project folder run the command
`go build -o ./bin/user-likes-unprotected ./src/synchronization/shared-values/user-likes/write-only/unprotected`

### run

From the GO-CLASS project folder run the command
`./bin/user-likes-unprotected`

## Protecting the likes field with a mutex

If we add to the user struct a mutex, we can use it to protect the method that increments the likes

### build

From the GO-CLASS project folder run the command
`go build -o ./bin/user-likes-protected ./src/synchronization/shared-values/user-likes/write-only/protected`

### run

From the GO-CLASS project folder run the command
`./bin/user-likes-protected`

## Strict read write protection

If we add to the user struct a mutex, we can use it to protect the method that increments the likes and also to protect the method that reads the likes

### build

From the GO-CLASS project folder run the command
`go build -o ./bin/user-likes-mutex-protection ./src/synchronization/shared-values/user-likes/read-write/mutex-protection`

### run

From the GO-CLASS project folder run the command
`./bin/user-likes-mutex-protection`

## Read write protection with RWMutex

If we RWMutex, we can use it to protect the method that increments the likes with the usual Lock, but then we can protect the method that reads the likes using the RLock method which does not synchronize (i.e. sequentialize) reac operations.

### build

From the GO-CLASS project folder run the command
`go build -o ./bin/user-likes-read-write-mutex-protection ./src/synchronization/shared-values/user-likes/read-write/read-write-mutex-protection`

### run

From the GO-CLASS project folder run the command
`./bin/user-likes-read-write-mutex-protection`
