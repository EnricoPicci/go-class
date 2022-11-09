# use external modules with internal folder

This is an example of a package, the main package stored in the `src/use-external-module-internal/main` folder, that uses an external module stored in a git repo.

This external module has some packages residing below a folder names `internal`. The exported names (e.g. variables and functions) exported by the pacakges below the `internal` folder are not visible for this package even if they are exported (i.e. they start with a capital letter). These exported names are visible only within the module that defines them.

This can be seen in the [main.go](./main/main.go) file. If we remove the comment from the `import "github.com/EnricoPicci/go-class-hello-with-internal/src/internal/print"` line, we see that an error appears.

## build

Build the executable running the following command from the go-class project directory

`go build -o ./bin/use-external-module-internal ./src/use-external-module-internal/main`

## run

Run the example calling the following command from the go-class project directory

`./bin/use-external-module-internal`
