# use external modules but retract some versions

This is an example of a package, the main package stored in the `src/use-external-module-retract/main` folder, that uses an external module stored in a git repo. Such external module has the version v0.1.2 which was accidentally published and that has subsequently deprecated (retracted) and therefore is not avaible to be used by this main package.

This can be seen if we run the command

`go get github.com/EnricoPicci/go-class-hello-retract@v0.1.2`

which will return an error stating that this version has been retracted.

The [go.mod](../../go.mod) and the [go.sum](../../go.sum) files are those of the GO-CLASS project since the `src/use-external-module` folder does not contain and `go.mod` file and therefore is not a Go module.

## build

Build the executable running the following command from the go-class project directory

`go build -o ./bin/use-external-module-retract ./src/use-external-module-retract/main`

## run

Run the example calling the following command from the go-class project directory

`./bin/use-external-module-retract`
