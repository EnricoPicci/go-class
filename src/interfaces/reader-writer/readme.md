# Reader-Writer

This is an example of how to create types that comply with interfaces defined in other packages and use the APIs provided by such packages.

This program, once launched, copies what is keyed in the console (stdin) to a file. The file can be local or can be remote. For copying to a remote file to work, a server (coded in the `server` folder) has to be launched.

In this particular example, we create a type `main.stdin` that implements `io.Reader` interface and a type `remote.File` that implements the `io.WriteCloser` interface. The real stdin is wrapped by a custom struct type `main.stdin`, so that it is possible to end the program by writing a specific "quit" command.

The program uses also the type `os.File` which also implements the `io.WriteCloser` interface.

## Where interfaces are used

The function that shows how to use concrete types that implement the required interfaces is `copyFromStdinTo`.

## Run the example

To run the example program start first the server with the following command

`go run ./src/interfaces/reader-writer/server/`

and then start the client with the command

`go run ./src/interfaces/reader-writer/stdin-to-file/ -file ./out/abc.txt` to write to the local abc.txt file

or with the command

`go run ./src/interfaces/reader-writer/stdin-to-file/ -file abc.txt -remote` to write to a remote file abc.txt
