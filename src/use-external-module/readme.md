# use external modules

This is an example of a package, the main package stored in the `src/use-external-module/main` folder, that uses an external module stored in a git repo.

The [go.mod](../../go.mod) and the [go.sum](../../go.sum) files are those of the GO-CLASS project since the `src/use-external-module` folder does not contain and `go.mod` file and therefore is not a Go module.

## build

Build the executable running the following command from the go-class project directory

`go build -o ./bin/use-external-module ./src/use-external-module/main`

## run

Run the example calling the following command from the go-class project directory

`./bin/use-external-module`
